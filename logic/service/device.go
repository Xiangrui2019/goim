package service

import (
	"goim/logic/cache"
	"goim/logic/dao"
	"goim/logic/model"

	"goim/public/imctx"
	"goim/public/logger"

	"github.com/satori/go.uuid"
)

const (
	DeviceOnline  = 1
	DeviceOffline = 0
)

type deviceService struct{}

var DeviceService = new(deviceService)

// Regist 注册设备
func (*deviceService) Regist(ctx *imctx.Context, device model.Device) (int64, string, error) {
	err := ctx.Session.Begin()
	if err != nil {
		logger.Sugar.Error(err)
		return 0, "", err
	}
	defer ctx.Session.Rollback()

	UUID := uuid.NewV4()
	device.Token = UUID.String()
	id, err := dao.DeviceDao.Add(ctx, device)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, "", err
	}

	err = dao.SyncSequenceDao.Add(ctx, device.AppId, device.Id, 0)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, "", err
	}

	err = ctx.Session.Commit()
	if err != nil {
		logger.Sugar.Error(err)
		return 0, "", err
	}
	return id, device.Token, nil
}

// ListOnlineByUserId 获取用户的所有在线设备
func (*deviceService) ListOnlineByUserId(ctx *imctx.Context, appId, userId int64) ([]model.Device, error) {
	devices, err := cache.DeviceCache.GetAll(appId, userId)
	if err != nil && err != cache.ErrKeyNotExists {
		logger.Sugar.Error(err)
		return nil, err
	}

	if err != cache.ErrKeyNotExists {
		return devices, nil
	}

	devices, err = dao.DeviceDao.ListOnlineByUserId(ctx, appId, userId)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

// ListOnlineByUserId 获取用户的所有在线设备
func (*deviceService) AddToUser(ctx *imctx.Context, appId, userId int64) ([]model.Device, error) {
	devices, err := cache.DeviceCache.GetAll(appId, userId)
	if err != nil && err != cache.ErrKeyNotExists {
		logger.Sugar.Error(err)
		return nil, err
	}

	if err != cache.ErrKeyNotExists {
		return devices, nil
	}

	devices, err = dao.DeviceDao.ListOnlineByUserId(ctx, appId, userId)
	if err != nil {
		return nil, err
	}
	return devices, nil
}
