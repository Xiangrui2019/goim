package server

import (
	"goim/logic/db"
	"goim/logic/model"
	"goim/logic/service"
	"goim/public/imctx"
	"goim/public/imerror"
	"goim/public/logger"
	"goim/public/pb"
	"goim/public/transfer"

	"github.com/golang/protobuf/proto"
)

func Context() *imctx.Context {
	return imctx.NewContext(db.Factoty.GetSession())
}

type LogicRPCServer struct{}

// SignIn 设备登录
func (s *LogicRPCServer) SignIn(req transfer.SignInReq, resp *transfer.SignInResp) error {
	signInReq := pb.SignInReq{}
	err := proto.Unmarshal(req.Bytes, &signInReq)
	if err != nil {
		logger.Sugar.Error(err)
		transfer.ErrorToSignInResp(err)
		return nil
	}

	err = service.AuthService.SignIn(Context(), signInReq.AppId, signInReq.UserId, signInReq.DeviceId, signInReq.Token)
	if err != nil {
		logger.Sugar.Error(err)
	}
	resp = transfer.ErrorToSignInResp(err)
	return nil
}

// Sync 设备同步消息
func (s *LogicRPCServer) Sync(req transfer.SyncReq, resp *transfer.SyncResp) error {
	if !req.IsSignIn {
		resp = transfer.ErrorToSyncResp(imerror.ErrUnauthorized, nil)
		return nil
	}

	syncReq := pb.SyncReq{}
	err := proto.Unmarshal(req.Bytes, &syncReq)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	message, err := service.MessageService.ListByUserIdAndSeq(Context(), req.AppId, req.UserId, syncReq.DeviceAck)
	if err != nil {
		logger.Sugar.Error(err)
		resp = transfer.ErrorToSyncResp(imerror.ErrUnauthorized, nil)
		return nil
	}
	resp = transfer.ErrorToSyncResp(nil, model.MessagesToPB(message))
	return nil
}

// SendMessage 发送消息
func (s *LogicRPCServer) SendMessage(req transfer.SendMessageReq, resp *transfer.SendMessageResp) error {
	sendMessageReq := pb.SendMessageReq{}
	err := proto.Unmarshal(req.Bytes, &sendMessageReq)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	if !req.IsSignIn {
		resp = transfer.ErrorToSendMessageResp(imerror.ErrUnauthorized, sendMessageReq.MessageId)
		return nil
	}

	return nil
}

// MessageACK 设备收到消息ack
func (s *LogicRPCServer) MessageACK(req transfer.MessageAckReq, resp *transfer.MessageAckResp) error {

	messageACK := pb.MessageACK{}
	err := proto.Unmarshal(req.Bytes, &messageACK)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// Offline 设备离线
func (s *LogicRPCServer) Offline(req transfer.OfflineReq, resp *transfer.OfflineResp) error {

	return nil
}
