package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

//ReadTemplate gets as input a filename and returns a Template object.
func ReadTemplate(filename string) (Template, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return Template{}, err
	}

	c := Template{}
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		return Template{}, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}
