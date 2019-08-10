package service

import (
	"goim/logic/dao"
	"goim/logic/model"
	"goim/public/imctx"
	"goim/public/logger"
)

type groupService struct{}

var GroupService = new(groupService)

// ListGroupUser 获取群组的用户信息
func (*groupService) GetUsers(ctx *imctx.Context, appId, groupId int64) ([]model.GroupUserInfo, error) {
	userInfos, err := dao.GroupUserDao.ListGroupUser(ctx, appId, groupId)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	return userInfos, err
}

// CreateAndAddUser 创建群组并且添加群成员
func (*groupService) CreateAndAddUser(ctx *imctx.Context, groupName string, userIds []int64) (int64, error) {
	err := ctx.Session.Begin()
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}
	defer ctx.Session.Rollback()

	id, err := dao.GroupDao.Add(ctx, groupName)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}

	for _, userId := range userIds {
		err := dao.GroupUserDao.Add(ctx, id, userId)
		if err != nil {
			logger.Sugar.Error(err)
			return 0, err
		}
	}
	ctx.Session.Commit()
	return id, nil
}
