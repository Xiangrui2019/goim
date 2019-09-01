package model

import (
	"time"
)

/**
CREATE TABLE `message` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `app_id` int(11) NOT NULL COMMENT 'app_id',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户id',
  `message_id` bigint(20) unsigned NOT NULL COMMENT '消息id',
  `sender_type` tinyint(3) NOT NULL COMMENT '发送者类型',
  `sender_id` bigint(20) unsigned NOT NULL COMMENT '发送者id',
  `sender_device_id` bigint(20) unsigned NOT NULL COMMENT '发送设备id',
  `receiver_type` tinyint(3) NOT NULL COMMENT '接收者类型,1:个人；2：群组',
  `receiver_id` bigint(20) unsigned NOT NULL COMMENT '接收者id,如果是单聊信息，则为user_id，如果是群组消息，则为group_id',
  `to_user_ids` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '需要@的用户id列表，多个用户用，隔开',
  `message_content_id` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '内容',
  `user_seq` bigint(20) unsigned NOT NULL COMMENT '消息序列号',
  `send_time` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '消息发送时间',
  `status` tinyint(255) NOT NULL COMMENT '消息状态，0：未处理1：消息撤回',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_id_sequence` (`user_id`,`user_seq`),
  KEY `idx_message_id` (`message_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='消息';
*/

// Message 消息
type Message struct {
	Id             int64     `json:"id"`               // 自增主键
	MessageId      int64     `json:"message_id"`       // 消息id
	UserId         int64     `json:"user_id"`          // 用户id
	SenderType     int       `json:"sender_type"`      // 发送者类型
	SenderId       int64     `json:"sender"`           // 发送者账户id
	SenderDeviceId int64     `json:"sender_device_id"` // 发送者设备id
	ReceiverType   int       `json:"receiver_type"`    // 接收者账户id
	ReceiverId     int64     `json:"receiver"`         // 接收者id,如果是单聊信息，则为user_id，如果是群组消息，则为group_id
	Type           int       `json:"type"`             // 消息类型,0：文本；1：语音；2：图片
	Content        string    `json:"content"`          // 内容
	Sequence       int64     `json:"sequence"`         // 消息同步序列
	SendTime       time.Time `json:"send_time"`        // 消息发送时间
	CreateTime     time.Time `json:"create_time"`      // 创建时间
}
