# Submail SDK for Go

*[English](README.md) ∙ [简体中文](README_zh-CN.md)*

[![Github License](https://img.shields.io/github/license/sliveryou/submail-go-sdk.svg?style=flat)](https://github.com/sliveryou/submail-go-sdk/blob/master/LICENSE)
[![Go Doc](https://godoc.org/github.com/sliveryou/submail-go-sdk?status.svg)](https://pkg.go.dev/github.com/sliveryou/submail-go-sdk)
[![Go Report](https://goreportcard.com/badge/github.com/sliveryou/submail-go-sdk)](https://goreportcard.com/report/github.com/sliveryou/submail-go-sdk)
[![Github Latest Release](https://img.shields.io/github/release/sliveryou/submail-go-sdk.svg?style=flat)](https://github.com/sliveryou/submail-go-sdk/releases/latest)
[![Github Latest Tag](https://img.shields.io/github/tag/sliveryou/submail-go-sdk.svg?style=flat)](https://github.com/sliveryou/submail-go-sdk/tags)
[![Github Stars](https://img.shields.io/github/stars/sliveryou/submail-go-sdk.svg?style=flat)](https://github.com/sliveryou/submail-go-sdk/stargazers)

The Go SDK is based on the official APIs of [Submail](https://www.mysubmail.com/documents).

## Installation

Download package by using:

```shell script
$ go get github.com/sliveryou/submail-go-sdk
```

## Usage Example

```golang
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
```

More example projects can be found at *github.com/sliveryou/submail-go-sdk/examples*.
