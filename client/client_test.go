package client

import (
	"testing"
)

var c *Client

func TestNew(t *testing.T) {
	c = New("appId", "appKey", "sha1")
	t.Log(c)
}

func TestClient_GetTimestamp(t *testing.T) {
	c = New("appId", "appKey", "sha1")
	t.Log(c)

	timestamp, err := c.GetTimestamp()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(timestamp)
	}
}
