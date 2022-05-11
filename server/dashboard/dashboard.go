package dashboard

import (
	"bytes"
	"context"
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
	dbName := os.Getenv("MONGO_DB")
	client := db.ConnectDB(connString)
	// ------- debug -------
	if client != nil {
		fmt.Println("DASHBOARD: Connected to MongoDB!")
	}
	baseTemplatePath := "./server/dashboard/templates/"
	tmpl, err := template.ParseFiles(baseTemplatePath+"index.html",
		baseTemplatePath+"navbar.html",
		baseTemplatePath+"latest.html",
		baseTemplatePath+"footer.html")
	if err != nil {
		log.Fatal(err)
	}

	// Routes setup
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		database := db.GetDatabase(client, dbName)
		collection := db.GetLogs(database)
		logs, err := db.GetLatestNLogs(client, collection, ctx, 30)
		if err != nil {
			log.Fatal(err)
		}

		buf := &bytes.Buffer{}
		err = tmpl.Execute(buf, logs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = buf.WriteTo(w)
		if err != nil {
			log.Fatal(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fs := http.FileServer(http.Dir("assets"))
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
