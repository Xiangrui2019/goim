package model

import (
	"time"
)

/**
CREATE TABLE `friend` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `app_id` bigint(20) NOT NULL COMMENT 'app_id',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '账户id',
  `friend_id` bigint(20) unsigned NOT NULL COMMENT '好友账户id',
  `label` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '备注',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_app_id_user_id_friend_id` (`app_id`,`user_id`,`friend_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='好友关系';
*/

// Friend 好友关系
type Friend struct {
	Id         int64     `json:"id"`          // 自增主键
	AppId      int64     `json:"app_id"`      // app_id
	UserId     int64     `json:"user_id"`     // 账户id
	FriendId   int64     `json:"friend_id"`   // 好友账户id
	Label      string    `json:"label"`       // 备注，标签
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
}

// UserFriend 用户好友信息
type UserFriend struct {
	UserId   int64  `json:"user_id"` // 用户id
	Label    string `json:"lable"`   // 用户对好友的标签
	Number   string `json:"number"`  // 手机号
	Nickname string `json:"name"`    // 昵称
	Sex      int    `json:"sex"`     // 性别，1:男；2:女
	Avatar   string `json:"avatar"`  // 用户头像
}
