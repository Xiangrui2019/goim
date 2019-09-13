package transfer

import (
	"goim/public/imerror"
	"goim/public/logger"
	"goim/public/pb"

	"github.com/golang/protobuf/proto"
)

type SendMessageReq struct {
	IsSignIn bool   // 标记用户是否登录成功
	DeviceId int64  // 设备id
	UserId   int64  // 用户id
	Bytes    []byte // 消息包
}

type SendMessageResp struct {
	ConnectStatus int    // 连接状态
	Bytes         []byte // 字节包
}

func NewSendMessageResp(code int32, message string, messageId string) *SendMessageResp {
	sendMessageResp := pb.SendMessageResp{
		Code:      code,
		Message:   "",
		MessageId: messageId,
	}

	bytes, err := proto.Marshal(&sendMessageResp)
	if err != nil {
		logger.Sugar.Error(err)
	}
	connectStatus := ConnectStatusBreak
	if code == imerror.CodeSuccess {
		connectStatus = ConnectStatusOK
	}
	return &SendMessageResp{
		ConnectStatus: connectStatus,
		Bytes:         bytes,
	}
}

func ErrorToSendMessageResp(err error, messageId string) *SendMessageResp {
	if err != nil {
		e, ok := err.(*imerror.Error)
		if ok {
			return NewSendMessageResp(e.Code, e.Message, messageId)
		} else {
			return NewSendMessageResp(imerror.ErrUnknown.Code, imerror.ErrUnknown.Message, messageId)
		}
	}
	return NewSendMessageResp(imerror.CodeSuccess, imerror.MessageSuccess, messageId)
}
