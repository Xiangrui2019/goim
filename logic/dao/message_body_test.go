package dao

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMessageBodyDao_Add(t *testing.T) {
	fmt.Println(MessageBodyDao.Add(ctx, "message_body", 1, "1"))
}

func TestMessageBodyDao_List(t *testing.T) {
	bodys, err := MessageBodyDao.List(ctx, "message_body", []int64{1, 2})
	fmt.Println(err)
	bytes, _ := json.Marshal(bodys)
	fmt.Println(string(bytes))
}
