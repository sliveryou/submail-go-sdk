package client

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

// OptionFunc represents the submail common client option func.
type OptionFunc func(c *Client)

// WithHTTPClient makes the submail common client using the http client.
func WithHTTPClient(client *http.Client) OptionFunc {
	return func(c *Client) {
		c.client = client
	}
}

// New returns the submail common client.
func New(appId, appKey, signType string, opts ...OptionFunc) *Client {
	c := &Client{}
	c.appId = appId
	c.appKey = appKey
	c.signType = signType
	c.client = &http.Client{}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// GetTimestamp requests the submail timestamp service and returns the unix timestamp.
func (c *Client) GetTimestamp() (string, error) {
	resp, err := c.Get(timestampURL)
	if err != nil {
		return "", err
	}

	var t timestampResp
	err = json.Unmarshal(resp, &t)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(t.Timestamp, 10), nil
}

// Sign signs the request param and returns the url.Values constructed by it.
func (c *Client) Sign(param Param) (url.Values, error) {
	params, err := param.Params()
	if err != nil {
		return nil, err
	}
	params.Add("appid", c.appId)

	if c.signType != SignTypeNormal {
		timestamp, err := c.GetTimestamp()
		if err != nil {
			return nil, err
		}

		params.Add("timestamp", timestamp)
		params.Add("sign_type", c.signType)
		params.Add("sign_version", "2")
	}

	signature := c.appKey

	if c.signType == SignTypeSha1 {
		s := sha1.New()
		s.Write([]byte(c.genSignStr(params)))
		signature = hex.EncodeToString(s.Sum(nil))
	} else if c.signType == SignTypeMd5 {
		m := md5.New()
		m.Write([]byte(c.genSignStr(params)))
		signature = hex.EncodeToString(m.Sum(nil))
	}

	params.Add("signature", signature)

	return params, nil
}

// genSignStr returns the string to be signed
func (c *Client) genSignStr(params url.Values) string {
	values := url.Values{}
	for key := range params {
		if _, ok := notSignParams[key]; !ok {
			values.Add(key, params.Get(key))
		}
	}

	e, _ := url.QueryUnescape(values.Encode())
	signStr := c.appId + c.appKey + e + c.appId + c.appKey

	return signStr
}

// Do requests the submail service by param.
func (c *Client) Do(param Param, enableMultipart ...bool) error {
	enable := false
	if len(enableMultipart) != 0 {
		enable = enableMultipart[0]
	}

	params, err := c.Sign(param)
	if err != nil {
		return err
	}

	var resp []byte
	if !enable {
		resp, err = c.Post(param.RequestURL(), params)
	} else {
		resp, err = c.PostMultipart(param.RequestURL(), params)
	}
	if err != nil {
		return err
	}

	var result commonResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return err
	}
	// fmt.Printf("%+v\n", result)

	if result.Status != StatusSuccess {
		return errors.New(result.Msg)
	}
	return nil
}
