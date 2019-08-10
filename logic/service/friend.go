package service

import (
	"goim/logic/cache"
	"goim/logic/dao"
	"goim/logic/model"
	"goim/public/imctx"
	"goim/public/logger"
)

type friendService struct{}

var FriendService = new(friendService)

// List 获取用户好友列表
func (*friendService) ListUserFriend(ctx *imctx.Context, appId, userId int64) ([]model.UserFriend, error) {
	friends, err := dao.FriendDao.ListUserFriend(ctx, appId, userId)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	return friends, err
}

// Add 添加好友关系
func (*friendService) Add(ctx *imctx.Context, appId, userId int64, friendId int64, label1, label2, extra1, extra2 string) error {
	err := ctx.Session.Begin()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	defer func() {
		err = ctx.Session.Rollback()
		if err != nil {
			logger.Sugar.Error(err)
			return
		}
	}()

	err = dao.FriendDao.Add(ctx, appId, userId, friendId, label1, extra1)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	err = dao.FriendDao.Add(ctx, appId, friendId, userId, label2, extra2)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	err = ctx.Session.Commit()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// Delete 删除好友关系
func (*friendService) Delete(ctx *imctx.Context, appId, userId, friendId int64) error {
	err := ctx.Session.Begin()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	defer func() {
		err = ctx.Session.Rollback()
		if err != nil {
			logger.Sugar.Error(err)
			return
		}
	}()

	err = dao.FriendDao.Delete(ctx, appId, userId, friendId)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	err = dao.FriendDao.Delete(ctx, appId, friendId, userId)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	err = cache.FriendCache.Delele(ctx, appId, userId, friendId)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	err = cache.FriendCache.Delele(ctx, appId, friendId, userId)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	err = ctx.Session.Commit()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// Get 获取朋友信息
func (*friendService) Get(ctx *imctx.Context, appId, userId, friendId int64) (*model.Friend, error) {
	friend, err := cache.FriendCache.Get(ctx, appId, userId, friendId)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	if friend != nil {
		return friend, nil
	}

	friend, err = dao.FriendDao.Get(ctx, appId, userId, friendId)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	return friend, nil
}
