package model

import (
	"time"
)

// Friend 好友关系
type Friend struct {
	Id         int64     `json:"id"`          // 自增主键
	AppId      int64     `json:"app_id"`      // app_id
	UserId     int64     `json:"user_id"`     // 账户id
	FriendId   int64     `json:"friend_id"`   // 好友账户id
	Label      string    `json:"label"`       // 备注，标签
	Extra      string    `json:"extra"`       // 附加属性
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
}

// UserFriend 用户好友信息
type UserFriend struct {
	UserId      int64     `json:"user_id"`      // 用户id
	Label       string    `json:"lable"`        // 用户对好友的标签
	FriendExtra string    `json:"friend_extra"` // 朋友附加属性
	Nickname    string    `json:"name"`         // 昵称
	Sex         int       `json:"sex"`          // 性别，1:男；2:女
	AvatarUrl   string    `json:"avatar"`       // 用户头像
	UserExtra   string    `json:"user_extra"`   // 用户附加属性
	CreateTime  time.Time `json:"create_time"`  // 创建时间
	UpdateTime  time.Time `json:"update_time"`  // 更新时间
}
