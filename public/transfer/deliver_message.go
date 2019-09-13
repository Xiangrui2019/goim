package transfer

const (
	MessageTypeSync = 1 // 消息同步
	MessageTypeMail = 2 // 消息投递
)

type DeliverMessageReq struct {
	DeviceId int64  // 设备id
	Bytes    []byte // 消息投递字节包
}

type DeliverMessageResp struct {
}
