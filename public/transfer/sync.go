package transfer

import (
	"github.com/golang/protobuf/proto"
	"goim/public/imerror"
	"goim/public/pb"
)

// 同步消息触发
type SyncReq struct {
	DeviceId int64  `json:"device_id"`     // 设备id
	UserId   int64  `json:"user_id"`       // 用户id
	Bytes    []byte `json:"sync_sequence"` // 已经同步的消息序列号
}

// 同步消息触发
type SyncResp struct {
	Code    int32
	Message string
	Bytes   []byte
}

// NewSyncResp 创建NewSyncResp
func NewSyncResp(code int32, message string, messages []*pb.MessageItem) *SyncResp {
	syncResp := pb.SyncResp{
		Code:     code,
		Message:  "",
		Messages: messages,
	}

	bytes, err := proto.Marshal(&syncResp)
	if err != nil
	return &SyncResp{
		Code:    code,
		Message: message,
		Bytes:   bytes,
	}
}

func ErrorToSyncResp(err error, messages []*pb.MessageItem) *SyncResp {
	if err != nil {
		e, ok := err.(*imerror.Error)
		if ok {
			return NewSyncResp(e.Code, e.Message, nil)
		} else {
			return NewSyncResp(imerror.ErrUnknown.Code, imerror.ErrUnknown.Message, nil)
		}
	}
	return NewSyncResp(imerror.CodeSuccess, imerror.MessageSuccess, messages)
}
