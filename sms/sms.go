package sms

import (
	"github.com/sliveryou/submail-go-sdk/client"
)

// https://www.mysubmail.com/documents/LJ4xa2
func New(appId, appKey, signType string) *Client {
	return &Client{Client: client.New(appId, appKey, signType)}
}

// https://www.mysubmail.com/documents/FppOR3
func (c *Client) Send(p *SendParam) error {
	return c.Do(p)
}

// https://www.mysubmail.com/documents/OOVyh
func (c *Client) XSend(p *XSendParam) error {
	return c.Do(p)
}

// https://www.mysubmail.com/documents/KZjET3
func (c *Client) MultiSend(p *MultiSendParam) error {
	return c.Do(p)
}

// https://www.mysubmail.com/documents/eM4rY2
func (c *Client) MultiXSend(p *MultiXSendParam) error {
	return c.Do(p)
}
