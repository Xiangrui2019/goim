package service

import (
	"goim/logic/dao"
	"goim/logic/model"
	"goim/public/imctx"
	"goim/public/logger"
)

type groupUserService struct{}

var GroupUserService = new(groupUserService)

// ListByUserId 获取用户的群组
func (*groupUserService) ListByUserId(ctx *imctx.Context, appId, userId int64) ([]model.Group, error) {
	groups, err := dao.GroupUserDao.ListByUserId(ctx, appId, userId)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	return groups, nil
}

// AddUser 给群组添加用户
func (*groupUserService) AddUsers(ctx *imctx.Context, groupId int64, userIds []int64) error {
	err := ctx.Session.Begin()
	if err != nil {
		logger.Sugar.Error(err)
		return nil
	}
	defer ctx.Session.Rollback()

	for _, userId := range userIds {
		err := dao.GroupUserDao.Add(ctx, appId, groupId, userId)
		if err != nil {
			logger.Sugar.Error(err)
			return err
		}
	}
	ctx.Session.Commit()
	return nil
}

// DeleteUser 从群组移除用户
func (*groupUserService) DeleteUser(ctx *imctx.Context, groupId int64, userIds []int64) error {
	err := ctx.Session.Begin()
	if err != nil {
		logger.Sugar.Error(err)
		return nil
	}
	defer ctx.Session.Rollback()

	for _, userId := range userIds {
		err := dao.GroupUserDao.Delete(ctx, groupId, userId)
		if err != nil {
			logger.Sugar.Error(err)
			return err
		}
	}
	ctx.Session.Commit()
	return nil
}

func (*groupUserService) UpdateLabel(ctx *imctx.Context, groupId int64, userId int64, label string) error {
	return dao.GroupUserDao.UpdateLabel(ctx, groupId, userId, label)
}
