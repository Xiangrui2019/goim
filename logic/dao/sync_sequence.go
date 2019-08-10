package dao

import (
	"goim/public/imctx"
	"goim/public/logger"
)

type syncSequenceDao struct{}

var SyncSequenceDao = new(syncSequenceDao)

// Add 添加设备同步序列号记录
func (*syncSequenceDao) Add(ctx *imctx.Context, appId, deviceId int64, syncSequence int64) error {
	_, err := ctx.Session.Exec("insert into sync_sequence(app_id,device_id,sync_sequence) values(?,?,?)",
		appId, deviceId, syncSequence)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// Get 获取设备同步序列号
func (*syncSequenceDao) Get(ctx *imctx.Context, appId, deviceId int64) (int64, error) {
	row := ctx.Session.QueryRow("select sync_sequence from device_sync_sequence where app_id = ? and device_id = ?", appId, deviceId)
	var syncSequence int64
	err := row.Scan(&syncSequence)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}
	return syncSequence, nil
}

// GetMaxSyncSequenceByUserId 获取用户最大的同步序列号
func (*syncSequenceDao) GetMaxSyncSequenceByUserId(ctx *imctx.Context, appId, userId int64) (int64, error) {
	row := ctx.Session.QueryRow(`
		select max(s.sync_sequence) 
		from device d
		left join sync_sequence s on d.app_id = s.app_id and d.device_id = s.device_id  
		where app_id = ? and user_id = ?`, appId, userId)
	var syncSequence int64
	err := row.Scan(&syncSequence)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}
	return syncSequence, nil
}

// UpdateSyncSequence 更新设备同步序列号
func (*syncSequenceDao) UpdateSyncSequence(ctx *imctx.Context, appId, deviceId, sequence int64) error {
	_, err := ctx.Session.Exec("update sync_sequence set sync_sequence = ? where app_id = and device_id = ?",
		sequence, appId, deviceId)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}
