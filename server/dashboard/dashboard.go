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
	@Author:		edoardottt, https://www.edoardoottavianelli.it
	@License:		https://github.com/edoardottt/boggart/blob/main/LICENSE
*/

package dashboard

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/edoardottt/boggart/db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ContextBackgroundDuration = 3
	WriteTimeoutDuration      = 15
	ReadTimeoutDuration       = 15
	LatestNLogs               = 30
	baseTemplatePath          = "./server/dashboard/templates/"
)

// Start starts the dashboard.
func Start() {
	// DB setup.
	connString := os.Getenv("MONGO_CONN") // "mongodb://hostname:27017".
	dbName := os.Getenv("MONGO_DB")
	client, err := db.ConnectDB(connString)
	// ------- debug -------.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DASHBOARD: Connected to MongoDB!")

	var funcs = template.FuncMap{
		"idtostring": func(value primitive.ObjectID) string {
			return value.Hex()
		},
		"maptostring": func(input map[string][]string) string {
			var result = ""
			for k, v := range input {
				result += html.EscapeString(k) + ": "
				var safeValues []string
				for _, x := range v {
					safeValues = append(safeValues, html.EscapeString(x))
				}
				result += strings.Join(safeValues, ",") + "<br>"
			}
			return result
		},
		"timetostring": func(input int64) string {
			return time.Unix(input, 0).Format("01-02-2006 15:04:05")
		},
		"escapehtml": func(input string) string {
			return html.EscapeString(input)
		},
	}

	// Routes setup.
	router := mux.NewRouter()

	tmpl, err := template.New("index.html").Funcs(funcs).ParseFiles(baseTemplatePath+"index.html",
		baseTemplatePath+"head.html",
		baseTemplatePath+"footer.html")
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dashboardIndexHandler(w, tmpl)
	})

	tmplOverview, err := template.New("overview.html").Funcs(funcs).ParseFiles(baseTemplatePath+"overview.html",
		baseTemplatePath+"head.html",
		baseTemplatePath+"footer.html")
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/overview", func(w http.ResponseWriter, r *http.Request) {
		dashboardOverviewHandler(w, client, dbName, tmplOverview)
	})

	tmplQuery, err := template.New("query.html").Funcs(funcs).ParseFiles(baseTemplatePath+"query.html",
		baseTemplatePath+"head.html",
		baseTemplatePath+"footer.html")
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		dashboardQueryHandler(w, tmplQuery)
	})

	tmplResult, err := template.New("result.html").Funcs(funcs).ParseFiles(baseTemplatePath+"result.html",
		baseTemplatePath+"head.html",
		baseTemplatePath+"footer.html")
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		dashboardResultHandler(r, w, client, dbName, tmplResult)
	})

	tmplLatest, err := template.New("latest.html").Funcs(funcs).ParseFiles(baseTemplatePath+"latest.html",
		baseTemplatePath+"head.html",
		baseTemplatePath+"footer.html")
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/latest", func(w http.ResponseWriter, r *http.Request) {
		dashboardLatestHandler(w, client, dbName, tmplLatest)
	})

	tmplID, err := template.New("id.html").Funcs(funcs).ParseFiles(baseTemplatePath+"id.html",
		baseTemplatePath+"head.html",
		baseTemplatePath+"footer.html")
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/id/{id}", func(w http.ResponseWriter, r *http.Request) {
		dashboardIDHandler(w, client, dbName, tmplID, mux.Vars(r)["id"])
	})

	tmplDetection, err := template.New("detection.html").Funcs(funcs).ParseFiles(baseTemplatePath+"detection.html",
		baseTemplatePath+"head.html",
		baseTemplatePath+"footer.html")
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/detection", func(w http.ResponseWriter, r *http.Request) {
		dashboardDetectionHandler(w, tmplDetection)
	})

	tmplStatus, err := template.New("status.html").Funcs(funcs).ParseFiles(baseTemplatePath+"status.html",
		baseTemplatePath+"head.html",
		baseTemplatePath+"footer.html")
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		dashboardStatusHandler(w, tmplStatus)
	})

	cssHandler := http.FileServer(http.Dir("./server/dashboard/assets/css/"))
	jsHandler := http.FileServer(http.Dir("./server/dashboard/assets/js/"))

	router.Handle("/assets/css/{asset}", http.StripPrefix("/assets/css/", cssHandler))
	router.Handle("/assets/js/{asset}", http.StripPrefix("/assets/js/", jsHandler))

	srv := &http.Server{
		Handler: router,
		Addr:    ":8093",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: WriteTimeoutDuration * time.Second,
		ReadTimeout:  ReadTimeoutDuration * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
