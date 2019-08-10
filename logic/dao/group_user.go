package dao

import (
	"database/sql"
	"goim/logic/model"
	"goim/public/imctx"
	"goim/public/logger"
)

type groupUserDao struct{}

var GroupUserDao = new(groupUserDao)

// ListGroupUser 获取群组用户信息
func (*groupUserDao) ListGroupUser(ctx *imctx.Context, appId, groupId int64) ([]model.GroupUser, error) {
	rows, err := ctx.Session.Query(`
		select user_id,label,extra,create_time,update_time 
		from group_user
		where app_id = ? and group_id = ?`, appId, groupId)
	if err != nil {
		return nil, err
	}
	groupUsers := make([]model.GroupUser, 0, 5)
	for rows.Next() {
		var groupUser = model.GroupUser{
			AppId:   appId,
			GroupId: groupId,
		}
		err := rows.Scan(&groupUser.UserId, &groupUser.Label, &groupUser.Extra, &groupUser.CreateTime, &groupUser.UpdateTime)
		if err != nil {
			logger.Sugar.Error(err)
			return nil, err
		}
		groupUsers = append(groupUsers, groupUser)
	}
	return groupUsers, nil
}

// ListByUser 获取用户加入的群组信息
func (*groupUserDao) ListByUserId(ctx *imctx.Context, appId, userId int64) ([]model.Group, error) {
	rows, err := ctx.Session.Query(`
		select group_id,name,introduction,user_num,type,extra,create_time,update_time  
		from group_user u 
		left join group g u.app_id = g.app_id and u.group_id = g.group_id
		where app_id = ? and user_id = ?`,
		appId, userId)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	var groups []model.Group
	var group model.Group
	for rows.Next() {
		err := rows.Scan(&group.GroupId, &group.Name, &group.Introduction, &group.UserNum, &group.Type, &group.Extra, &group.CreateTime, &group.UpdateTime)
		if err != nil {
			logger.Sugar.Error(err)
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

// Add 将用户添加到群组
func (*groupUserDao) Add(ctx *imctx.Context, appId, groupId, userId int64, lable, extra string) error {
	_, err := ctx.Session.Exec("insert ignore into group_user(app_id,group_id,user_id,lable,extra) values(?,?,?,?,?)",
		appId, groupId, userId, lable, extra)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// Delete 将用户从群组删除
func (d *groupUserDao) Delete(ctx *imctx.Context, appId int64, groupId int64, userId int64) error {
	_, err := ctx.Session.Exec("delete from group_user where app_id = ? and group_id = ? and user_id = ?",
		appId, groupId, userId)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// UpdateLabel 更新用户群组信息
func (*groupUserDao) Update(ctx *imctx.Context, appId, groupId, userId int64, label string, extra string) error {
	_, err := ctx.Session.Exec("update group_user set label = ?,extra = ? where app_id = ? and group_id = ? and user_id = ?",
		label, extra, appId, groupId, userId)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// GetGroupUser 获取群组用户信息,用户不存在返回nil
func (*groupUserDao) GetGroupUser(ctx *imctx.Context, appId, groupId, userId int64) (*model.GroupUser, error) {
	var groupUser = model.GroupUser{
		AppId:   appId,
		GroupId: groupId,
		UserId:  userId,
	}
	err := ctx.Session.QueryRow("select lable,extra from group_user where app_id = ? and group_id = ? and user_id = ?",
		appId, groupId, userId).
		Scan(&groupUser.Label, &groupUser.Extra)
	if err != nil && err != sql.ErrNoRows {
		logger.Sugar.Error(err)
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &groupUser, nil
}
