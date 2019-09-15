package model

import (
	"goim/public/logger"
	"goim/public/pb"
	"goim/public/util"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type MessageBody struct {
	Id            int64     // 自增主键
	MessageBodyId int64     // 消息体id
	MessageType   int32     // 消息类型
	Content       string    // 消息内容
	CreateTime    time.Time // 创建时间
}

/**

// 消息类型
enum MessageType {
    DEFALUT = 0; // 占位，无意义
    TEXT = 1; // 文本
    FACE = 2; // 表情
    VOICE = 3; // 语音消息
    IMAGE = 4; // 图片
    FILE = 5; // 文件
    LOCATION = 6; // 地理位置
    COMMAND = 7; // 指令推送
    CUSTOM = 8; // 自定义
}
*/

func PBToMessageBody(pbBody *pb.MessageBody) *MessageBody {
	if pbBody.MessageType == pb.MessageType_DEFALUT {
		logger.Logger.Error("error message type")
		return nil
	}

	var content interface{}
	switch pbBody.MessageType {
	case pb.MessageType_TEXT:
		content = pbBody.MessageContent.GetText()
	case pb.MessageType_FACE:
		content = pbBody.MessageContent.GetFace()
	case pb.MessageType_VOICE:
		content = pbBody.MessageContent.GetVoice()
	case pb.MessageType_IMAGE:
		content = pbBody.MessageContent.GetImage()
	case pb.MessageType_FILE:
		content = pbBody.MessageContent.GetFile()
	case pb.MessageType_LOCATION:
		content = pbBody.MessageContent.GetLocation()
	case pb.MessageType_COMMAND:
		content = pbBody.MessageContent.GetCommand()
	case pb.MessageType_CUSTOM:
		content = pbBody.MessageContent.GetCustom()
	}

	bytes, err := jsoniter.Marshal(content)
	if err != nil {
		logger.Sugar.Error(err)
		return nil
	}
	return &MessageBody{
		MessageType: int32(pbBody.MessageType),
		Content:     util.Bytes2str(bytes),
	}

}

func MessageBodyToPB(body *MessageBody) *pb.MessageBody {
	content := pb.MessageContent{}
	switch pb.MessageType(body.MessageType) {
	case pb.MessageType_TEXT:
		var text pb.Text
		err := jsoniter.Unmarshal(util.Str2bytes(body.Content), &text)
		if err != nil {
			logger.Sugar.Error(err)
			return nil
		}
		content.Content = &pb.MessageContent_Text{Text: &text}
	case pb.MessageType_FACE:
		var face pb.Face
		err := jsoniter.Unmarshal(util.Str2bytes(body.Content), &face)
		if err != nil {
			logger.Sugar.Error(err)
			return nil
		}
		content.Content = &pb.MessageContent_Face{Face: &face}
	case pb.MessageType_VOICE:
		var voice pb.Voice
		err := jsoniter.Unmarshal(util.Str2bytes(body.Content), &voice)
		if err != nil {
			logger.Sugar.Error(err)
			return nil
		}
		content.Content = &pb.MessageContent_Voice{Voice: &voice}
	case pb.MessageType_IMAGE:
		var image pb.Image
		err := jsoniter.Unmarshal(util.Str2bytes(body.Content), &image)
		if err != nil {
			logger.Sugar.Error(err)
			return nil
		}
		content.Content = &pb.MessageContent_Image{Image: &image}
	case pb.MessageType_FILE:
		var file pb.File
		err := jsoniter.Unmarshal(util.Str2bytes(body.Content), &file)
		if err != nil {
			logger.Sugar.Error(err)
			return nil
		}
		content.Content = &pb.MessageContent_File{File: &file}
	case pb.MessageType_LOCATION:
		var location pb.Location
		err := jsoniter.Unmarshal(util.Str2bytes(body.Content), &location)
		if err != nil {
			logger.Sugar.Error(err)
			return nil
		}
		content.Content = &pb.MessageContent_Location{Location: &location}
	case pb.MessageType_COMMAND:
		var command pb.Command
		err := jsoniter.Unmarshal(util.Str2bytes(body.Content), &command)
		if err != nil {
			logger.Sugar.Error(err)
			return nil
		}
		content.Content = &pb.MessageContent_Command{Command: &command}
	case pb.MessageType_CUSTOM:
		var custom pb.Custom
		err := jsoniter.Unmarshal(util.Str2bytes(body.Content), &custom)
		if err != nil {
			logger.Sugar.Error(err)
			return nil
		}
		content.Content = &pb.MessageContent_Custom{Custom: &custom}
	}

	return &pb.MessageBody{
		MessageType:    pb.MessageType(body.MessageType),
		MessageContent: &content,
	}

}
