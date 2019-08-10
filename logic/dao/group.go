package dao

import (
	"goim/logic/model"
	"goim/public/imctx"
	"goim/public/logger"
)

type groupDao struct{}

var GroupDao = new(groupDao)

// Get 获取群组信息
func (*groupDao) Get(ctx *imctx.Context, appId, groupId int64) (*model.Group, error) {
	row := ctx.Session.QueryRow("select name,introduction,user_num,type,extra,create_time,update_time from `group` where app_id = ? and group_id = ?",
		appId, groupId)
	group := model.Group{
		AppId:   appId,
		GroupId: groupId,
	}
	err := row.Scan(&group.Name, &group.Introduction, &group.UserNum, &group.Type, &group.Extra, &group.CreateTime, &group.UpdateTime)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	return &group, nil
}

// Insert 插入一条群组
func (*groupDao) Add(ctx *imctx.Context, group model.Group) (int64, error) {
	result, err := ctx.Session.Exec("insert into `group`(app_id,group_id,name,introduction,type,extra) value(?,?,?,?,?,?)",
		group.AppId, group.GroupId, group.Name, group.Introduction, group.Type, group.Extra)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}
	num, err := result.RowsAffected()
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}
	return num, nil
}
