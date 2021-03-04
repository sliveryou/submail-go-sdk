package sms

import (
	"github.com/sliveryou/submail-go-sdk/client"
)

// New returns the submail sms client,
// more look at: https://www.mysubmail.com/documents/LJ4xa2
func New(appId, appKey, signType string, opts ...client.OptionFunc) *Client {
	return &Client{Client: client.New(appId, appKey, signType, opts...)}
}

// Send requests the submail sms send service by param,
// more look at: https://www.mysubmail.com/documents/FppOR3
func (c *Client) Send(p *SendParam) error {
	return c.Do(p)
}

// XSend requests the submail sms xsend service by param,
// more look at: https://www.mysubmail.com/documents/OOVyh
func (c *Client) XSend(p *XSendParam) error {
	return c.Do(p)
}

// MultiSend requests the submail sms multisend service by param,
// more look at: https://www.mysubmail.com/documents/KZjET3
func (c *Client) MultiSend(p *MultiSendParam) error {
	return c.Do(p)
}

// MultiXSend requests the submail sms multixsend service by param,
// more look at: https://www.mysubmail.com/documents/eM4rY2
func (c *Client) MultiXSend(p *MultiXSendParam) error {
	return c.Do(p)
}
