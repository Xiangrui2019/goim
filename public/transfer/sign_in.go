package transfer

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
