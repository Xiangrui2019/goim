package cache

import (
	"errors"
	"goim/logic/db"
	"goim/public/lib"
	"goim/public/logger"
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

// hget 从指定key的hash获取指定field的值，假如这个key存在
func hget(key string, field string, value interface{}) error {
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

	err = jsoniter.Unmarshal(lib.Str2bytes(results[1].(string)), value)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// hset 将哈希表key中的字段field的值设为value,假如key的字段存在
func hset(key string, field string, value interface{}) error {
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

// hdel 删除一个字段，假如这个key存在
func hdel(key string, field string) error {
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
