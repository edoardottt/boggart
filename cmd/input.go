package main

import (
	"fmt"
	"io/ioutil"

	"github.com/edoardottt/boggart/pkg/template"

	"gopkg.in/yaml.v3"
)

//ReadTemplate gets as input a filename and returns a Template object.
//The filename should be a YAML file.
//To check if the template is valid YAML check the error.
func ReadTemplate(filename string) (template.Template, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return template.Template{}, err
	}

	result := template.Template{}
	err = yaml.Unmarshal(buf, &result)
	if err != nil {
		return template.Template{}, fmt.Errorf("in file %q: %v", filename, err)
	}

	return result, nil
}
