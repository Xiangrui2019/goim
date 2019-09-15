package transfer

import (
	"goim/public/pb"
)

type MessageAckReq struct {
	IsSignIn bool // 标记用户是否登录成功
	AppId    int64
	DeviceId int64  // 设备id
	UserId   int64  // 用户id
	Bytes    []byte // 字节数组
}

type MessageAckResp struct {
	ConnectStatus int // 连接状态
}

// NewSyncResp 创建MessageAckResp
func NewMessageAckResp(code int32, message string, messages []*pb.MessageItem) *SyncResp {
	return nil
}

// ErrorMessageAckResp 将err转化为MessageAckResp
func ErrorMessageAckResp(err error, messages []*pb.MessageItem) *SyncResp {
	return nil
}
