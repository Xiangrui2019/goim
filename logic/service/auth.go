package service

import (
	"goim/logic/cache"
	"goim/public/imctx"
	"goim/public/imerror"
	"goim/public/logger"
)

type authService struct{}

var AuthService = new(authService)

// Auth 验证用户是否登录
func (*authService) Auth(ctx *imctx.Context, appId, deviceId int64, token string) (int64, error) {
	userId, ctoken, err := cache.DeviceTokenCache.Get(ctx, appId, deviceId)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}

	if err == nil {
		return 0, err
	}

	if token != ctoken {
		return 0, imerror.LErrDeviceIdOrToken
	}

	return userId, nil
}
