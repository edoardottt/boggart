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

// GetShodanAPIKey returns the Shodan Api Key.
func GetShodanAPIKey() (string, error) {
	apiKey := os.Getenv("SHODAN_KEY")
	if strings.Trim(apiKey, " ") == "" {
		return "", ErrAPIKeyEmpty
	}

	return apiKey, nil
}

// APIInfoResponse defines the structure of the Info
// Api response.
type APIInfoResponse struct {
	ScanCredits int `json:"scanCredits"`
	UsageLimits struct {
		ScanCredits  int `json:"scanCredits"`
		QueryCredits int `json:"queryCredits"`
		MonitoredIps int `json:"monitoredIps"`
	} `json:"usageLimits"`
	Plan         string `json:"plan"`
	HTTPS        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
	QueryCredits int    `json:"queryCredits"`
	MonitoredIps int    `json:"monitoredIps"`
	UnlockedLeft int    `json:"unlockedLeft"`
	Telnet       bool   `json:"telnet"`
}

// APIInfo returns the struct ApiInfoResponse filled
// with the Api account information.
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

// HostIPInfo > TODO.
func HostIPInfo(hostIP string, apiKey string) {
	url := baseURL + hostIPEndpoint + hostIP + "?key=" + apiKey
	fmt.Println(url)
}

// Host.
type Host struct {
	RegionCode  string        `json:"regionCode"`
	Tags        []interface{} `json:"tags"`
	IP          int64         `json:"ip"`
	AreaCode    interface{}   `json:"areaCode"`
	Domains     []string      `json:"domains"`
	Hostnames   []string      `json:"hostnames"`
	CountryCode string        `json:"countryCode"`
	Org         string        `json:"org"`
	Data        []struct {
		IP     int64 `json:"ip"`
		Shodan struct {
			ID      string `json:"id"`
			Ptr     bool   `json:"ptr"`
			Options struct {
			} `json:"options"`
			Module  string `json:"module"`
			Crawler string `json:"crawler"`
		} `json:"shodan"`
		Product string `json:"product"`
		HTTP    struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robotsHash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemapHash"`
			Robots      interface{}   `json:"robots"`
			Server      string        `json:"server"`
			HeadersHash int           `json:"headersHash"`
			Host        string        `json:"host"`
			HTML        string        `json:"html"`
			Location    string        `json:"location"`
			Components  struct {
			} `json:"components"`
			SecuritytxtHash interface{} `json:"securitytxtHash"`
			Sitemap         interface{} `json:"sitemap"`
			HTMLHash        int         `json:"htmlHash"`
		} `json:"http,omitempty"`
		Os   interface{} `json:"os"`
		Opts struct {
		} `json:"opts"`
		Timestamp string   `json:"timestamp"`
		Isp       string   `json:"isp"`
		Cpe       []string `json:"cpe"`
		IPStr     string   `json:"ipStr"`
		Asn       string   `json:"asn"`
		Hostnames []string `json:"hostnames"`
		Cpe23     []string `json:"cpe23"`
		Org       string   `json:"org"`
		Domains   []string `json:"domains"`
		Hash      int      `json:"hash"`
		Data      string   `json:"data"`
		Port      int      `json:"port"`
		Transport string   `json:"transport"`
		Location  struct {
			City        string      `json:"city"`
			RegionCode  string      `json:"regionCode"`
			AreaCode    interface{} `json:"areaCode"`
			Longitude   float64     `json:"longitude"`
			CountryName string      `json:"countryName"`
			CountryCode string      `json:"countryCode"`
			Latitude    float64     `json:"latitude"`
		} `json:"location"`
	} `json:"data"`
	Asn         string      `json:"asn"`
	City        string      `json:"city"`
	Latitude    float64     `json:"latitude"`
	Isp         string      `json:"isp"`
	Longitude   float64     `json:"longitude"`
	LastUpdate  string      `json:"lastUpdate"`
	CountryName string      `json:"countryName"`
	IPStr       string      `json:"ipStr"`
	Os          interface{} `json:"os"`
	Ports       []int       `json:"ports"`
}
