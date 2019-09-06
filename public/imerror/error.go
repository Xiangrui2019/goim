package imerror

import "goim/public/pb"

var (
	CCodeSuccess = 0 // 成功发送
)

// Error 接入层调用错误
type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

var (
	ErrUnknown         = NewError(int(pb.ErrCode_SERVER_UNKNOWN), "error unknown error") // 服务器位置错误
	ErrDeviceIdOrToken = NewError(1001, "error device token")                            // 设备id或者token错误
	ErrNotIsFriend     = NewError(2, "error not is friend")                              // 非好友关系
	ErrNotInGroup      = NewError(3, "error not in group")                               // 没有在群组内
)
