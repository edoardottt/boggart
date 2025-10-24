/*
=======================
	boggart
=======================

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with this program.  If not, see http://www.gnu.org/licenses/.

	@Repository:	https://github.com/edoardottt/boggart
	@Author:		edoardottt, https://edoardottt.com
	@License:		https://github.com/edoardottt/boggart/blob/main/LICENSE
*/

package honeypot

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/edoardottt/boggart/db"
	"github.com/edoardottt/boggart/internal/file"
	"github.com/edoardottt/boggart/pkg/template"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ContextBackgroundDuration = 3
	WriteTimeoutDuration      = 15
	ReadTimeoutDuration       = 15
)

func genericWriter(w http.ResponseWriter, req *http.Request, dbName string,
	client *mongo.Client, tmpl template.Template, response string) {
	if !ignorePath(req.URL.Path, tmpl) {
		ctx, cancel := context.WithTimeout(context.Background(), ContextBackgroundDuration*time.Second)
		defer cancel()

		database := db.GetDatabase(client, dbName)
		collection := db.GetLogs(database)

		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err)

			return
		}

		bodyString := string(bodyBytes)
		remoteIP := strings.Split(req.RemoteAddr, ":")[0]

		db.InsertLog(ctx, client, collection, db.Log{
			IP:        remoteIP,
			Method:    req.Method,
			Path:      req.RequestURI,
			Headers:   req.Header,
			Body:      bodyString,
			Timestamp: time.Now().Unix(),
		})
	}

	fmt.Fprint(w, response)
}

func fileWriter(w http.ResponseWriter, req *http.Request, dbName string,
	client *mongo.Client, tmpl template.Template, inputFile string) {
	content, err := file.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	if !ignorePath(req.URL.Path, tmpl) {
		ctx, cancel := context.WithTimeout(context.Background(), ContextBackgroundDuration*time.Second)
		defer cancel()

		database := db.GetDatabase(client, dbName)
		collection := db.GetLogs(database)

		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err)

			return
		}

		bodyString := string(bodyBytes)
		remoteIP := strings.Split(req.RemoteAddr, ":")[0]

		db.InsertLog(ctx, client, collection, db.Log{
			IP:        remoteIP,
			Method:    req.Method,
			Path:      req.RequestURI,
			Headers:   req.Header,
			Body:      bodyString,
			Timestamp: time.Now().Unix(),
		})
	}

	fmt.Fprint(w, content)
}

// Raw > to be filled.
func Raw(tmpl template.Template) {
	// DB setup
	connString := os.Getenv("MONGO_CONN") // "mongodb://hostname:27017".
	dbName := os.Getenv("MONGO_DB")
	client, err := db.ConnectDB(connString)
	// ------- debug -------.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("HONEYPOT: Connected to MongoDB!")

	// Routes setup.
	router := mux.NewRouter()

	var staticPath = "public/honeypot/"

	// registering endpoints.
	for _, request := range tmpl.Requests {
		if request.ID != "default" {
			request2 := request
			switch request2.ResponseType {
			case "raw":
				router.HandleFunc(request2.Endpoint, func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Content-Type", request2.ContentType)

					// Add custom headers.
					for _, elem := range request2.Headers {
						values := strings.Split(elem, ":")
						w.Header().Add(values[0], values[1])
					}

					genericWriter(w, r, dbName, client, tmpl, request2.Content)
				}).Methods(template.HTTPMethodsAsString(request2.Methods)...)
			case "file":
				router.HandleFunc(request2.Endpoint, func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Content-Type", request2.ContentType)

					// Add custom headers.
					for _, elem := range request2.Headers {
						values := strings.Split(elem, ":")
						w.Header().Add(values[0], values[1])
					}

					fileWriter(w, r, dbName, client, tmpl, staticPath+request2.Content)
				}).Methods(template.HTTPMethodsAsString(request2.Methods)...)
			}
		}
	}

	// default response.
	defaultRequest := template.Default(tmpl)
	switch defaultRequest.ResponseType {
	case "raw":
		router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", defaultRequest.ContentType)

			// Add custom headers.
			for _, elem := range defaultRequest.Headers {
				values := strings.Split(elem, ":")
				w.Header().Add(values[0], values[1])
			}

			genericWriter(w, r, dbName, client, tmpl, defaultRequest.Content)
		})
	case "file":
		router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", defaultRequest.ContentType)

			// Add custom headers.
			for _, elem := range defaultRequest.Headers {
				values := strings.Split(elem, ":")
				w.Header().Add(values[0], values[1])
			}

			fileWriter(w, r, dbName, client, tmpl, staticPath+defaultRequest.Content)
		})
	}

	srv := &http.Server{
		Handler: router,
		Addr:    ":8092",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: WriteTimeoutDuration * time.Second,
		ReadTimeout:  ReadTimeoutDuration * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// ignorePath.
func ignorePath(path string, tmpl template.Template) bool {
	for _, elem := range tmpl.Ignore {
		if elem == path {
			return true
		}
	}

	return false
}
