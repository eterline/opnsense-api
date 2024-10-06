package opnsenseapi

import (
	"encoding/json"
	"log"

	"github.com/eterline/opnsense-api/utillis"
)

type DhcpV6Status struct {
	Status string `json:"status"`
	Widget struct {
		CaptionStop    string `json:"caption_stop"`
		CaptionStart   string `json:"caption_start"`
		CaptionRestart string `json:"caption_restart"`
	} `json:"widget"`
}

type DhcpV6 OpnsenseClient

func InitDhcpV6(cl OpnsenseClient) DhcpV6 {
	return DhcpV6(cl)
}

func (oc *DhcpV6) Status() DhcpV6Status {
	var res DhcpV6Status
	body, err := utillis.GetRequest(oc.HostURL, oc.BasicToken, "api/dhcpv6/service/status")
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func (oc *DhcpV6) Restart() error {
	return utillis.PostRequest(oc.HostURL, oc.BasicToken, "api/dhcpv6/service/restart")
}

func (oc *DhcpV6) Stop() error {
	return utillis.PostRequest(oc.HostURL, oc.BasicToken, "api/dhcpv6/service/stop")
}

func (oc *DhcpV6) Start() error {
	return utillis.PostRequest(oc.HostURL, oc.BasicToken, "api/dhcpv6/service/start")
}
