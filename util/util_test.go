package util

import (
	"net/url"
	"testing"
)

func TestGet(t *testing.T) {
	resp, err := Get("https://www.httpbin.org/get")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(string(resp))
	}
}

func TestPost(t *testing.T) {
	params := url.Values{}
	params.Add("user", "sliver")
	params.Add("password", "123456")

	resp, err := Post("https://www.httpbin.org/post", params)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(string(resp))
	}
}

func TestPostMultipart(t *testing.T) {
	params := url.Values{}
	params.Add("user", "sliver")
	params.Add("password", "123456")

	resp, err := PostMultipart("https://www.httpbin.org/post", params)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(string(resp))
	}
}
