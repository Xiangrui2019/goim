package service

import (
	"goim/logic/dao"
	"goim/logic/model"
	"goim/public/imctx"
)

type messageBodyService struct{}

var MessageBodyService = new(messageBodyService)

// todo 此处应该支持分库分表
// Add 添加消息内容
func (*messageBodyService) Add(ctx *imctx.Context, messageBodyId int64, content string) error {
	return dao.MessageBodyDao.Add(ctx, "message_body", messageBodyId, content)
}

// Get 获取消息内容
func (*messageBodyService) Get(ctx *imctx.Context, messageBodyId int64) (*model.MessageBody, error) {
	return dao.MessageBodyDao.Get(ctx, "message_body", messageBodyId)
}
