package service

import (
	"goim/logic/cache"
	"goim/logic/dao"
	"goim/logic/model"
	"goim/public/imctx"
	"goim/public/imerror"
	"goim/public/logger"
)

type userService struct{}

var UserService = new(userService)

// AddUser 添加用户（将业务账号导入IM系统账户）
//1.添加用户，2.添加用户消息序列号
func (*userService) AddUser(ctx *imctx.Context, deviceId int64, user model.User) error {
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

	affected, err := dao.UserDao.Add(ctx, user)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	if affected == 0 {
		return imerror.LErrNumberUsed
	}

	err = dao.SequenceDao.Add(ctx, user.AppId, user.UserId, 0)
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

// SignIn 登录
func (*userService) SignIn(ctx *imctx.Context, appId int64, deviceId int64, token string, userId int64, secretKey string) error {
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

// VerifySecretKey 对用户秘钥进行校验
func VerifySecretKey(appid int64, userId int64, secretKey string) bool {
	return true
}

// Get 获取用户信息
func (*userService) Get(ctx *imctx.Context, appId, userId int64) (*model.User, error) {
	user, err := dao.UserDao.Get(ctx, appId, userId)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	user.Id = userId
	return user, err
}
