package model

import (
	"time"
)

// User 账户
type User struct {
	Id         int64     `json:"id"`       // 用户id
	AppId      int64     `json:"app_id"`   // app_id
	UserId     int64     `json:"number"`   // 手机号
	Nickname   string    `json:"nickname"` // 昵称
	Sex        int       `json:"sex"`      // 性别，1:男；2:女
	AvatarUrl  string    `json:"avatar"`   // 用户头像
	Extra      string    `json:"extra"`    // 附加属性
	CreateTime time.Time `json:"-"`        // 创建时间
	UpdateTime time.Time `json:"-"`        // 更新时间
}

// UserRegist 用户注册
type AddUser struct {
	AppId     int64  `json:"app_id"`   // app_id
	UserId    int64  `json:"user_id"`  // 用户id
	Nickname  string `json:"nickname"` // 昵称
	Sex       int    `json:"sex"`      // 性别，1:男；2:女
	AvatarUrl string `json:"avatar"`   // 用户头像
	Extra     string `json:"extra"`    // 附加属性
}

// SignIn 登录结构体
type SignIn struct {
	Number   string `json:"number"`
	Password string `json:"password"`
}

// SignInResp 登录响应
type SignInResp struct {
	SendSequence int64 `json:"send_sequence"` // 发送序列号
	SyncSequence int64 `json:"sync_sequence"` // 同步序列号
}
