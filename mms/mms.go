package mms

import (
	"github.com/sliveryou/submail-go-sdk/client"
)

// New returns the submail mms client,
// more look at: https://www.mysubmail.com/documents/Ww6KQ3
func New(appId, appKey, signType string, opts ...client.OptionFunc) *Client {
	return &Client{Client: client.New(appId, appKey, signType, opts...)}
}

// XSend requests the submail mms xsend service by param,
// more look at: https://www.mysubmail.com/documents/N6ktR
func (c *Client) XSend(p *XSendParam) error {
	return c.Do(p)
}

// MultiXSend requests the submail mms multixsend service by param,
// more look at: https://www.mysubmail.com/documents/UPKTG2
func (c *Client) MultiXSend(p *MultiXSendParam) error {
	return c.Do(p)
}
