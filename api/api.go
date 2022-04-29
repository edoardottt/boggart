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

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/edoardottt/boggart/db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//HealthCheck tells you if the API server is listening
func HealthCheck(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	//specify status code
	w.WriteHeader(http.StatusOK)

	//update response writer
	fmt.Fprintf(w, "OK")
}

//LogsDateHandler >
func LogsDateHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	database := db.GetDatabase(client, dbName)
	collection := db.GetLogs(database)

	vars := mux.Vars(req)
	dateT, err := time.Parse("2006-01-02", vars["date"])

	// 400 BAD REQUEST: Time format != YYYY-MM-DD
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Cannot convert the date. Time format is YYYY-MM-DD.")
		return
	}

	filter := db.BuildFilter(map[string]interface{}{})
	db.AddMultipleCondition(filter, "$and", []bson.M{
		{"timestamp": bson.M{"$gte": dateT.Unix()}},
		{"timestamp": bson.M{"$lt": dateT.Add(time.Hour * 24).Unix()}},
	})
	logs, err := db.GetLogsWithFilter(client, collection, ctx, filter)

	// 500 INTERNAL SERVER ERROR: generic error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error while retrieving data.")
		fmt.Println(err)
		return
	}

	if len(logs) == 0 {
		fmt.Fprintf(w, "{}")
	} else {
		err = json.NewEncoder(w).Encode(logs)
		if err != nil {
			fmt.Println(err) //DEBUG: logging!
		}
	}
}

//LogsRangeHandler >
func LogsRangeHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "TODO")
}

//IPInfoHandler >
func IPInfoHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "TODO")
}

//LogsIPHandler >
func LogsIPHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "TODO")
}

//LogsIPDateHandler >
func LogsIPDateHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	database := db.GetDatabase(client, dbName)
	collection := db.GetLogs(database)

	vars := mux.Vars(req)
	ip := vars["ip"]
	dateT, err := time.Parse("2006-01-02", vars["date"])

	// 400 BAD REQUEST: Time format != YYYY-MM-DD
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Cannot convert the date. Time format is YYYY-MM-DD.")
		return
	}

	filter := db.BuildFilter(map[string]interface{}{})
	filter = db.AddMultipleCondition(filter, "$and", []bson.M{
		{"timestamp": bson.M{"$gte": dateT.Unix()}},
		{"timestamp": bson.M{"$lt": dateT.Add(time.Hour * 24).Unix()}},
	})
	filter = db.AddCondition(filter, "ip", ip)
	logs, err := db.GetLogsWithFilter(client, collection, ctx, filter)

	// 500 INTERNAL SERVER ERROR: generic error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error while retrieving data.")
		fmt.Println(err)
		return
	}

	if len(logs) == 0 {
		fmt.Fprintf(w, "{}")
	} else {
		err = json.NewEncoder(w).Encode(logs)
		if err != nil {
			fmt.Println(err) //DEBUG: logging!
		}
	}
}

//LogsIPRangeHandler >
func LogsIPRangeHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "TODO")
}

//LogsPathHandler >
func LogsPathHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "TODO")
}

//LogsPathDateHandler >
func LogsPathDateHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "TODO")
}

//LogsPathRangeHandler >
func LogsPathRangeHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "TODO")
}

//LogsMethodHandler >
func LogsMethodHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "TODO")
}

//LogsMethodDateHandler >
func LogsMethodDateHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	database := db.GetDatabase(client, dbName)
	collection := db.GetLogs(database)

	vars := mux.Vars(req)
	date := vars["date"]
	method := vars["method"]
	dateT, err := time.Parse("2006-01-02", date)

	// 400 BAD REQUEST: Time format != YYYY-MM-DD
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Cannot convert the date. Time format is YYYY-MM-DD.")
		return
	}
	logs, err := db.GetLogsByMethodAndDate(client, collection, ctx, dateT, method)

	// 500 INTERNAL SERVER ERROR: generic error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error while retrieving data.")
		fmt.Println(err)
		return
	}

	if len(logs) == 0 {
		fmt.Fprintf(w, "{}")
	} else {
		err = json.NewEncoder(w).Encode(logs)
		if err != nil {
			fmt.Println(err) //DEBUG: logging!
		}
	}
}

//LogsMethodRangeHandler >
func LogsMethodRangeHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "TODO")
}
