package dao

import (
	"fmt"
	"goim/logic/db"
	"goim/logic/model"
	"goim/public/imctx"
	"testing"
)

var ctx = imctx.NewContext(db.Factoty.GetSession())

func TestDeviceDao_Add(t *testing.T) {
	device := model.Device{
		AppId:         1,
		DeviceId:      1,
		UserId:        1,
		Token:         "token",
		Type:          1,
		Brand:         "huawei",
		Model:         "huawei P10",
		SystemVersion: "8.0.0",
		SDKVersion:    "1.0.0",
		Status:        1,
	}
	fmt.Println(DeviceDao.Add(ctx, device))
}

func TestDeviceDao_Get(t *testing.T) {
	fmt.Println(DeviceDao.Get(ctx, 1, 1))
}

func TestDeviceDao_ListOnlineByUserId(t *testing.T) {
	fmt.Println(DeviceDao.ListOnlineByUserId(ctx, 1, 1))
}

func TestDeviceDao_UpdateUserId(t *testing.T) {
	fmt.Println(DeviceDao.UpdateUserId(ctx, 1, 1, 2))
}

func TestDeviceDao_UpdateStatus(t *testing.T) {
	fmt.Println(DeviceDao.UpdateStatus(ctx, 1, 1, 0))
}

func TestDeviceDao_Upgrade(t *testing.T) {
	fmt.Println(DeviceDao.Upgrade(ctx, 1, 1, "9.0.0", "2.0.0"))
}
