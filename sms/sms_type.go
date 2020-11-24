package sms

import (
	"encoding/json"
	"net/url"

	"github.com/sliveryou/submail-go-sdk/client"
)

const (
	// Submail sms send service url.
	sendURL = client.APIDomain + "/message/send"
	// Submail sms xsend service url.
	xsendURL = client.APIDomain + "/message/xsend"
	// Submail sms multisend service url.
	multisendURL = client.APIDomain + "/message/multisend"
	// Submail sms multixsend service url.
	multixsendURL = client.APIDomain + "/message/multixsend"
)

// Client represents the submail sms client.
type Client struct {
	*client.Client
}

// SendParam represents the sms send param and implements the client.Param interface.
type SendParam struct {
	To      string
	Content string
	Tag     string
}

// Params implements the client.Param interface Params method.
func (p *SendParam) Params() (url.Values, error) {
	params := url.Values{}
	params.Add("to", p.To)
	params.Add("content", p.Content)

	if p.Tag != "" {
		params.Add("tag", p.Tag)
	}

	return params, nil
}

// RequestURL implements the client.Param interface RequestURL method.
func (p *SendParam) RequestURL() string {
	return sendURL
}

// XSendParam represents the sms xsend param and implements the client.Param interface.
type XSendParam struct {
	To      string
	Project string
	Vars    map[string]string
	Tag     string
}

// Params implements the client.Param interface Params method.
func (p *XSendParam) Params() (url.Values, error) {
	params := url.Values{}
	params.Add("to", p.To)
	params.Add("project", p.Project)

	if p.Tag != "" {
		params.Add("tag", p.Tag)
	}

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

// MultiParam represents the sms multi param.
type MultiParam struct {
	To   string
	Vars map[string]string
}

// MultiSendParam represents the sms multisend param and implements the client.Param interface.
type MultiSendParam struct {
	Content string
	Multi   []*MultiParam
	Tag     string
}

// Params implements the client.Param interface Params method.
func (p *MultiSendParam) Params() (url.Values, error) {
	params := url.Values{}
	params.Add("content", p.Content)

	if p.Tag != "" {
		params.Add("tag", p.Tag)
	}

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
func (p *MultiSendParam) RequestURL() string {
	return multisendURL
}

// MultiXSendParam represents the sms multixsend param and implements the client.Param interface.
type MultiXSendParam struct {
	Project string
	Multi   []*MultiParam
	Tag     string
}

// Params implements the client.Param interface Params method.
func (p *MultiXSendParam) Params() (url.Values, error) {
	params := url.Values{}
	params.Add("project", p.Project)

	if p.Tag != "" {
		params.Add("tag", p.Tag)
	}

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
