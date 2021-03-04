package mail

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
		To:           []*ToParam{{Name: "alice", Address: "alice@qq.com"}},
		From:         "sender@qq.com",
		FromName:     "sender",
		Subject:      "subject",
		Text:         "text",
		Vars:         map[string]string{"code": "123456"},
		Asynchronous: false,
	}

	err := c.Send(sp)
	t.Log(err)
}

func TestClient_XSend(t *testing.T) {
	c = New("appId", "appKey", "sha1")
	t.Log(c)

	xsp := &XSendParam{
		To:           []*ToParam{{Name: "alice", Address: "alice@qq.com"}},
		From:         "sender@qq.com",
		FromName:     "sender",
		Subject:      "subject",
		Project:      "project",
		Vars:         map[string]string{"code": "123456"},
		Asynchronous: false,
	}

	err := c.XSend(xsp)
	t.Log(err)
}
