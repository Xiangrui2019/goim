package model

import (
	"goim/public/pb"
	"time"
)

// Message 消息
type Message struct {
	Id             int64     // 自增主键
	MessageId      string    // 消息id
	AppId          int64     // appId
	UserId         int64     // 用户id
	SenderType     int32     // 发送者类型
	SenderId       int64     // 发送者账户id
	SenderDeviceId int64     // 发送者设备id
	ReceiverType   int32     // 接收者账户id
	ReceiverId     int64     // 接收者id,如果是单聊信息，则为user_id，如果是群组消息，则为group_id
	ToUserIds      string    // 需要@的用户id列表，多个用户用，隔开
	MessageBodyId  int64     // 消息体id
	MessageBody    string    `gorm:"-"` // 消息体
	UserSeq        int64     // 消息同步序列
	SendTime       time.Time // 消息发送时间
	Status         int32     // 创建时间
}

func MessageToPb(message *Message) *pb.MessageItem {
	return &pb.MessageItem{
		MessageId:      message.MessageId,
		SenderType:     message.SenderType,
		SenderId:       message.SenderId,
		SenderDeviceId: message.SenderDeviceId,
		ReceiverType:   message.ReceiverType,
		ReceiverId:     message.ReceiverId,
		ToUserIds:      []int64{},
		MessageBody:    nil,
		UserSeq:        message.UserSeq,
		SendTime:       message.SendTime.Unix(),
		Status:         message.Status,
	}
}
