package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

//ReadTemplate gets as input a filename and returns a Template object.
//The filename should be a YAML file.
//To check if the template is valid YAML check the error.
func ReadTemplate(filename string) (Template, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return Template{}, err
	}

	result := Template{}
	err = yaml.Unmarshal(buf, &result)
	if err != nil {
		return Template{}, fmt.Errorf("in file %q: %v", filename, err)
	}

	return result, nil
}

//TemplateIdUnique checks if in a raw template there are
//duplicate request IDs.
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
