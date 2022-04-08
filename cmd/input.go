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
