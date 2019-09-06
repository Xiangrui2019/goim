package transfer

type OfflineReq struct {
	DeviceId int64 `json:"device_id"`
	UserId   int64 `json:"user_id"`
}

type OfflineResp struct {
}
