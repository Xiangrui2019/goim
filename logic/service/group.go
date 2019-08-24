package service

import (
	"goim/logic/dao"
	"goim/public/imctx"
	"goim/public/logger"
)

type groupService struct{}

var GroupService = new(groupService)

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
