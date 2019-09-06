package transfer

type MessageAckReq struct {
	DeviceId int64
	UserId   int64
	Bytes    []byte
}

type MessageAckResp struct {
}
