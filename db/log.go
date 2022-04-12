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
	Ip        string `json:"ip Str"`
	Raw       string `json:"raw Str"`
	Timestamp int64  `json:"timestamp Int"`
}

//InsertLog inserts a log record into the logs collection
func InsertLog(client *mongo.Client, collection *mongo.Collection, record Log) {
	result, err := collection.InsertOne(context.TODO(), record) //result
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted: ", result.InsertedID)
}
