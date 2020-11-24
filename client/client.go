package client

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/url"
	"sort"
	"strconv"

	"github.com/sliveryou/submail-go-sdk/util"
)

// New returns the submail common client.
func New(appId, appKey, signType string) *Client {
	return &Client{
		appId:    appId,
		appKey:   appKey,
		signType: signType,
	}
}

// GetTimestamp requests the submail timestamp service and returns the unix timestamp.
func (c *Client) GetTimestamp() (string, error) {
	resp, err := util.Get(timestampURL)
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

		params.Add("timestampResp", timestamp)
		params.Add("sign_type", c.signType)
		params.Add("sign_version", "2")
	}

	var keys []string
	for key := range params {
		if _, ok := notSignParams[key]; !ok {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)

	values := url.Values{}
	for _, key := range keys {
		values.Add(key, params.Get(key))
	}

	signature := c.appKey
	signStr := c.appId + c.appKey + values.Encode() + c.appId + c.appKey

	if c.signType == SignTypeSha1 {
		s := sha1.New()
		s.Write([]byte(signStr))
		signature = hex.EncodeToString(s.Sum(nil))
	} else if c.signType == SignTypeMd5 {
		m := md5.New()
		m.Write([]byte(signStr))
		signature = hex.EncodeToString(m.Sum(nil))
	}

	params.Add("signature", signature)

	return params, nil
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
		resp, err = util.Post(param.RequestURL(), params)
	} else {
		resp, err = util.PostMultipart(param.RequestURL(), params)
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
