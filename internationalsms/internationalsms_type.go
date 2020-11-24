package internationalsms

import (
	"encoding/json"
	"net/url"

	"github.com/sliveryou/submail-go-sdk/client"
)

const (
	// Submail internationalsms send service url.
	sendURL = client.APIDomain + "/internationalsms/send"
	// Submail internationalsms xsend service url.
	xsendURL = client.APIDomain + "/internationalsms/xsend"
	// Submail internationalsms multixsend service url.
	multixsendURL = client.APIDomain + "/internationalsms/multixsend"
)

// Client represents the submail internationalsms client.
type Client struct {
	*client.Client
}

// SendParam represents the internationalsms send param and implements the client.Param interface.
type SendParam struct {
	To      string
	Content string
}

// Params implements the client.Param interface Params method.
func (p *SendParam) Params() (url.Values, error) {
	params := url.Values{}
	params.Add("to", p.To)
	params.Add("content", p.Content)

	return params, nil
}

// RequestURL implements the client.Param interface RequestURL method.
func (p *SendParam) RequestURL() string {
	return sendURL
}

// XSendParam represents the internationalsms xsend param and implements the client.Param interface.
type XSendParam struct {
	To      string
	Project string
	Vars    map[string]string
}

// Params implements the client.Param interface Params method.
func (p *XSendParam) Params() (url.Values, error) {
	params := url.Values{}
	params.Add("to", p.To)
	params.Add("project", p.Project)

	if len(p.Vars) > 0 {
		vars, err := json.Marshal(p.Vars)
		if err != nil {
			return nil, err
		}

		params.Add("vars", string(vars))
	}

	return params, nil
}

// RequestURL implements the client.Param interface RequestURL method.
func (p *XSendParam) RequestURL() string {
	return xsendURL
}

// MultiParam represents the internationalsms multi param.
type MultiParam struct {
	To   string
	Vars map[string]string
}

// MultiXSendParam represents the internationalsms multisend param and implements the client.Param interface.
type MultiXSendParam struct {
	Project string
	Multi   []*MultiParam
}

// Params implements the client.Param interface Params method.
func (p *MultiXSendParam) Params() (url.Values, error) {
	params := url.Values{}
	params.Add("project", p.Project)

	if len(p.Multi) > 0 {
		multi, err := json.Marshal(p.Multi)
		if err != nil {
			return nil, err
		}

		params.Add("multi", string(multi))
	}

	return params, nil
}

// RequestURL implements the client.Param interface RequestURL method.
func (p *MultiXSendParam) RequestURL() string {
	return multixsendURL
}
