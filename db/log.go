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

package db

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Log defines the structure of a log record
//in the database
type Log struct {
	ID        primitive.ObjectID  `json:"_id"`
	IP        string              `json:"ip"`
	Method    string              `json:"method"`
	Path      string              `json:"path"`
	Headers   map[string][]string `json:"headers"`
	Body      string              `json:"body"`
	Timestamp int64               `json:"timestamp"`
}

//IsEmpty checks if a Log is a new one (just created)
func (log Log) IsEmpty() bool {
	return reflect.DeepEqual(log, Log{})
}

//InsertLog inserts a log record into the logs collection
func InsertLog(client *mongo.Client, collection *mongo.Collection, ctx context.Context, record Log) interface{} {
	record.ID = primitive.NewObjectID()
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

//GetLogsByMethod returns a slice of logs with the defined HTTP Method.
//If the Method is not present in the database err won't be nil.
func GetLogsByMethod(client *mongo.Client, collection *mongo.Collection, ctx context.Context, method string) ([]Log, error) {
	var result []Log
	filter := bson.M{"method": method}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return result, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		return result, err
	}
	return result, nil
}

//GetLogsByPath returns a slice of logs with the defined Path.
//If the Path is not present in the database err won't be nil.
func GetLogsByPath(client *mongo.Client, collection *mongo.Collection, ctx context.Context, path string) ([]Log, error) {
	var result []Log
	filter := bson.M{"path": path}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return result, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		return result, err
	}
	return result, nil
}

//GetLogsByBody returns a slice of logs with the defined Body.
//If the Body is not present in the database err won't be nil.
func GetLogsByBody(client *mongo.Client, collection *mongo.Collection, ctx context.Context, body string) ([]Log, error) {
	var result []Log
	filter := bson.M{"body": body}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return result, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		return result, err
	}
	return result, nil
}

//GetLogsByDate returns a slice of logs within the defined date.
//If the Date is not present in the database err won't be nil.
func GetLogsByDate(client *mongo.Client, collection *mongo.Collection, ctx context.Context, date time.Time) ([]Log, error) {
	var result []Log
	nextDateInt := date.Add(time.Hour * 24).Unix()
	filter := bson.M{
		"$and": []bson.M{
			{"timestamp": bson.M{"$gte": date.Unix()}},
			{"timestamp": bson.M{"$lt": nextDateInt}},
		},
	}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return result, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		return result, err
	}
	return result, nil
}

//GetLogsByRange returns a slice of logs within the defined range (date to date).
//If the Range is not present in the database err won't be nil.
func GetLogsByRange(client *mongo.Client, collection *mongo.Collection, ctx context.Context, dateStart time.Time, dateEnd time.Time) ([]Log, error) {
	var result []Log
	filter := bson.M{
		"$and": []bson.M{
			{"timestamp": bson.M{"$gte": dateStart.Unix()}},
			{"timestamp": bson.M{"$lt": dateEnd.Unix()}},
		},
	}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return result, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		return result, err
	}
	return result, nil
}

//GetLatestNLogs returns a slice of the latest inserted N logs.
//If they are not present in the database err won't be nil.
func GetLatestNLogs(client *mongo.Client, collection *mongo.Collection, ctx context.Context, n int64) ([]Log, error) {
	var result []Log
	findOptions := options.Find()
	// Sort by `timestamp` field descending
	findOptions.SetSort(bson.D{{Key: "timestamp", Value: -1}})
	findOptions.Limit = &n
	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return result, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		return result, err
	}
	return result, nil
}
