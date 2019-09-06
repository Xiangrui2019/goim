package model

import "time"

type MessageBody struct {
	Id            int64
	MessageBodyId int64
	Content       string
	CreateTime    time.Time
}
