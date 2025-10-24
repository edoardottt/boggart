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
	@Author:		edoardottt, https://edoardottt.com
	@License:		https://github.com/edoardottt/boggart/blob/main/LICENSE
*/

package template

import (
	"fmt"
	"strings"

	"github.com/edoardottt/boggart/internal/slice"
)

// Type contains all the types that a template can assume.
type Type string

// Types a template can assume.
const (
	RawTemplateType    Type = "raw"
	ShodanTemplateType Type = "shodan"
)

const keyValueLen = 2

// HTTPMethod contains all the methods that a HTTP request can assume.
type HTTPMethod string

// HTTP Methods a request can have.
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

// ResponseType contains all the types that a HTTP response can assume.
type ResponseType string

// Types a response can have.
const (
	RawResponseType  ResponseType = "raw"
	FileResponseType ResponseType = "file"
)

// Request is the struct defining an HTTP request structure in a
// valid template.
type Request struct {
	ID           string       `yaml:"id"` // Id is mandatory
	Methods      []HTTPMethod `yaml:"methods,omitempty"`
	Endpoint     string       `yaml:"endpoint,omitempty"`
	ResponseType ResponseType `yaml:"response-type,omitempty"`
	ContentType  string       `yaml:"content-type,omitempty"`
	Content      string       `yaml:"content,omitempty"`
	Headers      []string     `yaml:"headers,omitempty"`
}

// Template is the struct defining the structure of a configuration template.
// The configuration file has to be a valid YAML file.
type Template struct {
	Type     Type      `yaml:"type,omitempty"`
	Requests []Request `yaml:"requests,omitempty"`
	Ignore   []string  `yaml:"ignore,omitempty"`
	IP       string    `yaml:"ip,omitempty"`
}

const DefaultID string = "default"

// ---------------------------------------
// -------------- HELPERS ----------------
// ---------------------------------------

// CheckTemplate checks if a generic template is formatted in a proper way.
func CheckTemplate(tmpl Template) error {
	if tmpl.Type == "" {
		return ErrMissingType
	}

	if tmpl.Type == RawTemplateType {
		return CheckRawTeplate(tmpl)
	}

	if tmpl.Type == "shodan" {
		return CheckShodanTemplate(tmpl)
	}

	return nil
}

// CheckRawTeplate checks if a raw template is formatted in a proper way.
func CheckRawTeplate(tmpl Template) error {
	if !IDUnique(tmpl) {
		return ErrUniqueRequestID
	}

	if !EndpointUnique(tmpl) {
		return ErrUniqueRequestEndpoint
	}

	if MissingTemplateDefault(tmpl) {
		return ErrMissingDefaultRequest
	}

	err := CheckRequests(tmpl)
	if err != nil {
		return err
	}

	err = CheckDefaultRequest(tmpl)
	if err != nil {
		return err
	}

	err = CheckIgnore(tmpl)
	if err != nil {
		return err
	}

	return nil
}

// CheckShodanTemplate checks if a shodan template is formatted in a proper way.
func CheckShodanTemplate(tmpl Template) error {
	if tmpl.IP != "" {
		return ErrMandatoryIP
	}

	return nil
}

