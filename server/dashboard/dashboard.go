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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var baseTemplatePath = "./server/dashboard/templates/"
var funcs = template.FuncMap{"idtostring": func(value primitive.ObjectID) string { return value.Hex() }}

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
	tmpl, err := template.New("index.html").Funcs(funcs).ParseFiles(baseTemplatePath+"index.html",
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

	cssHandler := http.FileServer(http.Dir("./server/dashboard/assets/css/"))
	jsHandler := http.FileServer(http.Dir("./server/dashboard/assets/js/"))

	r.Handle("/assets/css/{asset}", http.StripPrefix("/assets/css/", cssHandler))
	r.Handle("/assets/js/{asset}", http.StripPrefix("/assets/js/", jsHandler))

	srv := &http.Server{
		Handler: r,
		Addr:    ":8093",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
