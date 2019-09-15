package dao

import (
	"fmt"
	"goim/logic/model"
	"goim/public/imctx"
	"goim/public/logger"
)

type messageDao struct{}

var MessageDao = new(messageDao)

// Add 插入一条消息
func (*messageDao) Add(ctx *imctx.Context, tableName string, message model.Message) error {
	sql := fmt.Sprintf(`insert into %s(message_id,app_id,user_id,sender_type,sender_id,sender_device_id,receiver_type,receiver_id,to_user_ids,message_body_id,user_seq,send_time) values(?,?,?,?,?,?,?,?,?,?,?,?)`, tableName)
	_, err := ctx.Session.Exec(sql, message.MessageId, message.AppId, message.UserId, message.SenderType, message.SenderId, message.SenderDeviceId,
		message.ReceiverType, message.ReceiverId, message.ToUserIds, message.MessageBodyId, message.UserSeq, message.SendTime)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// ListByUserIdAndUserSeq 根据用户id查询大于序号大于sequence的消息
func (*messageDao) ListByUserIdAndUserSeq(ctx *imctx.Context, tableName string, appId, userId, userSeq int64) ([]model.Message, error) {
	sql := fmt.Sprintf(`select message_id,app_id,user_id,sender_type,sender_id,sender_device_id,receiver_type,receiver_id,to_user_ids,message_body_id,user_seq,send_time from %s where app_id = ? and user_id = ? and user_seq > ?`, tableName)
	rows, err := ctx.Session.Query(sql, appId, userId, userSeq)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	messages := make([]*model.Message, 0, 5)
	for rows.Next() {
		message := new(model.Message)
		err := rows.Scan(&message.MessageId, &message.AppId, &message.UserId, &message.SenderType, &message.SenderId, &message.SenderDeviceId, &message.ReceiverType,
			&message.ReceiverId, &message.ToUserIds, &message.MessageBodyId, &message.UserSeq, &message.SendTime)
		if err != nil {
			logger.Sugar.Error(err)
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
