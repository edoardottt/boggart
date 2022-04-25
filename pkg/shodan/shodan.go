/*
=======================
	boggart
=======================

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with this program.  If not, see http://www.gnu.org/licenses/.

	@Repository:	https://github.com/edoardottt/boggart
	@Author:		edoardottt, https://www.edoardoottavianelli.it
	@License:		https://github.com/edoardottt/boggart/blob/main/LICENSE
*/

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
