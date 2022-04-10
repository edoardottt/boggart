package honeypot

import (
	"fmt"
	"net/http"
)

func genericWriter(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

//Honeypot
func Honeypot() {

	http.HandleFunc("/", genericWriter)

	http.ListenAndServe(":8090", nil)
}
