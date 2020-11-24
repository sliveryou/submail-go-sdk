package internationalsms

import (
	"github.com/sliveryou/submail-go-sdk/client"
)

// New returns the submail internationalsms client,
// more look at: https://www.mysubmail.com/documents/gQn2R3
func New(appId, appKey, signType string) *Client {
	return &Client{Client: client.New(appId, appKey, signType)}
}

// Send requests the submail internationalsms send service by param,
// more look at: https://www.mysubmail.com/documents/3UQA3
func (c *Client) Send(p *SendParam) error {
	return c.Do(p)
}

// XSend requests the submail internationalsms xsend service by param,
// more look at: https://www.mysubmail.com/documents/87QTB2
func (c *Client) XSend(p *XSendParam) error {
	return c.Do(p)
}

// MultiXSend requests the submail internationalsms multixsend service by param,
// more look at: https://www.mysubmail.com/documents/B70hy
func (c *Client) MultiXSend(p *MultiXSendParam) error {
	return c.Do(p)
}
