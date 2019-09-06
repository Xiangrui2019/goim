package dao

import (
	"fmt"
	"goim/logic/model"
	"goim/public/imctx"
	"goim/public/logger"
	"strconv"
	"strings"
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
func (*messageBodyDao) List(ctx *imctx.Context, tableName string, messageBodyIds []int64) ([]*model.MessageBody, error) {
	var build strings.Builder
	for i := range messageBodyIds {
		build.WriteString(strconv.FormatInt(messageBodyIds[i], 10))
		if i != len(messageBodyIds)-1 {
			build.WriteString(",")
		}
	}

	sql := fmt.Sprintf(`select message_body_id,content from %s where message_body_id in (%s)`, tableName, build.String())
	rows, err := ctx.Session.Query(sql)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	bodys := make([]*model.MessageBody, 0, len(messageBodyIds))
	for rows.Next() {
		body := new(model.MessageBody)
		err := rows.Scan(&body.MessageBodyId, &body.Content)
		if err != nil {
			logger.Sugar.Error(err)
			return nil, err
		}
		bodys = append(bodys, body)
	}
	return bodys, nil
}
