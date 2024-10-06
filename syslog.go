package opnsenseapi

import (
	"encoding/json"
	"log"

	"github.com/eterline/opnsense-api/utillis"
)

type SyslogStats struct {
	Total    int `json:"total"`
	RowCount int `json:"rowCount"`
	Current  int `json:"current"`
	Rows     []struct {
		Code           string `json:"#"`
		Description    string `json:"Description"`
		SourceName     string `json:"SourceName"`
		SourceID       string `json:"SourceId"`
		SourceInstance string `json:"SourceInstance"`
		State          string `json:"State"`
		Type           string `json:"Type"`
		Number         string `json:"Number"`
	} `json:"rows"`
}

type SyslogStatus struct {
	Status string `json:"status"`
	Widget struct {
		CaptionStop    string `json:"caption_stop"`
		CaptionStart   string `json:"caption_start"`
		CaptionRestart string `json:"caption_restart"`
	} `json:"widget"`
}

type Syslog OpnsenseClient

func InitSyslog(cl OpnsenseClient) Syslog {
	return Syslog(cl)
}

func (oc *Syslog) Status() SyslogStatus {
	var res SyslogStatus
	body, err := utillis.GetRequest(oc.HostURL, oc.BasicToken, "api/syslog/service/status")
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func (oc *Syslog) Restart() error {
	return utillis.PostRequest(oc.HostURL, oc.BasicToken, "api/syslog/service/restart")
}

func (oc *Syslog) Stop() error {
	return utillis.PostRequest(oc.HostURL, oc.BasicToken, "api/syslog/service/stop")
}

func (oc *Syslog) Start() error {
	return utillis.PostRequest(oc.HostURL, oc.BasicToken, "api/syslog/service/start")
}

func (oc *Syslog) Stats() SyslogStats {
	var res SyslogStats
	body, err := utillis.GetRequest(oc.HostURL, oc.BasicToken, "api/syslog/service/stats")
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Panic(err)
	}
	return res
}
