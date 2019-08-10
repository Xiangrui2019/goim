package db

import (
	"goim/conf"

	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
)

var RedisCli *redis.Client

const DeviceIdPre = "connect:device_id:"

func init() {
	addr := conf.RedisIP

	RedisCli = redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	_, err := RedisCli.Ping().Result()
	if err != nil {
		logs.Error("redis err ")
		panic(err)
	}
}
