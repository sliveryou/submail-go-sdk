package sms

import (
	"testing"
)

var c *Client

func TestNew(t *testing.T) {
	c = New("appId", "appKey", "sha1")
	t.Log(c)
}

func TestClient_Send(t *testing.T) {
	c = New("appId", "appKey", "sha1")
	t.Log(c)

	sp := &SendParam{
		To:      "12345678910",
		Content: "content",
	}

	err := c.Send(sp)
	t.Log(err)
}

func TestClient_XSend(t *testing.T) {
	c = New("appId", "appKey", "sha1")
	t.Log(c)

	xsp := &XSendParam{
		To:      "12345678910",
		Project: "project",
		Vars:    map[string]string{"code": "123456"},
	}

	err := c.XSend(xsp)
	t.Log(err)
}

func TestClient_MultiSend(t *testing.T) {
	c = New("appId", "appKey", "sha1")
	t.Log(c)

	msp := &MultiSendParam{
		Content: "content",
		Multi: []*MultiParam{
			{To: "12345678910", Vars: map[string]string{"code": "123456"}},
			{To: "13579246810", Vars: map[string]string{"code": "456123"}},
		},
	}

	err := c.MultiSend(msp)
	t.Log(err)
}

func TestClient_MultiXSend(t *testing.T) {
	c = New("appId", "appKey", "sha1")
	t.Log(c)

	mxsp := &MultiXSendParam{
		Project: "project",
		Multi: []*MultiParam{
			{To: "12345678910", Vars: map[string]string{"code": "123456"}},
			{To: "13579246810", Vars: map[string]string{"code": "456123"}},
		},
	}

	err := c.MultiXSend(mxsp)
	t.Log(err)
}
