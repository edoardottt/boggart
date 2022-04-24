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

package template

import (
	"errors"
	"strings"
)

//Type contains all the types that a template can assume
type Type string

//Types a template can assume
const (
	RawTemplateType    Type = "raw"
	ShodanTemplateType Type = "shodan"
)

//HTTPMethod contains all the methods that a HTTP request can assume
type HTTPMethod string

//HTTP Methods a request can have
const (
	MethodGet     HTTPMethod = "GET"
	MethodHead    HTTPMethod = "HEAD"
	MethodPost    HTTPMethod = "POST"
	MethodPut     HTTPMethod = "PUT"
	MethodPatch   HTTPMethod = "PATCH" // RFC 5789
	MethodDelete  HTTPMethod = "DELETE"
	MethodConnect HTTPMethod = "CONNECT"
	MethodOptions HTTPMethod = "OPTIONS"
	MethodTrace   HTTPMethod = "TRACE"
)

//ResponseType contains all the types that a HTTP response can assume
type ResponseType string

//Types a response can have
const (
	RawResponseType  ResponseType = "raw"
	FileResponseType ResponseType = "file"
)

//Request is the struct defining an HTTP request structure in a
//valid template
type Request struct {
	ID           string       `yaml:"id"` //Id is mandatory
	Methods      []HTTPMethod `yaml:"methods,omitempty"`
	Endpoint     string       `yaml:"endpoint,omitempty"`
	ResponseType ResponseType `yaml:"response-type,omitempty"`
	ContentType  string       `yaml:"content-type,omitempty"`
	Content      string       `yaml:"content,omitempty"`
}

//Template is the struct defining the structure of a configuration template.
//The configuration file has to be a valid YAML file.
type Template struct {
	Type     Type      `yaml:"type,omitempty"`
	Requests []Request `yaml:"requests,omitempty"`
	IP       string    `yaml:"ip,omitempty"`
}

//CheckTemplate checks if a generic template is formatted in a proper way.
func CheckTemplate(tmpl Template) (bool, error) {
	if tmpl.Type == "" {
		return false, errors.New("template: missing template type")
	}
	if tmpl.Type == "raw" {
		return CheckRawTeplate(tmpl)
	}
	if tmpl.Type == "shodan" {
		return CheckShodanTemplate(tmpl), nil
	}
	return true, nil
}

//CheckRawTeplate checks if a raw template is formatted in a proper way.
func CheckRawTeplate(tmpl Template) (bool, error) {
	if !IDUnique(tmpl) {
		return false, errors.New("template: request IDs are not unique")
	}
	if MissingTemplateDefault(tmpl) {
		return false, errors.New("template: missing default request")
	}
	_, err := CheckRequests(tmpl)
	if err != nil {
		return false, err
	}
	_, err = CheckDefaultRequest(tmpl)
	if err != nil {
		return false, err
	}
	return true, nil
}

//CheckShodanTemplate checks if a shodan template is formatted in a proper way.
func CheckShodanTemplate(tmpl Template) bool {
	return tmpl.IP != ""
}

//IDUnique checks if in a raw template there are
//duplicate request IDs.
//True for shodan template
func IDUnique(tmpl Template) bool {
	if tmpl.Type == "raw" {
		keys := make(map[string]bool)
		list := []string{}
		for _, entry := range tmpl.Requests {
			if _, value := keys[entry.ID]; !value {
				keys[entry.ID] = true
				list = append(list, entry.ID)
			}
		}
		return len(tmpl.Requests) == len(list)
	}
	return true
}

//MissingTemplateDefault checks if in a raw template there is
//a request with a default action.
//True for shodan template
func MissingTemplateDefault(tmpl Template) bool {
	var missing = true
	if tmpl.Type == "raw" {
		for _, entry := range tmpl.Requests {
			if entry.ID == "default" {
				missing = false
			}
		}
	} else {
		return false
	}
	return missing
}

//RootEndpointExists checks if a request handling for
//the root endpoint exists.
//True for shodan template
func RootEndpointExists(tmpl Template) bool {
	if tmpl.Type == "raw" {
		for _, entry := range tmpl.Requests {
			if entry.Endpoint == "/" {
				return true
			}
		}
	}
	return true
}

//Default returns the default response.
//Empty request for shodan template
func Default(tmpl Template) Request {
	if tmpl.Type == "raw" {
		for _, entry := range tmpl.Requests {
			if entry.ID == "default" {
				return entry
			}
		}
	}
	return Request{}
}

//HTTPMethodsAsString transforms a slice of HttpMethod to a
//slice of strings.
func HTTPMethodsAsString(methods []HTTPMethod) []string {
	var result []string
	for _, method := range methods {
		result = append(result, string(method))
	}
	return result
}

//CheckRequests checks if the requests (except for default one)
//are ok. True if everything is correct.
//True for shodan template
func CheckRequests(tmpl Template) (bool, error) {
	for _, entry := range tmpl.Requests {
		if strings.Trim(entry.ID, " ") == "" {
			return false, errors.New("template: missing id in request")
		}
		if entry.ID != "default" {
			if strings.Trim(entry.Endpoint, " ") == "" {
				return false, errors.New("template: missing endpoint in request with id " + entry.ID)
			}
			if len(entry.Methods) == 0 {
				return false, errors.New("template: missing methods in request with id " + entry.ID)
			}
			if strings.Trim(string(entry.ResponseType), " ") == "" {
				return false, errors.New("template: missing response type in request with id " + entry.ID)
			}
			if strings.Trim(entry.ContentType, " ") == "" {
				return false, errors.New("template: missing content type in request with id " + entry.ID)
			}
			if strings.Trim(entry.Content, " ") == "" {
				return false, errors.New("template: missing content in request with id " + entry.ID)
			}
		}
	}
	return true, nil
}

//CheckDefaultRequest checks if the default request
//is ok. True if everything is correct.
//True for shodan template
func CheckDefaultRequest(tmpl Template) (bool, error) {
	entry := Default(tmpl)
	if strings.Trim(string(entry.ResponseType), " ") == "" {
		return false, errors.New("template: missing response type in default request")
	}
	if strings.Trim(entry.ContentType, " ") == "" {
		return false, errors.New("template: missing content type in default request")
	}
	if strings.Trim(entry.Content, " ") == "" {
		return false, errors.New("template: missing content in default request")
	}
	return true, nil
}
