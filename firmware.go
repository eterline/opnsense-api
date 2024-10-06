package opnsenseapi

import (
	"encoding/json"
	"log"

	"github.com/eterline/opnsense-api/utillis"
)

type Firmware OpnsenseClient

func InitFirmware(cl OpnsenseClient) Firmware {
	return Firmware(cl)
}

type FirmwareData struct {
	Firmware FirmwareParam `json:"firmware"`
}
type FirmwareParam struct {
	Mirror       string `json:"mirror"`
	Flavour      string `json:"flavour"`
	Plugins      string `json:"plugins"`
	Type         string `json:"type"`
	Subscription string `json:"subscription"`
	Reboot       string `json:"reboot"`
}

type FirmwareOpts struct {
	Families                Families `json:"families"`
	FamiliesAllowCustom     int      `json:"families_allow_custom"`
	FamiliesHasSubscription []string `json:"families_has_subscription"`
	Flavours                Flavours `json:"flavours"`
	FlavoursAllowCustom     bool     `json:"flavours_allow_custom"`
	FlavoursHasSubscription []string `json:"flavours_has_subscription"`
	Mirrors                 Mirrors  `json:"mirrors"`
	MirrorsAllowCustom      bool     `json:"mirrors_allow_custom"`
	MirrorsHasSubscription  []string `json:"mirrors_has_subscription"`
}
type Families struct {
	Business      string `json:"business"`
	NAMING_FAILED string `json:""`
	Devel         string `json:"devel"`
}
type Flavours struct {
	NAMING_FAILED string `json:""`
}
type Mirrors struct {
	NAMING_FAILED                            string `json:""`
	HTTPSMirrorsDotsrcOrgOpnsense            string `json:"https://mirrors.dotsrc.org/opnsense"`
	HTTPSOpnsenseAivianOrg                   string `json:"https://opnsense.aivian.org"`
	HTTPSOpnsenseUpdateDecisoCom             string `json:"https://opnsense-update.deciso.com"`
	HTTPSMirrorDNSRootDeOpnsense             string `json:"https://mirror.dns-root.de/opnsense"`
	HTTPSOpnsenseC0UrierNet                  string `json:"https://opnsense.c0urier.net"`
	HTTPSMirrorCloudfenceComBrOpnsense       string `json:"https://mirror.cloudfence.com.br/opnsense"`
	HTTPSOpnsenseMirrorHihoCh                string `json:"https://opnsense-mirror.hiho.ch"`
	HTTPSMirrorOpnsenseServerbaseCh          string `json:"https://mirror-opnsense.serverbase.ch"`
	HTTPSMirrorAms1NlLeasewebNetOpnsense     string `json:"https://mirror.ams1.nl.leaseweb.net/opnsense"`
	HTTPSMirrorFra10DeLeasewebNetOpnsense    string `json:"https://mirror.fra10.de.leaseweb.net/opnsense"`
	HTTPSMirrorSfo12UsLeasewebNetOpnsense    string `json:"https://mirror.sfo12.us.leaseweb.net/opnsense"`
	HTTPSMirrorWdc1UsLeasewebNetOpnsense     string `json:"https://mirror.wdc1.us.leaseweb.net/opnsense"`
	HTTPQuantumMirrorHuMirrorsPubOpnsense    string `json:"http://quantum-mirror.hu/mirrors/pub/opnsense"`
	HTTPMirrorMarwanMaOpnsense               string `json:"http://mirror.marwan.ma/opnsense/"`
	HTTPSMirrorsNycbugOrgPubOpnsense         string `json:"https://mirrors.nycbug.org/pub/opnsense"`
	HTTPSPkgOpnsenseOrg                      string `json:"https://pkg.opnsense.org"`
	HTTPMirrorTerrahostNoOpnsense            string `json:"http://mirror.terrahost.no/opnsense"`
	HTTPSWwwMirrorserviceOrgSitesOpnsenseOrg string `json:"https://www.mirrorservice.org/sites/opnsense.org"`
	HTTPMirrorVenturasystemsTechOpnsense     string `json:"http://mirror.venturasystems.tech/opnsense"`
}

type FirmwareInfo struct {
	ProductID      string      `json:"product_id"`
	ProductVersion string      `json:"product_version"`
	Packages       []Package   `json:"package"`
	Plugins        []Plugin    `json:"plugin"`
	ChangeLogs     []ChangeLog `json:"changelog"`
	Product        Product     `json:"product"`
}

type Package struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Comment    string `json:"comment"`
	Flatsize   string `json:"flatsize"`
	Locked     string `json:"locked"`
	Automatic  string `json:"automatic"`
	License    string `json:"license"`
	Repository string `json:"repository"`
	Origin     string `json:"origin"`
	Provided   string `json:"provided"`
	Installed  string `json:"installed"`
	Path       string `json:"path"`
	Configured string `json:"configured"`
}

type Plugin struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Comment    string `json:"comment"`
	Flatsize   string `json:"flatsize"`
	Locked     string `json:"locked"`
	Automatic  string `json:"automatic"`
	License    string `json:"license"`
	Repository string `json:"repository"`
	Origin     string `json:"origin"`
	Provided   string `json:"provided"`
	Installed  string `json:"installed"`
	Path       string `json:"path"`
	Configured string `json:"configured"`
	Tier       string `json:"tier"`
}

type Product struct {
	Abi            string `json:"product_abi"`
	Arch           string `json:"product_arch"`
	Check          any    `json:"product_check"`
	Conflicts      string `json:"product_conflicts"`
	CopyrightOwner string `json:"product_copyright_owner"`
	CopyrightURL   string `json:"product_copyright_url"`
	CopyrightYears string `json:"product_copyright_years"`
	Email          string `json:"product_email"`
	Hash           string `json:"product_hash"`
	ID             string `json:"product_id"`
	Latest         string `json:"product_latest"`
	License        []any  `json:"product_license"`
	Log            int    `json:"product_log"`
	Mirror         string `json:"product_mirror"`
	Name           string `json:"product_name"`
	Nickname       string `json:"product_nickname"`
	Repos          string `json:"product_repos"`
	Series         string `json:"product_series"`
	Tier           string `json:"product_tier"`
	Time           string `json:"product_time"`
	Version        string `json:"product_version"`
	Website        string `json:"product_website"`
}

type ChangeLog struct {
	Series  string `json:"series"`
	Version string `json:"version"`
	Date    string `json:"date"`
}

type RunningStatus struct {
	Status string `json:"status"`
}

func (oc *Firmware) FirmwareGet() FirmwareData {
	var res FirmwareData
	body, err := utillis.GetRequest(oc.HostURL, oc.BasicToken, "api/core/firmware/get")
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func (oc *Firmware) FirmwareOptions() FirmwareOpts {
	var res FirmwareOpts
	body, err := utillis.GetRequest(oc.HostURL, oc.BasicToken, "api/core/firmware/getOptions")
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func (oc *Firmware) FirmwareInfo() FirmwareInfo {
	var res FirmwareInfo
	body, err := utillis.GetRequest(oc.HostURL, oc.BasicToken, "api/core/firmware/info")
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func (oc *Firmware) FirmwareRunning() RunningStatus {
	var res RunningStatus
	body, err := utillis.GetRequest(oc.HostURL, oc.BasicToken, "api/core/firmware/running")
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func (oc *DhcpV6) Reboot() error {
	return utillis.PostRequest(oc.HostURL, oc.BasicToken, "api/core/firmware/reboot")
}

func (oc *DhcpV6) PowerOff() error {
	return utillis.PostRequest(oc.HostURL, oc.BasicToken, "api/core/firmware/poweroff")
}
