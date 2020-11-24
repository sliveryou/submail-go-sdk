package main

import (
	"github.com/sliveryou/submail-go-sdk/internationalsms"
)

func main() {
	client := internationalsms.New("AppId", "AppKey", "SignType")

	sp := &internationalsms.SendParam{
		To:      "12345678910",
		Content: "content",
	}

	err := client.Send(sp)
	if err != nil {
		// HandleError(err)
	}

	xsp := &internationalsms.XSendParam{
		To:      "12345678910",
		Project: "project",
		Vars:    map[string]string{"code": "123456"},
	}

	err = client.XSend(xsp)
	if err != nil {
		// HandleError(err)
	}

	mxsp := &internationalsms.MultiXSendParam{
		Project: "project",
		Multi: []*internationalsms.MultiParam{
			{To: "12345678910", Vars: map[string]string{"code": "123456"}},
			{To: "13579246810", Vars: map[string]string{"code": "456123"}},
		},
	}

	err = client.MultiXSend(mxsp)
	if err != nil {
		// HandleError(err)
	}
}
