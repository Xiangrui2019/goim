package cache

import (
	"goim/logic/db"
	"goim/public/logger"
	"time"

	"github.com/json-iterator/go"
)

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
