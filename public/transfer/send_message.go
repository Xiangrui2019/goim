package transfer

type SendMessageReq struct {
	DeviceId int64
	UserId   int64
	Bytes    []byte
}

type SendMessageResp struct {
	Bytes []byte
}
