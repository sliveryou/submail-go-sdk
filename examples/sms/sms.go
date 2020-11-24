package main

import (
	"github.com/sliveryou/submail-go-sdk/sms"
)

func main() {
	client := sms.New("AppId", "AppKey", "SignType")

	sp := &sms.SendParam{
		To:      "12345678910",
		Content: "content",
		Tag:     "tag",
	}

	err := client.Send(sp)
	if err != nil {
		// HandleError(err)
	}

	xsp := &sms.XSendParam{
		To:      "12345678910",
		Project: "project",
		Vars:    map[string]string{"code": "123456"},
		Tag:     "tag",
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
		Tag: "tag",
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
