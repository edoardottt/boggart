package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

//Log defines the structure of a log record
//in the database
type Log struct {
	Ip        string              `json:"ip Str"`
	Method    string              `json:"method Str"`
	Path      string              `json:"path Str"`
	Headers   map[string][]string `json:"headers Str"`
	Body      string              `json:"body Str"`
	Timestamp int64               `json:"timestamp Int"`
}

//InsertLog inserts a log record into the logs collection
func InsertLog(client *mongo.Client, collection *mongo.Collection, record Log) {
	result, err := collection.InsertOne(context.TODO(), record) //result
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted: ", result.InsertedID)
}

/*
TEST

connString := os.Getenv("MONGO_CONN")
	//connString := "mongodb://hostname:27017"
	dbName := os.Getenv("DB_NAME")
	client, _ := db.ConnectDB(connString)

	database := db.GetDatabase(client, dbName)

	collection := db.GetLogs(database)

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	db.InsertLog(client, collection, db.Log{...})
*/
