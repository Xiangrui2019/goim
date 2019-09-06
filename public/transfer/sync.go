package transfer

// 同步消息触发
type SyncReq struct {
	DeviceId int64  `json:"device_id"`     // 设备id
	UserId   int64  `json:"user_id"`       // 用户id
	Bytes    []byte `json:"sync_sequence"` // 已经同步的消息序列号
}

// 同步消息触发
type SyncResp struct {
	Bytes []byte
}
