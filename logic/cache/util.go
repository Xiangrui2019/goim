package cache

import (
	"errors"
	"goim/logic/db"
	"goim/public/logger"
	"goim/public/util"
	"time"

	"github.com/json-iterator/go"
)

// set 将指定值设置到redis中，使用json的序列化方式
func set(key string, value interface{}, duration time.Duration) error {
	bytes, err := jsoniter.Marshal(value)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	err = db.RedisCli.Set(key, bytes, duration).Err()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// get 从redis中读取指定值，使用json的反序列化方式
func get(key string, value interface{}) error {
	bytes, err := db.RedisCli.Get(key).Bytes()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	err = jsoniter.Unmarshal(bytes, value)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

const (
	KeyNotExists int64 = 0
	KeyExists    int64 = 1
)

// "return {KEYS[1],KEYS[2],ARGV[1],ARGV[2]}"

const (
	hGetAllLua = `
local exists= redis.call("EXISTS",KEYS[1])
if(exists == 0)
then
    return {0,""}
end

local value = redis.call("HGETALL",KEYS[1])
return {1,value}
`
	hGetLua = `
local exists= redis.call("EXISTS",KEYS[1])
if(exists == 0)
then
    return {0,""}
end

local value = redis.call("HGET",KEYS[1],KEYS[2])
return {1,value}
`
	hSetLua = `
local exists= redis.call("EXISTS",KEYS[1])
if(exists == 0)
then
    return {0,""}
end

local result = redis.call("HSET",KEYS[1],KEYS[2],ARGV[1])
return {1,result}
`
	hDelLua = `
local exists= redis.call("EXISTS",KEYS[1])
if(exists == 0)
then
    return {0,""}
end

local result = redis.call("HDEL",KEYS[1],KEYS[2])
return {1,result}
`
)

var (
	ErrKeyNotExists   = errors.New("error key not exist")
	ErrFieldNotExists = errors.New("error field not exist")
)

// getHashAll 获取hash的所有键值，如果这个hash不存在，放回ErrKeyNotExists错误
func getHashAll(key string) (map[string][]byte, error) {
	result, err := db.RedisCli.Eval(hGetAllLua, []string{key}).Result()
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	results := result.([]interface{})

	if results[0].(int64) == KeyNotExists {
		return nil, ErrKeyNotExists
	}

	values := results[1].([]interface{})

	var m = make(map[string][]byte, len(values)/2)
	for i := range values {
		if i%2 == 1 {
			m[values[i-1].(string)] = util.Str2bytes(values[i].(string))
		}
	}

	return m, nil
}

// getHash 从指定key的hash获取指定field的值，假如这个key存在
func getHash(key string, field string, value interface{}) error {
	result, err := db.RedisCli.Eval(hGetLua, []string{key, field}).Result()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	results := result.([]interface{})

	if results[0].(int64) == KeyNotExists {
		return ErrKeyNotExists
	}

	if results[1] == nil {
		return ErrFieldNotExists
	}

	err = jsoniter.Unmarshal(util.Str2bytes(results[1].(string)), value)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// setHash 将哈希表key中的字段field的值设为value,假如key的字段存在
func setHash(key string, field string, value interface{}) error {
	bytes, err := jsoniter.Marshal(value)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	result, err := db.RedisCli.Eval(hSetLua, []string{key, field}, bytes).Result()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	results := result.([]interface{})

	if results[0].(int64) == KeyNotExists {
		return ErrKeyNotExists
	}
	return nil
}

// delHasg 删除一个字段，假如这个key存在
func delHash(key string, field string) error {
	result, err := db.RedisCli.Eval(hDelLua, []string{key, field}).Result()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	results := result.([]interface{})

	if results[0].(int64) == KeyNotExists {
		return ErrKeyNotExists
	}
	return nil
}
