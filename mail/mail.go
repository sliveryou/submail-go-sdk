package mail

import (
	"github.com/sliveryou/submail-go-sdk/client"
)

// New returns the submail mail client,
// more look at: https://www.mysubmail.com/documents/jkh2j
func New(appId, appKey, signType string, opts ...client.OptionFunc) *Client {
	return &Client{Client: client.New(appId, appKey, signType, opts...)}
}

// Send requests the submail mail send service by param,
// more look at: https://www.mysubmail.com/documents/4MfRT2
func (c *Client) Send(p *SendParam) error {
	return c.Do(p, true)
}

// XSend requests the submail mail xsend service by param,
// more look at: https://www.mysubmail.com/documents/Vu8Qh3
func (c *Client) XSend(p *XSendParam) error {
	return c.Do(p, true)
}
