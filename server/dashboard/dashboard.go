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
	"go.mongodb.org/mongo-driver/mongo"
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
	}

	tmpl, err := template.New("index.html").Funcs(funcs).ParseFiles(baseTemplatePath+"index.html",
		baseTemplatePath+"navbar.html",
		baseTemplatePath+"latest.html",
		baseTemplatePath+"footer.html")
	if err != nil {
		log.Fatal(err)
	}

	tmplID, err := template.New("id.html").Funcs(funcs).ParseFiles(baseTemplatePath+"id.html",
		baseTemplatePath+"navbar.html",
		baseTemplatePath+"latest.html",
		baseTemplatePath+"footer.html")
	if err != nil {
		log.Fatal(err)
	}

	// Routes setup.
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dashboardHandler(w, client, dbName, tmpl)
	})

	router.HandleFunc("/id/{id}", func(w http.ResponseWriter, r *http.Request) {
		dashboardIDHandler(w, client, dbName, tmplID, mux.Vars(r)["id"])
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

func dashboardHandler(w http.ResponseWriter, client *mongo.Client, dbName string,
	tmpl *template.Template) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextBackgroundDuration*time.Second)
	defer cancel()

	database := db.GetDatabase(client, dbName)
	collection := db.GetLogs(database)

	logs, err := db.GetLatestNLogs(ctx, client, collection, LatestNLogs)
	if err != nil {
		fmt.Println(err)

		return
	}

	buf := &bytes.Buffer{}

	err = tmpl.Execute(buf, logs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func dashboardIDHandler(w http.ResponseWriter, client *mongo.Client, dbName string,
	tmpl *template.Template, id string) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextBackgroundDuration*time.Second)
	defer cancel()

	database := db.GetDatabase(client, dbName)
	collection := db.GetLogs(database)
	logID, err := db.GetLogByID(ctx, client, collection, id)

	if err != nil {
		fmt.Println(err)
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, logID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
