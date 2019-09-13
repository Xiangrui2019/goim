package imerror

import (
	"goim/public/pb"
)

const (
	CodeSuccess    = 0    // code成功
	MessageSuccess = "OK" // message成功
)

// Error 接入层调用错误
type Error struct {
	Code    int32
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code int32, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

var (
	ErrUnknown         = NewError(int32(pb.ErrCode_SERVER_UNKNOWN), "error unknown error") // 服务器未知错误
	ErrUnauthorized    = NewError(int32(pb.ErrCode_UNAUTHORIZED), "error unauthorized")    // 未登录
	ErrDeviceIdOrToken = NewError(1001, "error device token")                              // 设备id或者token错误
	ErrNotIsFriend     = NewError(2, "error not is friend")                                // 非好友关系
	ErrNotInGroup      = NewError(3, "error not in group")                                 // 没有在群组内
	LErrToken          = NewLError(1003, "error secret key")                               // 用户秘钥错误
)
