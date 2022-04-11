package honeypot

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/edoardottt/boggart/pkg/template"
	"github.com/gorilla/mux"
)

func genericWriter(w http.ResponseWriter, req *http.Request, response string) {
	fmt.Fprint(w, response)
}

//RawHoneypot
func RawHoneypot(tmpl template.Template) {

	r := mux.NewRouter()

	for _, request := range tmpl.Requests {
		if request.Id != "default" {
			request2 := request
			r.HandleFunc(request2.Endpoint, func(w http.ResponseWriter, r *http.Request) {
				genericWriter(w, r, request2.Content)
			}).Methods(template.HttpMethodsAsString(request2.Methods)...)
		}
	}
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, template.DefaultResponse(tmpl))
	})

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8090",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
