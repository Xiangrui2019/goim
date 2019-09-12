package transfer

type SendMessageReq struct {
	DeviceId int64
	UserId   int64
	Bytes    []byte
}

type SendMessageResp struct {
	Code    int32
	Message string
	Bytes   []byte
}

func NewSendMessageResp() {

}
