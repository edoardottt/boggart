package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Log defines the structure of a log record
//in the database
type Log struct {
	IP        string              `json:"ip"`
	Method    string              `json:"method"`
	Path      string              `json:"path"`
	Headers   map[string][]string `json:"headers"`
	Body      string              `json:"body"`
	Timestamp int64               `json:"timestamp"`
}

//InsertLog inserts a log record into the logs collection
func InsertLog(client *mongo.Client, collection *mongo.Collection, ctx context.Context, record Log) interface{} {
	result, err := collection.InsertOne(ctx, record) //result
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted: ", result.InsertedID)
	return result.InsertedID
}

//GetLogByID returns a log struct with the defined ID.
//If the ID is not present in the database err won't be nil.
func GetLogByID(client *mongo.Client, collection *mongo.Collection, ctx context.Context, ID interface{}) (Log, error) {
	var result Log
	filter := bson.M{"_id": ID}

	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		return result, err
	}
	return result, nil
}

//GetLogsByIP returns a slice of logs with the defined IP.
//If the IP is not present in the database err won't be nil.
func GetLogsByIP(client *mongo.Client, collection *mongo.Collection, ctx context.Context, IP string) ([]Log, error) {
	var result []Log
	filter := bson.M{"ip": IP}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return result, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		return result, err
	}
	return result, nil
}
