package dao

import (
	"goim/public/imctx"
	"goim/public/logger"
)

type uesrSeqDao struct{}

var UserSeqDao = new(uesrSeqDao)

// Add 添加 todo redis持久层
func (*uesrSeqDao) Add(ctx *imctx.Context, appId, userId int64, seq int64) error {
	_, err := ctx.Session.Exec("insert into user_seq (app_id,user_id,seq) values(?,?,?)", appId, userId, seq)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// Increase sequence++
func (*uesrSeqDao) Increase(ctx *imctx.Context, appId int64, userId int64) error {
	_, err := ctx.Session.Exec("update user_seq set seq = seq + 1 where app_id = ? and user_id = ?", appId, userId)
	if err != nil {
		logger.Sugar.Error(err)
	}
	return err
}

// GetSequence 获取自增序列
func (*uesrSeqDao) Get(ctx *imctx.Context, appId, userId int64) (int64, error) {
	var sequence int64
	err := ctx.Session.QueryRow("select seq from user_seq where app_id = ? and user_id = ?", appId, userId).
		Scan(&sequence)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err

	}
	return sequence, nil
}
