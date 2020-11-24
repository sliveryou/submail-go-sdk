package voice

import (
	"github.com/sliveryou/submail-go-sdk/client"
)

// New returns the submail voice client,
// more look at: https://www.mysubmail.com/documents/PopwU2
func New(appId, appKey, signType string) *Client {
	return &Client{Client: client.New(appId, appKey, signType)}
}

// Send requests the submail voice send service by param,
// more look at: https://www.mysubmail.com/documents/meE3C1
func (c *Client) Send(p *SendParam) error {
	return c.Do(p)
}

// XSend requests the submail voice xsend service by param,
// more look at: https://www.mysubmail.com/documents/KbG03
func (c *Client) XSend(p *XSendParam) error {
	return c.Do(p)
}

// MultiXSend requests the submail voice multixsend service by param,
// more look at: https://www.mysubmail.com/documents/FkgkM2
func (c *Client) MultiXSend(p *MultiXSendParam) error {
	return c.Do(p)
}
