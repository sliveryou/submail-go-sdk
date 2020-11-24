package client

import (
	"net/url"
)

const (
	// SignTypeNormal represents the request sign type normal.
	SignTypeNormal = "normal"
	// SignTypeSha1 represents the request sign type sha1.
	SignTypeSha1 = "sha1"
	// SignTypeMd5 represents the request sign type md5.
	SignTypeMd5 = "md5"

	// StatusSuccess represents the response status success.
	StatusSuccess = "success"
	// StatusError represents the response status error.
	StatusError = "error"

	// APIDomain represents the submail api domain.
	APIDomain = "https://api.mysubmail.com"

	// timestampURL represents the submail timestamp service url.
	timestampURL = APIDomain + "/service/timestamp"
)

var (
	// notSignParams represents the request not sign params.
	notSignParams = map[string]struct{}{
		"content":     {},
		"vars":        {},
		"multi":       {},
		"html":        {},
		"text":        {},
		"links":       {},
		"headers":     {},
		"attachments": {},
	}
)

// Param is the interface that wraps the Params and RequestURL methods.
type Param interface {
	// Params returns the params in the form of url.Values.
	Params() (url.Values, error)
	// RequestURL returns the request url of the param.
	RequestURL() string
}

// Client represents the submail common client.
type Client struct {
	appId    string
	appKey   string
	signType string
}

// timestampResp represents the timestamp service response.
type timestampResp struct {
	Timestamp int64 `json:"timestamp"`
}

// commonResp represents the common service response.
type commonResp struct {
	Status                  string `json:"status"`
	Code                    int    `json:"code"`
	Msg                     string `json:"msg"`
	SendId                  string `json:"send_id"`
	Fee                     int    `json:"fee"`
	SmsCredits              string `json:"sms_credits"`
	TransactionalSmsCredits string `json:"transactional_sms_credits"`
}
