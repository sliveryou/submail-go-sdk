package main

import (
	"github.com/sliveryou/submail-go-sdk/mail"
)

func main() {
	client := mail.New("AppId", "AppKey", "SignType")

	sp := &mail.SendParam{
		To:           []*mail.ToParam{{Name: "alice", Address: "alice@qq.com"}},
		From:         "sender@qq.com",
		FromName:     "sender",
		Subject:      "subject",
		Text:         "text",
		Vars:         map[string]string{"code": "123456"},
		Asynchronous: false,
		Tag:          "tag",
	}

	err := client.Send(sp)
	if err != nil {
		// HandleError(err)
	}

	xsp := &mail.XSendParam{
		To:           []*mail.ToParam{{Name: "alice", Address: "alice@qq.com"}},
		From:         "sender@qq.com",
		FromName:     "sender",
		Subject:      "subject",
		Project:      "project",
		Vars:         map[string]string{"code": "123456"},
		Asynchronous: false,
		Tag:          "tag",
	}

	err = client.XSend(xsp)
	if err != nil {
		// HandleError(err)
	}
}
