package transfer

import (
	"goim/public/imerror"
	"goim/public/logger"
	"goim/public/pb"

	"github.com/golang/protobuf/proto"
)

// SignIn 设备登录
type SignInReq struct {
	Bytes []byte
}

//  SignInACK 设备登录回执
type SignInResp struct {
	ConnectStatus int    // 连接状态
	Bytes         []byte // 设备登录响应消息包
}

// NewSignInResp 创建NewSignInResp
func NewSignInResp(code int32, message string) *SignInResp {
	pbResp := pb.SignInResp{
		Code:    code,
		Message: message,
	}
	bytes, err := proto.Marshal(&pbResp)
	if err != nil {
		logger.Sugar.Error(err)
	}
	connectStatus := ConnectStatusBreak
	if code == imerror.CodeSuccess {
		connectStatus = ConnectStatusOK
	}
	return &SignInResp{
		ConnectStatus: connectStatus,
		Bytes:         bytes,
	}
}

// ErrorToSignInResp 将error转化成SignInResp
func ErrorToSignInResp(err error) *SignInResp {
	if err != nil {
		e, ok := err.(*imerror.Error)
		if ok {
			return NewSignInResp(e.Code, e.Message)
		} else {
			return NewSignInResp(imerror.ErrUnknown.Code, imerror.ErrUnknown.Message)
		}
	}
	return NewSignInResp(imerror.CodeSuccess, imerror.MessageSuccess)
}
