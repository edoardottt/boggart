package honeypot

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/edoardottt/boggart/db"
	"github.com/edoardottt/boggart/internal/file"
	"github.com/edoardottt/boggart/pkg/template"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func genericWriter(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client, response string) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	database := db.GetDatabase(client, dbName)
	collection := db.GetLogs(database)

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	db.InsertLog(client, collection, ctx, db.Log{
		IP:        req.RemoteAddr,
		Method:    req.Method,
		Path:      req.RequestURI,
		Headers:   req.Header,
		Body:      bodyString,
		Timestamp: time.Now().Unix(),
	})

	/* =========== DEBUG ============
	insertedID := db.InsertLog(client, collection, ctx, db.Log{
		IP:        req.RemoteAddr,
		Method:    req.Method,
		Path:      req.RequestURI,
		Headers:   req.Header,
		Body:      bodyString,
		Timestamp: time.Now().Unix(),
	})
	fmt.Println("Inserted: ", insertedID)

	insertedLog, err := db.GetLogByID(client, collection, ctx, insertedID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(insertedLog)
	*/

	fmt.Fprint(w, response)
}

func fileWriter(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client, inputFile string) {
	content, err := file.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	database := db.GetDatabase(client, dbName)
	collection := db.GetLogs(database)

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	db.InsertLog(client, collection, ctx, db.Log{
		IP:        req.RemoteAddr,
		Method:    req.Method,
		Path:      req.RequestURI,
		Headers:   req.Header,
		Body:      bodyString,
		Timestamp: time.Now().Unix(),
	})

	fmt.Fprint(w, content)
}

//Raw > to be filled
func Raw(tmpl template.Template) {

	// DB setup
	connString := os.Getenv("MONGO_CONN") // "mongodb://hostname:27017"
	dbName := os.Getenv("DB_NAME")
	client, _ := db.ConnectDB(connString)

	// Routes setup
	r := mux.NewRouter()
	var staticPath = "public/honeypot/"

	//registering endpoints
	for _, request := range tmpl.Requests {
		if request.ID != "default" {
			request2 := request
			if request2.ResponseType == "raw" {

				r.HandleFunc(request2.Endpoint, func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Content-Type", request2.ContentType)
					genericWriter(w, r, dbName, client, request2.Content)
				}).Methods(template.HTTPMethodsAsString(request2.Methods)...)

			} else if request2.ResponseType == "file" {

				r.HandleFunc(request2.Endpoint, func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Content-Type", request2.ContentType)
					fileWriter(w, r, dbName, client, staticPath+request2.Content)
				}).Methods(template.HTTPMethodsAsString(request2.Methods)...)
			}
		}
	}

	//default response
	defaultRequest := template.Default(tmpl)
	if defaultRequest.ResponseType == "raw" {

		r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", defaultRequest.ContentType)
			genericWriter(w, r, dbName, client, defaultRequest.Content)
		})

	} else if defaultRequest.ResponseType == "file" {

		r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", defaultRequest.ContentType)
			fileWriter(w, r, dbName, client, staticPath+defaultRequest.Content)
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
