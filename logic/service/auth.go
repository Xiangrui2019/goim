package service

import (
	"goim/logic/cache"
	"goim/logic/dao"
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

// SignIn 登录
func (*authService) SignIn(ctx *imctx.Context, appId int64, deviceId int64, token string, userId int64, secretKey string) error {
	err := ctx.Session.Begin()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	defer func() {
		err = ctx.Session.Rollback()
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()

	// 用户验证
	if !VerifySecretKey(appId, userId, secretKey) {
		return imerror.LErrSecretKey
	}

	// 设备验证
	device, err := dao.DeviceDao.Get(ctx, appId, deviceId)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	if device == nil {
		return imerror.LErrDeviceNotFound
	}

	if device.Token != token {
		return imerror.LErrToken
	}

	user, err := dao.UserDao.Get(ctx, appId, userId)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	if user == nil {
		return imerror.LErrUserNotFound
	}

	err = dao.DeviceDao.UpdateUserId(ctx, appId, deviceId, userId)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	err = cache.DeviceTokenCache.Set(ctx, appId, deviceId, userId, token)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	err = ctx.Session.Commit()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	return nil
}
