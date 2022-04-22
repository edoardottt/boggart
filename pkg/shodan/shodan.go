package shodan

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
reference: https://developer.shodan.io/api
*/

const (
	baseURL        = "https://api.shodan.io"
	hostIPEndpoint = "/shodan/host/"
	apiInfo        = "/api-info"
)

//GetShodanAPIKey returns the Shodan Api Key
func GetShodanAPIKey() (string, error) {
	apiKey := os.Getenv("SHODAN_KEY")
	if strings.Trim(apiKey, " ") == "" {
		return "", errors.New("shodan: Api key is empty")
	}
	return apiKey, nil
}

//APIInfoResponse defines the structure of the Info
//Api response
type APIInfoResponse struct {
	ScanCredits int `json:"scan_credits"`
	UsageLimits struct {
		ScanCredits  int `json:"scan_credits"`
		QueryCredits int `json:"query_credits"`
		MonitoredIps int `json:"monitored_ips"`
	} `json:"usage_limits"`
	Plan         string `json:"plan"`
	HTTPS        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
	QueryCredits int    `json:"query_credits"`
	MonitoredIps int    `json:"monitored_ips"`
	UnlockedLeft int    `json:"unlocked_left"`
	Telnet       bool   `json:"telnet"`
}

//APIInfo returns the struct ApiInfoResponse filled
//with the Api account information
func APIInfo(apiKey string) APIInfoResponse {
	resp, err := http.Get(baseURL + apiInfo + "?key=" + apiKey)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response APIInfoResponse
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		log.Fatal(err)
	}
	return response
}

//HostIPInfo > TODO
func HostIPInfo(hostIP string, apiKey string) {
	url := baseURL + hostIPEndpoint + hostIP + "?key=" + apiKey
	fmt.Println(url)
}
