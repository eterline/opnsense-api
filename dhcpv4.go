package opnsenseapi

import (
	"encoding/json"
	"log"

	"github.com/eterline/opnsense-api/utillis"
)

type DhcpV4Status struct {
	Status string `json:"status"`
	Widget struct {
		CaptionStop    string `json:"caption_stop"`
		CaptionStart   string `json:"caption_start"`
		CaptionRestart string `json:"caption_restart"`
	} `json:"widget"`
}

type DhcpV4 OpnsenseClient

func InitDhcpV4(cl OpnsenseClient) DhcpV4 {
	return DhcpV4(cl)
}

func (oc *DhcpV4) Status() DhcpV4Status {
	var res DhcpV4Status
	body, err := utillis.GetRequest(oc.HostURL, oc.BasicToken, "api/dhcpv4/service/status")
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func (oc *DhcpV4) Restart() error {
	return utillis.PostRequest(oc.HostURL, oc.BasicToken, "api/dhcpv4/service/restart")
}

func (oc *DhcpV4) Stop() error {
	return utillis.PostRequest(oc.HostURL, oc.BasicToken, "api/dhcpv4/service/stop")
}

func (oc *DhcpV4) Start() error {
	return utillis.PostRequest(oc.HostURL, oc.BasicToken, "api/dhcpv4/service/start")
}
