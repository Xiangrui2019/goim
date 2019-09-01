package dao

import (
	"goim/public/imctx"
	"goim/public/logger"
)

type deviceAckDao struct{}

var DeviceAckDao = new(deviceAckDao)

// Add 添加设备同步序列号记录
func (*deviceAckDao) Add(ctx *imctx.Context, appId, deviceId int64, ack int64) error {
	_, err := ctx.Session.Exec("insert into device_ack(app_id,device_id,ack) values(?,?,?)",
		appId, deviceId, ack)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// Get 获取设备同步序列号
func (*deviceAckDao) Get(ctx *imctx.Context, appId, deviceId int64) (int64, error) {
	row := ctx.Session.QueryRow("select ack from device_ack where app_id = ? and device_id = ?", appId, deviceId)
	var syncSequence int64
	err := row.Scan(&syncSequence)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}
	return syncSequence, nil
}

// UpdateSyncSequence 更新设备同步序列号
func (*deviceAckDao) Update(ctx *imctx.Context, appId, deviceId, ack int64) error {
	_, err := ctx.Session.Exec("update device_ack set ack = ? where app_id = ? and device_id = ?",
		ack, appId, deviceId)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// GetMaxByUserId 获取用户最大的同步序列号
func (*deviceAckDao) GetMaxByUserId(ctx *imctx.Context, appId, userId int64) (int64, error) {
	row := ctx.Session.QueryRow(`
		select max(a.ack) 
		from device d
		left join device_ack a on d.app_id = a.app_id and d.device_id = a.device_id  
		where d.app_id = ? and d.user_id = ?`, appId, userId)
	var syncSequence int64
	err := row.Scan(&syncSequence)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}
	return syncSequence, nil
}
