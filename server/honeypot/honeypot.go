package honeypot

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/edoardottt/boggart/internal/file"
	"github.com/edoardottt/boggart/pkg/template"
	"github.com/gorilla/mux"
)

func genericWriter(w http.ResponseWriter, req *http.Request, response string) {
	fmt.Fprint(w, response)
}

func fileWriter(w http.ResponseWriter, req *http.Request, inputFile string) {
	content, err := file.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, content)
}

//Raw > to be filled
func Raw(tmpl template.Template) {

	r := mux.NewRouter()
	var staticPath = "public/honeypot/"

	//registering endpoints
	for _, request := range tmpl.Requests {
		if request.ID != "default" {
			request2 := request
			if request2.ResponseType == "raw" {

				r.HandleFunc(request2.Endpoint, func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Content-Type", request2.ContentType)
					genericWriter(w, r, request2.Content)
				}).Methods(template.HTTPMethodsAsString(request2.Methods)...)

			} else if request2.ResponseType == "file" {

				r.HandleFunc(request2.Endpoint, func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Content-Type", request2.ContentType)
					fileWriter(w, r, staticPath+request2.Content)
				}).Methods(template.HTTPMethodsAsString(request2.Methods)...)
			}
		}
	}

	//default response
	defaultRequest := template.Default(tmpl)
	if defaultRequest.ResponseType == "raw" {

		r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", defaultRequest.ContentType)
			genericWriter(w, r, defaultRequest.Content)
		})

	} else if defaultRequest.ResponseType == "file" {

		r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", defaultRequest.ContentType)
			fileWriter(w, r, staticPath+defaultRequest.Content)
		})
	}

	srv := &http.Server{
		Handler: r,
		Addr:    ":8090",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
