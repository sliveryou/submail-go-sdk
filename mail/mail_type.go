package mail

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/sliveryou/submail-go-sdk/client"
)

const (
	// sendURL represents the submail mail send service url.
	sendURL = client.APIDomain + "/mail/send"
	// xsendURL represents the submail mail xsend service url.
	xsendURL = client.APIDomain + "/mail/xsend"
)

// Client represents the submail mail client.
type Client struct {
	*client.Client
}

// ToParam represents the mail to param.
type ToParam struct {
	Name    string
	Address string
}

// SendParam represents the mail send param and implements the client.Param interface.
type SendParam struct {
	To           []*ToParam
	From         string
	FromName     string
	Reply        string
	Cc           []string
	Bcc          []string
	Subject      string
	Html         string
	Text         string
	Vars         map[string]string
	Links        map[string]string
	Headers      map[string]string
	Asynchronous bool
	Attachments  []string
}

// Params implements the client.Param interface Params method.
func (p *SendParam) Params() (url.Values, error) {
	params := url.Values{}

	if len(p.To) > 0 {
		var receivers []string
		for _, to := range p.To {
			receivers = append(receivers, fmt.Sprintf("%s<%s>", to.Name, to.Address))
		}
		params.Add("to", strings.Join(receivers, ","))
	}

	if p.Subject != "" {
		params.Add("subject", p.Subject)
	}

	if p.From != "" {
		params.Add("from", p.From)
	}

	if p.FromName != "" {
		params.Add("from_name", p.FromName)
	}

	if p.Reply != "" {
		params.Add("reply", p.Reply)
	}

	if len(p.Cc) > 0 {
		params.Add("cc", strings.Join(p.Cc, ","))
	}

	if len(p.Bcc) > 0 {
		params.Add("bcc", strings.Join(p.Bcc, ","))
	}

	if p.Asynchronous {
		params.Add("asynchronous", "true")
	} else {
		params.Add("asynchronous", "false")
	}

	if p.Html != "" {
		params.Add("html", p.Html)
	}

	if p.Text != "" {
		params.Add("text", p.Text)
	}

	if len(p.Vars) > 0 {
		vars, err := json.Marshal(p.Vars)
		if err != nil {
			return nil, err
		}

		params.Add("vars", string(vars))
	}

	if len(p.Links) > 0 {
		links, err := json.Marshal(p.Links)
		if err != nil {
			return nil, err
		}

		params.Add("links", string(links))
	}

	if len(p.Headers) > 0 {
		headers, err := json.Marshal(p.Headers)
		if err != nil {
			return nil, err
		}

		params.Add("headers", string(headers))
	}

	if len(p.Attachments) > 0 {
		params.Add("attachments", strings.Join(p.Attachments, ","))
	}

	return params, nil
}

// RequestURL implements the client.Param interface RequestURL method.
func (p *SendParam) RequestURL() string {
	return sendURL
}

// XSendParam represents the mail xsend param and implements the client.Param interface.
type XSendParam struct {
	To           []*ToParam
	From         string
	FromName     string
	Reply        string
	Cc           []string
	Bcc          []string
	Subject      string
	Project      string
	Vars         map[string]string
	Links        map[string]string
	Headers      map[string]string
	Asynchronous bool
}

// Params implements the client.Param interface Params method.
func (p *XSendParam) Params() (url.Values, error) {
	params := url.Values{}
	params.Add("project", p.Project)

	if len(p.To) > 0 {
		var receivers []string
		for _, to := range p.To {
			receivers = append(receivers, fmt.Sprintf("%s<%s>", to.Name, to.Address))
		}
		params.Add("to", strings.Join(receivers, ","))
	}

	if p.Subject != "" {
		params.Add("subject", p.Subject)
	}

	if p.From != "" {
		params.Add("from", p.From)
	}

	if p.FromName != "" {
		params.Add("from_name", p.FromName)
	}

	if p.Reply != "" {
		params.Add("reply", p.Reply)
	}

	if len(p.Cc) > 0 {
		params.Add("cc", strings.Join(p.Cc, ","))
	}

	if len(p.Bcc) > 0 {
		params.Add("bcc", strings.Join(p.Bcc, ","))
	}

	if p.Asynchronous {
		params.Add("asynchronous", "true")
	} else {
		params.Add("asynchronous", "false")
	}

	if len(p.Vars) > 0 {
		vars, err := json.Marshal(p.Vars)
		if err != nil {
			return nil, err
		}

		params.Add("vars", string(vars))
	}

	if len(p.Links) > 0 {
		links, err := json.Marshal(p.Links)
		if err != nil {
			return nil, err
		}

		params.Add("links", string(links))
	}

	if len(p.Headers) > 0 {
		headers, err := json.Marshal(p.Headers)
		if err != nil {
			return nil, err
		}

		params.Add("headers", string(headers))
	}

	return params, nil
}

// RequestURL implements the client.Param interface RequestURL method.
func (p *XSendParam) RequestURL() string {
	return xsendURL
}
