package main

import (
	"log"
)

func main() {
	tmpl, err := ReadTemplate("../config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Default().Println(tmpl)
}
