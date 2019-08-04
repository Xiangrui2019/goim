package dao

import (
	"goim/logic/model"
	"goim/public/imctx"
	"goim/public/logger"
)

type groupDao struct{}

var GroupDao = new(groupDao)

// Get 获取群组信息
func (*groupDao) Get(ctx *imctx.Context, id int) (*model.Group, error) {
	row := ctx.Session.QueryRow("select name from `group` where id = ?", id)
	group := new(model.Group)
	err := row.Scan(&group.Name)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	return group, nil
}

// Insert 插入一条群组
func (*groupDao) Add(ctx *imctx.Context, appId int64, groupId int64, name string) (int64, error) {
	result, err := ctx.Session.Exec("insert into `group`(name) value(?)", name)
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
