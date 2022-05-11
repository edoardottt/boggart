package dashboard

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/edoardottt/boggart/db"
	"github.com/gorilla/mux"
)

//Start starts the dashboard >
func Start() {

	// DB setup
	connString := os.Getenv("MONGO_CONN") // "mongodb://hostname:27017"
	//dbName := os.Getenv("MONGO_DB")
	client := db.ConnectDB(connString)
	// ------- debug -------
	if client != nil {
		fmt.Println("DASHBOARD: Connected to MongoDB!")
	}
	baseTemplatePath := "./server/dashboard/templates/"

	// Routes setup
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("assets"))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(baseTemplatePath + "index.tmpl")
		if err != nil {
			log.Fatal(err)
		}
		buf := &bytes.Buffer{}
		err = tmpl.Execute(buf, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		buf.WriteTo(w)
	})

	r.Handle("/assets/", http.StripPrefix("/assets/", fs))

	srv := &http.Server{
		Handler: r,
		Addr:    ":8093",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
