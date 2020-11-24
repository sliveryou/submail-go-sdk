package main

import (
	"github.com/sliveryou/submail-go-sdk/mms"
)

func main() {
	client := mms.New("AppId", "AppKey", "SignType")

	xsp := &mms.XSendParam{
		To:      "12345678910",
		Project: "project",
		Vars:    map[string]string{"code": "123456"},
		Tag:     "tag",
	}

	err := client.XSend(xsp)
	if err != nil {
		// HandleError(err)
	}

	mxsp := &mms.MultiXSendParam{
		Project: "project",
		Multi: []*mms.MultiParam{
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
