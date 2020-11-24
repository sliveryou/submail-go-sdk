package main

import (
	"github.com/sliveryou/submail-go-sdk/voice"
)

func main() {
	client := voice.New("AppId", "AppKey", "SignType")

	sp := &voice.SendParam{
		To:      "12345678910",
		Content: "content",
	}

	err := client.Send(sp)
	if err != nil {
		// HandleError(err)
	}

	xsp := &voice.XSendParam{
		To:      "12345678910",
		Project: "project",
		Vars:    map[string]string{"code": "123456"},
	}

	err = client.XSend(xsp)
	if err != nil {
		// HandleError(err)
	}

	mxsp := &voice.MultiXSendParam{
		Project: "project",
		Multi: []*voice.MultiParam{
			{To: "12345678910", Vars: map[string]string{"code": "123456"}},
			{To: "13579246810", Vars: map[string]string{"code": "456123"}},
		},
	}

	err = client.MultiXSend(mxsp)
	if err != nil {
		// HandleError(err)
	}
}
