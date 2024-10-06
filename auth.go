package opnsenseapi

import (
	"net/url"

	"github.com/eterline/opnsense-api/utillis"
)

type OpnsenseClient struct {
	BasicToken string
	HostURL    *url.URL
	TLS        bool
}

func NewClient(user, pass, uri string) (OpnsenseClient, error) {
	if err := utillis.CorrectCerdentials(user, pass); err != nil {
		return OpnsenseClient{}, err
	}
	u, err := url.ParseRequestURI(uri)
	if err != nil {
		return OpnsenseClient{}, err
	}
	return OpnsenseClient{
		BasicToken: utillis.BasicAuthString(user, pass),
		HostURL:    u,
		TLS:        true,
	}, err
}

func (oc *OpnsenseClient) IgnoreSSL() {
	oc.TLS = false
}
