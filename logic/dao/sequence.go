package dao

import (
	"goim/public/imctx"
	"goim/public/logger"
)

type sequenceDao struct{}

var SequenceDao = new(sequenceDao)

// Add 添加 todo redis持久层
func (*sequenceDao) Add(ctx *imctx.Context, appId, userId int64, sequence int64) error {
	_, err := ctx.Session.Exec("insert into user_sequence (app_id,user_id,sequence) values(?,?,?)", appId, userId, sequence)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// Increase sequence++
func (*sequenceDao) Increase(ctx *imctx.Context, appId int64, userId int64) error {
	_, err := ctx.Session.Exec("update user_sequence set sequence = sequence + 1 where app_id = ? and user_id = ?", appId, userId)
	if err != nil {
		logger.Sugar.Error(err)
	}
	return err
}

// GetSequence 获取自增序列
func (*sequenceDao) GetSequence(ctx *imctx.Context, appId, userId int64) (int64, error) {
	var sequence int64
	err := ctx.Session.QueryRow("select sequence from user_sequence where app_id = ? and user_id = ?", appId, userId).
		Scan(&sequence)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err

	}
	return sequence, nil
}
