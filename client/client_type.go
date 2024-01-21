package client

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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
	APIDomain = "https://api-v4.mysubmail.com"

	// timestampURL represents the submail timestamp service url.
	timestampURL = APIDomain + "/service/timestamp"
)

// notSignParams represents the request not sign params.
var notSignParams = map[string]struct{}{
	"content":     {},
	"vars":        {},
	"multi":       {},
	"html":        {},
	"text":        {},
	"links":       {},
	"headers":     {},
	"attachments": {},
}

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
	client   *http.Client
}

// Get issues a get request to the specified URL and returns the response.
func (c *Client) Get(url string) ([]byte, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

// Post issues a post form request to the specified URL and returns the response.
func (c *Client) Post(url string, params url.Values) ([]byte, error) {
	resp, err := c.client.PostForm(url, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

// PostMultipart issues a post multipart request to the specified URL and returns the response.
func (c *Client) PostMultipart(url string, params url.Values) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key := range params {
		if key == "attachments" {
			attachments := strings.Split(params.Get(key), ",")
			if len(attachments) > 0 {
				for _, filename := range attachments {
					file, err := os.Open(filename)
					if err != nil {
						return nil, err
					}

					part, err := writer.CreateFormFile("attachments[]", filepath.Base(filename))
					if err != nil {
						return nil, err
					}

					_, err = io.Copy(part, file)
					if err != nil {
						return nil, err
					}

					_ = file.Close()
				}
			}
		} else {
			err := writer.WriteField(key, params.Get(key))
			if err != nil {
				return nil, err
			}
		}
	}

	contentType := writer.FormDataContentType()
	_ = writer.Close()

	resp, err := c.client.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

// timestampResp represents the timestamp service response.
type timestampResp struct {
	Timestamp int64 `json:"timestamp"`
}

// commonResp represents the common service response.
type commonResp struct {
	Status                  string      `json:"status"`
	Code                    int         `json:"code"`
	Msg                     string      `json:"msg"`
	SendId                  string      `json:"send_id"`
	Fee                     json.Number `json:"fee"`
	SmsCredits              string      `json:"sms_credits"`
	TransactionalSmsCredits string      `json:"transactional_sms_credits"`
}
