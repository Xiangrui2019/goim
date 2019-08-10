package cache

import (
	"goim/logic/model"
	"goim/public/imctx"
	"goim/public/logger"
	"strconv"
	"time"
)

const (
	deviceTokenKey    = "device_token:"
	deviceTokenExpire = 30 * 24 * time.Hour // 过期时间：30天
)

type deviceTokenCache struct{}

var DeviceTokenCache = new(deviceTokenCache)

func (*deviceTokenCache) Key(appId, deviceId int64) string {
	return deviceTokenKey + strconv.FormatInt(appId, 10) + ":" + strconv.FormatInt(deviceId, 10)
}

// Set 设置设备token
func (c *deviceTokenCache) Set(ctx *imctx.Context, appId, deviceId, userId int64, token string) error {
	err := set(c.Key(appId, deviceId), model.DeviceToken{
		UserId: userId,
		Token:  token,
	}, deviceTokenExpire)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// Get 获取设备token
func (c *deviceTokenCache) Get(ctx *imctx.Context, appId, deviceId int64) (int64, string, error) {
	deviceToken := model.DeviceToken{}
	err := get(c.Key(appId, deviceId), &deviceToken)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, "", err
	}
	return deviceToken.UserId, deviceToken.Token, nil
}
