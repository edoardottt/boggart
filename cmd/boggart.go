package main

import (
	"log"

	"github.com/edoardottt/boggart/pkg/template"
	"github.com/edoardottt/boggart/server/honeypot"
)

func main() {
	tmpl, err := ReadTemplate("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	//Check Template format
	if _, err := template.CheckTemplate(tmpl); err != nil {
		log.Fatal(err)
	}
	honeypot.Raw(tmpl)
}