// IDUnique checks if in a raw template there are
// duplicate request IDs.
// True for shodan template.
func IDUnique(tmpl Template) bool {
	if tmpl.Type == RawTemplateType {
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

// EndpointUnique checks if in a raw template there are
// duplicate request endpoints.
// True for shodan template.
func EndpointUnique(tmpl Template) bool {
	if tmpl.Type == RawTemplateType {
		keys := make(map[string]bool)
		list := []string{}

		for _, entry := range tmpl.Requests {
			if _, value := keys[entry.Endpoint]; !value {
				keys[entry.Endpoint] = true
				list = append(list, entry.Endpoint)
			}
		}

		return len(tmpl.Requests) == len(list)
	}

	return true
}

// MissingTemplateDefault checks if in a raw template there is
// a request with a default action.
// True for shodan template.
func MissingTemplateDefault(tmpl Template) bool {
	var missing = true

	if tmpl.Type == RawTemplateType {
		for _, entry := range tmpl.Requests {
			if entry.ID == DefaultID {
				missing = false
			}
		}
	} else {
		return false
	}

	return missing
}

// RootEndpointExists checks if a request handling for
// the root endpoint exists.
// True for shodan template.
func RootEndpointExists(tmpl Template) bool {
	if tmpl.Type == RawTemplateType {
		for _, entry := range tmpl.Requests {
			if entry.Endpoint == "/" {
				return true
			}
		}
	}

	return true
}

// Default returns the default response.
// Empty request for shodan template.
func Default(tmpl Template) Request {
	if tmpl.Type == RawTemplateType {
		for _, entry := range tmpl.Requests {
			if entry.ID == DefaultID {
				return entry
			}
		}
	}

	return Request{}
}

// HTTPMethodsAsString transforms a slice of HttpMethod to a
// slice of strings.
func HTTPMethodsAsString(methods []HTTPMethod) []string {
	result := make([]string, len(methods))
	for i, method := range methods {
		result[i] = string(method)
	}

	return result
}

// CheckRequests checks if the requests (except for default one)
// are ok. True if everything is correct.
// True for shodan template.
func CheckRequests(tmpl Template) error {
	for _, entry := range tmpl.Requests {
		if strings.Trim(entry.ID, " ") == "" {
			return ErrMissingID
		}

		if entry.ID != DefaultID {
			if strings.Trim(entry.Endpoint, " ") == "" {
				return fmt.Errorf("%w %s", ErrMissingEndpointID, entry.ID)
			}

			if len(entry.Methods) == 0 {
				return fmt.Errorf("%w %s", ErrMissingMethodsID, entry.ID)
			}

			if strings.Trim(string(entry.ResponseType), " ") == "" {
				return fmt.Errorf("%w %s", ErrMissingResponseTypeID, entry.ID)
			}

			if strings.Trim(entry.ContentType, " ") == "" {
				return fmt.Errorf("%w %s", ErrMissingContentTypeID, entry.ID)
			}

			if strings.Trim(entry.Content, " ") == "" {
				return fmt.Errorf("%w %s", ErrMissingContentID, entry.ID)
			}

			if len(entry.Headers) != 0 {
				for _, header := range entry.Headers {
					if len(strings.Split(header, ":")) != keyValueLen {
						return fmt.Errorf("%w %s", ErrWrongHeader, entry.ID)
					}
				}
			}
		}
	}

	return nil
}

// CheckDefaultRequest checks if the default request
// is ok. True if everything is correct.
// True for shodan template.
func CheckDefaultRequest(tmpl Template) error {
	entry := Default(tmpl)
	if strings.Trim(string(entry.ResponseType), " ") == "" {
		return ErrMissingDefaultResponseType
	}

	if strings.Trim(entry.ContentType, " ") == "" {
		return ErrMissingDefaultContentType
	}

	if strings.Trim(entry.Content, " ") == "" {
		return ErrMissingDefaultContent
	}

	return nil
}

// CheckIgnore checks if the ignore array
// is ok. True if everything is correct.
// True for shodan template.
func CheckIgnore(tmpl Template) error {
	input := tmpl.Ignore
	if len(input) == 0 {
		return nil
	}

	if len(input) != len(slice.RemoveDuplicateValues(input)) {
		return ErrDuplicatePathsIgnore
	}

	for _, path := range input {
		if path[0] != '/' {
			return ErrMissingSlashIgnore
		}
	}
	// here check if ignore is defined as endpoint in requests.
	for _, ignoreElem := range input {
		for _, request := range tmpl.Requests {
			if ignoreElem == request.Endpoint {
				return ErrPathRequestIgnore
			}
		}
	}

	return nil
}
