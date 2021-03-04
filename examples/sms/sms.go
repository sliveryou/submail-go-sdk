package main

import (
	"github.com/sliveryou/submail-go-sdk/sms"
)

func main() {
	client := sms.New("AppId", "AppKey", "SignType")

	sp := &sms.SendParam{
		To:      "12345678910",
		Content: "content",
	}

	err := client.Send(sp)
	if err != nil {
		// HandleError(err)
	}

	xsp := &sms.XSendParam{
		To:      "12345678910",
		Project: "project",
		Vars:    map[string]string{"code": "123456"},
	}

	err = client.XSend(xsp)
	if err != nil {
		// HandleError(err)
	}

	msp := &sms.MultiSendParam{
		Content: "content",
		Multi: []*sms.MultiParam{
			{To: "12345678910", Vars: map[string]string{"code": "123456"}},
			{To: "13579246810", Vars: map[string]string{"code": "456123"}},
		},
	}

	err = client.MultiSend(msp)
	if err != nil {
		// HandleError(err)
	}

	mxsp := &sms.MultiXSendParam{
		Project: "project",
		Multi: []*sms.MultiParam{
			{To: "12345678910", Vars: map[string]string{"code": "123456"}},
			{To: "13579246810", Vars: map[string]string{"code": "456123"}},
		},
		Tag: "tag",
	}

	err = client.MultiXSend(mxsp)
	if err != nil {
		// HandleError(err)
	}
}
