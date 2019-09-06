package dao

import (
	"encoding/json"
	"fmt"
	"goim/logic/model"
	"testing"
	"time"
)

func TestMessageDao_Add(t *testing.T) {
	message := model.Message{
		MessageId:      "1",
		AppId:          1,
		UserId:         2,
		SenderType:     2,
		SenderId:       2,
		SenderDeviceId: 2,
		ReceiverType:   2,
		ReceiverId:     2,
		ToUserIds:      "2",
		MessageBodyId:  2,
		UserSeq:        2,
		SendTime:       time.Now(),
	}
	fmt.Println(MessageDao.Add(ctx, "message", message))
}

func TestMessageDao_ListByUserIdAndUserSeq(t *testing.T) {
	messages, err := MessageDao.ListByUserIdAndUserSeq(ctx, "message", 1, 0, 0)
	fmt.Println(err)
	bytes, _ := json.Marshal(messages)
	fmt.Println(string(bytes))
}
