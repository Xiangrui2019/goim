package dao

import (
	"database/sql"
	"fmt"
	"goim/logic/model"
	"goim/public/imctx"
	"goim/public/logger"
)

type messageBodyDao struct{}

var MessageBodyDao = new(messageBodyDao)

// Add 插入一条消息体
func (*messageBodyDao) Add(ctx *imctx.Context, tableName string, messageBodyId int64, content string) error {
	sql := fmt.Sprintf(`insert into %s(message_body_id,content) values(?,?)`, tableName)
	_, err := ctx.Session.Exec(sql, messageBodyId, content)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// List 根据消息体id查询消息体
func (*messageBodyDao) Get(ctx *imctx.Context, tableName string, messageBodyId int64) (*model.MessageBody, error) {
	sqlStr := fmt.Sprintf(`select message_body_id,content from %s where message_body_id = ?`, tableName)
	row := ctx.Session.QueryRow(sqlStr, messageBodyId)

	body := new(model.MessageBody)
	err := row.Scan(&body.MessageBodyId, &body.Content)
	if err != nil && err != sql.ErrNoRows {
		logger.Sugar.Error(err)
		return nil, err
	}

	return body, nil
}
