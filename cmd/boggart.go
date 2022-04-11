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
	if !template.CheckTemplate(tmpl) {
		log.Fatal("wrong template format!")
	}
	honeypot.RawHoneypot(tmpl)
}
