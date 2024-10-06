package opnsenseapi

import (
	"encoding/json"
	"log"

	"github.com/eterline/opnsense-api/utillis"
)

type WgService struct {
	Total    int    `json:"total,omitempty"`
	RowCount int    `json:"rowCount,omitempty"`
	Current  int    `json:"current,omitempty"`
	Rows     []Rows `json:"rows,omitempty"`
}
type Rows struct {
	If                  string `json:"if,omitempty"`
	Type                string `json:"type,omitempty"`
	PublicKey           string `json:"public-key,omitempty"`
	ListenPort          string `json:"listen-port,omitempty"`
	Fwmark              string `json:"fwmark,omitempty"`
	Endpoint            string `json:"endpoint,omitempty"`
	Status              string `json:"status,omitempty"`
	Name                string `json:"name,omitempty"`
	Ifname              string `json:"ifname,omitempty"`
	AllowedIps          string `json:"allowed-ips,omitempty"`
	LatestHandshake     int    `json:"latest-handshake,omitempty"`
	TransferRx          int    `json:"transfer-rx,omitempty"`
	TransferTx          int    `json:"transfer-tx,omitempty"`
	PersistentKeepalive string `json:"persistent-keepalive,omitempty"`
}

type OpnsenseWg OpnsenseClient

func InitWg(cl OpnsenseClient) OpnsenseWg {
	return OpnsenseWg(cl)
}

func (oc *OpnsenseWg) ServiceShow() WgService {
	var res WgService
	body, err := utillis.GetRequest(oc.HostURL, oc.BasicToken, "api/wireguard/service/show")
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Panic(err)
	}
	return res
}
