package cache

import (
	"goim/logic/db"
	"goim/logic/model"
	"goim/public/logger"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
)

const (
	DeviceKey    = "device:"
	DeviceExpire = 2 * time.Hour
)

type deviceCache struct{}

var DeviceCache = new(deviceCache)

func (c *deviceCache) Key(appId, userId int64) string {
	return DeviceKey + strconv.FormatInt(appId, 10) + ":" + strconv.FormatInt(userId, 10)
}

// SetAll 将指定用户的所有在线设备存入缓存
func (c *deviceCache) SetAll(appId, userId, deviceId int64, devices []model.Device) error {
	deviceMap := make(map[string]interface{}, len(devices)+1)
	for _, device := range devices {
		bytes, err := jsoniter.Marshal(device)
		if err != nil {
			logger.Sugar.Error(err)
			return err
		}

		deviceMap[strconv.FormatInt(device.UserId, 10)] = bytes
	}

	key := c.Key(appId, userId)
	err := db.RedisCli.HMSet(key, deviceMap).Err()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	err = db.RedisCli.Expire(key, DeviceExpire).Err()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// GetAll 获取指定用户的所有在线设备
func (c *deviceCache) GetAll(appId, userId int64) ([]model.Device, error) {
	result, err := getHashAll(c.Key(appId, userId))
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	devices := make([]model.Device, 0, len(result))

	for _, v := range result {
		var device model.Device
		err = jsoniter.Unmarshal(v, &device)
		if err != nil {
			logger.Sugar.Error(err)
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}

// Get 获取某一个用户的在线设备
func (c *deviceCache) Get(appId, userId, deviceId int64) (*model.Device, error) {
	var device model.Device
	err := getHash(c.Key(appId, deviceId), strconv.FormatInt(userId, 10), &device)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	return &device, nil
}

// Set 设置某一用户的在线设备
func (c *deviceCache) Set(appId, userId, deviceId int64, device model.Device) error {
	err := setHash(c.Key(appId, userId), strconv.FormatInt(deviceId, 10), device)
	if err != nil && err != ErrKeyNotExists {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// Del 删除某一用户的在线设备
func (c *deviceCache) Del(appId, userId, deviceId int64) error {
	err := delHash(c.Key(appId, deviceId), strconv.FormatInt(userId, 10))
	if err != nil && err != ErrKeyNotExists {
		logger.Sugar.Error(err)
		return err
	}

	return nil
}
