package service

import (
	"goim/logic/dao"
	"goim/public/imctx"
	"goim/public/imerror"
	"goim/public/logger"
)

type authService struct{}

var AuthService = new(authService)

// SignIn 长连接登录
func (*authService) SignIn(ctx *imctx.Context, appId, userId, deviceId int64, token string) error {
	// 用户验证
	if !VerifyToken(appId, appId, deviceId, token) {
		return imerror.LErrToken
	}

	// 设备验证
	// todo 检查设备是否存在
	// todo 检查用户是否存在

	err := dao.DeviceDao.UpdateUserId(ctx, appId, deviceId, userId)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// Auth 验证用户是否登录
func (*authService) Auth(ctx *imctx.Context, appId, userId, deviceId int64, token string) error {
	// 进行一次鉴权
	return nil
}

// VerifySecretKey 对用户秘钥进行校验
func VerifyToken(appid, userId, diviceId int64, token string) bool {
	return true
}
