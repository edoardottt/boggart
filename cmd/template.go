package main

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

//Request is the struct defining an HTTP request
type Request struct {
	Id           string       `yaml:"id,omitempty"`
	Method       HttpMethod   `yaml:"method,omitempty"`
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
