package mms

import (
	"encoding/json"
	"net/url"

	"github.com/sliveryou/submail-go-sdk/client"
)

const (
	// Submail mms xsend service url.
	xsendURL = client.APIDomain + "/mms/xsend"
	// Submail mms multixsend service url.
	multixsendURL = client.APIDomain + "/mms/multixsend"
)

// Client represents the submail mms client.
type Client struct {
	*client.Client
}

// XSendParam represents the mms xsend param and implements the client.Param interface.
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

// MultiParam represents the mms multi param.
type MultiParam struct {
	To   string
	Vars map[string]string
}

// MultiXSendParam represents the mms multixsend param and implements the client.Param interface.
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
