package model

// Group 群组
type Group struct {
	Id           int64  `json:"id"`           // 群组id
	AppId        int64  `json:"app_id"`       // appId
	GroupId      int64  `json:"group_id"`     // 群组id
	Name         string `json:"name"`         // 组名
	Introduction string `json:"introduction"` // 群简介
	UserNum      int    `json:"user_num"`     // 群组人数
	Type         int    `json:"type"`         // 群组类型
	Extra        string `json:"extra"`        // 附加属性
	CreateTime   string `json:"create_time"`  // 创建时间
	UpdateTime   string `json:"update_time"`  // 更新时间
}

type GroupUserUpdate struct {
	GroupId int64   `json:"group_id"` // 群组名称
	UserIds []int64 `json:"user_ids"` // 群组成员
}
