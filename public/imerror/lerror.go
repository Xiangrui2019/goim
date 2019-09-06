package imerror

// LError HTTP调用错误
type LError struct {
	Code    int
	Message string
	Data    interface{}
}

func (e *LError) Error() string {
	return e.Message
}

func NewLError(code int, message string) *LError {
	return &LError{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

func WrapLErrorWithData(err *LError, data interface{}) *LError {
	return &LError{
		Code:    err.Code,
		Message: err.Message,
		Data:    data,
	}
}

// 通用错误
var (
	LErrUnauthorized      = NewLError(1, "error unauthorized")         // 需要认证
	LErrBadRequest        = NewLError(2, "error bad request")          // 请求错误
	LErrUnknown           = NewLError(3, "error unknown error")        // 未知错误
	LErrDeviceNotBindUser = NewLError(4, "error device not bind user") // 设备没有绑定用户
)

// 业务错误
var (
	LErrDeviceIdOrToken = NewLError(1001, "error device token")       // 设备id或者token错误
	LErrNumberUsed      = NewLError(1002, "error number has be used") // 手机号码已经被使用
	LErrSecretKey       = NewLError(1003, "error secret key")         // 用户秘钥错误
	LErrDeviceNotFound  = NewLError(1004, "error device not found")   // 设备找不到
	LErrToken           = NewLError(1005, "error token")              // token错误
	LErrUserNotFound    = NewLError(1006, "error user not found")     // 用户找不到
)
