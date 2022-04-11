package file

import (
	"io/ioutil"
)

//ReadFile reads a file and returns the content
//of the inputted file.
func ReadFile(inputFile string) (string, error) {
	var content = []byte{}
	content, err := ioutil.ReadFile(inputFile)
	return string(content), err
}
