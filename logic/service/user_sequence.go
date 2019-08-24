package service

import (
	"goim/logic/dao"
	"goim/public/imctx"
	"goim/public/logger"
)

type userRequenceService struct{}

var UserRequenceService = new(userRequenceService)

// GetNext 获取下一个序列
func (*userRequenceService) GetNext(ctx *imctx.Context, appId, userId int64) (int64, error) {
	err := ctx.Session.Begin()
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}
	defer func() {
		err = ctx.Session.Rollback()
		if err != nil {
			logger.Sugar.Error(err)
			return
		}
	}()

	err = dao.SequenceDao.Increase(ctx, appId, userId)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}

	sequence, err := dao.SequenceDao.GetSequence(ctx, appId, userId)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}

	err = ctx.Session.Commit()
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}
	return sequence, nil
}
