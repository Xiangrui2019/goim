package dao

import (
	"database/sql"
	"goim/logic/model"
	"goim/public/imctx"
	"goim/public/logger"
)

type deviceDao struct{}

var DeviceDao = new(deviceDao)

// Insert 插入一条设备信息
func (*deviceDao) Add(ctx *imctx.Context, device model.Device) (int64, error) {
	result, err := ctx.Session.Exec(`insert into device(app_id,device_id,user_id,token,type,brand,model,system_version,sdk_version,status) 
		values(?,?,?,?,?,?,?,?.?,?)`,
		device.AppId, device.DeviceId, device.UserId, device.Token, device.Type, device.Brand, device.Model, device.SystemVersion, device.SDKVersion)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}
	return id, nil
}

// Get 获取设备
func (*deviceDao) Get(ctx *imctx.Context, appId, deviceId int64) (*model.Device, error) {
	device := model.Device{
		AppId:    appId,
		DeviceId: deviceId,
	}
	row := ctx.Session.QueryRow(`
		select user_id,token,type,brand,model,system_version,sdk_version,status,create_time,update_time
		from device where app_id = ? and device_id = ?`, appId, deviceId)
	err := row.Scan(&device.UserId, &device.Token, &device.Type, &device.Brand, &device.Model, &device.SystemVersion, &device.SDKVersion,
		&device.Status, &device.CreateTime, &device.UpdateTime)
	if err != nil && err != sql.ErrNoRows {
		logger.Sugar.Error(err)
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &device, err
}

// ListUserOnline 查询用户所有的在线设备
func (*deviceDao) ListOnlineByUserId(ctx *imctx.Context, appId, userId int64) ([]*model.Device, error) {
	rows, err := ctx.Session.Query("select user_id,token,type,brand,model,system_version,sdk_version,status,create_time,update_time from device where app_id = ? and user_id = ? and status = ?",
		appId, userId, model.DeviceOnLine)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	devices := make([]*model.Device, 0, 5)
	for rows.Next() {
		device := new(model.Device)
		err = rows.Scan(&device.UserId, &device.Token, &device.Type, &device.Brand, &device.Model, &device.SystemVersion, &device.SDKVersion,
			&device.Status, &device.CreateTime, &device.UpdateTime)
		if err != nil {
			logger.Sugar.Error(err)
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}

// UpdateUserId 更新设备绑定用户
func (*deviceDao) UpdateUserId(ctx *imctx.Context, appId, deviceId, userId int64) error {
	_, err := ctx.Session.Exec("update device set user_id = ? where app_id = ? and device_id = ? ", userId, appId, deviceId)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// UpdateStatus 更新设备的在线状态
func (*deviceDao) UpdateStatus(ctx *imctx.Context, appId, deviceId int64, status int) error {
	_, err := ctx.Session.Exec("update device set status = ? where app_id = ? and id = ? ", status, appId, deviceId)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// Upgrade 升级设备
func (*deviceDao) Upgrade(ctx *imctx.Context, appId, deviceId int64, systemVersion, sdkVersion int) error {
	_, err := ctx.Session.Exec("update device set system_version = ?,sdk_version = ? where app_id = ? and id = ? ", systemVersion, sdkVersion, appId, deviceId)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}
