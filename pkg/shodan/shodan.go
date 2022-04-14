package shodan

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/*
reference: https://developer.shodan.io/api
*/

const (
	baseUrl        = "https://api.shodan.io"
	hostIpEndpoint = "/shodan/host/"
	apiInfo        = "/api-info"
)

//GetShodanApiKey returns the Shodan Api Key
func GetShodanApiKey() string {
	return os.Getenv("SHODAN_KEY")
}

//ApiInfoResponse defines the structure of the Info
//Api response
type ApiInfoResponse struct {
	ScanCredits int `json:"scan_credits"`
	UsageLimits struct {
		ScanCredits  int `json:"scan_credits"`
		QueryCredits int `json:"query_credits"`
		MonitoredIps int `json:"monitored_ips"`
	} `json:"usage_limits"`
	Plan         string `json:"plan"`
	Https        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
	QueryCredits int    `json:"query_credits"`
	MonitoredIps int    `json:"monitored_ips"`
	UnlockedLeft int    `json:"unlocked_left"`
	Telnet       bool   `json:"telnet"`
}

//ShodanApiInfo returns the struct ApiInfoResponse filled
//with the Api account information
func ApiInfo(apiKey string) ApiInfoResponse {
	resp, err := http.Get(baseUrl + apiInfo + "?=" + apiKey)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response ApiInfoResponse
	json.Unmarshal([]byte(body), &response)
	return response
}
