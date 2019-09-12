package transfer

import (
	"goim/public/imerror"
	"goim/public/logger"
	"goim/public/pb"

	"github.com/golang/protobuf/proto"
)

const (
	CodeSignInSuccess = 1
	CodeSignInFail    = 2
)

// SignIn 设备登录
type SignInReq struct {
	Bytes []byte
}

//  SignInACK 设备登录回执
type SignInResp struct {
	Result   bool
	DeviceId int64
	UserId   int64
	Bytes    []byte
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
	return &SignInResp{
		Result: code == imerror.CodeSuccess,
		Bytes:  bytes,
	}
}

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
