package template

//TemplateType contains all the types that a template can assume
type TemplateType string

const (
	RawTemplateType    TemplateType = "raw"
	ShodanTemplateType TemplateType = "shodan"
)

//HttpMethod contains all the methods that a HTTP request can assume
type HttpMethod string

const (
	MethodGet     HttpMethod = "GET"
	MethodHead    HttpMethod = "HEAD"
	MethodPost    HttpMethod = "POST"
	MethodPut     HttpMethod = "PUT"
	MethodPatch   HttpMethod = "PATCH" // RFC 5789
	MethodDelete  HttpMethod = "DELETE"
	MethodConnect HttpMethod = "CONNECT"
	MethodOptions HttpMethod = "OPTIONS"
	MethodTrace   HttpMethod = "TRACE"
)

//ResponseType contains all the types that a HTTP response can assume
type ResponseType string

const (
	RawResponseType  ResponseType = "raw"
	FileResponseType ResponseType = "file"
)

//Request is the struct defining an HTTP request structure in a
//valid template
type Request struct {
	Id           string       `yaml:"id"` //Id is mandatory
	Methods      []HttpMethod `yaml:"methods,omitempty"`
	Endpoint     string       `yaml:"endpoint,omitempty"`
	ResponseType ResponseType `yaml:"response-type,omitempty"`
	ContentType  string       `yaml:"content-type,omitempty"`
	Content      string       `yaml:"content,omitempty"`
}

//Template is the struct defining the structure of a configuration template.
//The configuration file has to be a valid YAML file.
type Template struct {
	Type     TemplateType `yaml:"type,omitempty"`
	Requests []Request    `yaml:"requests,omitempty"`
	Ip       string       `yaml:"ip,omitempty"`
}

//CheckTemplate checks if a generic template is formatted in a proper way.
func CheckTemplate(tmpl Template) bool {
	if tmpl.Type == "" {
		return false
	}
	if tmpl.Type == "raw" {
		return CheckRawTeplate(tmpl)
	}
	if tmpl.Type == "shodan" {
		return CheckShodanTemplate(tmpl)
	}
	return false
}

//CheckRawTeplate checks if a raw template is formatted in a proper way.
func CheckRawTeplate(tmpl Template) bool {
	if !TemplateIdUnique(tmpl) {
		return false
	}
	if MissingTemplateDefault(tmpl) {
		return false
	}
	return true
}

//CheckShodanTemplate checks if a shodan template is formatted in a proper way.
func CheckShodanTemplate(tmpl Template) bool {
	return tmpl.Ip != ""
}

//TemplateIdUnique checks if in a raw template there are
//duplicate request IDs.
//True for shodan template
func TemplateIdUnique(tmpl Template) bool {
	if tmpl.Type == "raw" {
		keys := make(map[string]bool)
		list := []string{}
		for _, entry := range tmpl.Requests {
			if _, value := keys[entry.Id]; !value {
				keys[entry.Id] = true
				list = append(list, entry.Id)
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
			if entry.Id == "default" {
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

//DefaultResponse returns the default response.
//Empty string for shodan template
func DefaultResponse(tmpl Template) string {
	if tmpl.Type == "raw" {
		for _, entry := range tmpl.Requests {
			if entry.Id == "default" {
				return entry.Content
			}
		}
	}
	return ""
}

//HttpMethodsAsString transforms a slice of HttpMethod to a
//slice of strings.
func HttpMethodsAsString(methods []HttpMethod) []string {
	var result []string
	for _, method := range methods {
		result = append(result, string(method))
	}
	return result
}
