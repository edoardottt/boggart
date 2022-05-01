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
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/edoardottt/boggart/db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//HealthCheck tells you if the API server is listening
func HealthCheck(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	//specify status code
	w.WriteHeader(http.StatusOK)
	//update response writer
	fmt.Fprintf(w, "OK")
}

//IPInfoResponse >
type IPInfoResponse struct {
	Logs         int
	LastActivity time.Time
	TopMethods   []string
	TopPaths     []string
	TopBodies    []string
}

//IPInfoHandler >
func IPInfoHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	database := db.GetDatabase(client, dbName)
	collection := db.GetLogs(database)

	vars := mux.Vars(req)
	ip := vars["ip"]
	topParam := req.URL.Query().Get("top")
	if topParam == "" {
		topParam = "10"
	}
	top, err := IsIntInTheRange(topParam, 4, 50)
	// 400 BAD REQUEST: top parameter not in the correct range
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "The parameter top accept an integer between 4 and 50.")
		return
	}

	topMethods, err := Top(w, req, dbName, client, "method", top, ip)
	if err != nil {
		return
	}
	topPaths, err := Top(w, req, dbName, client, "path", top, ip)
	if err != nil {
		return
	}
	topBodies, err := Top(w, req, dbName, client, "body", top, ip)
	if err != nil {
		return
	}

	filter := db.BuildFilter(map[string]interface{}{"ip": ip})
	findOptions := options.Find()
	// Sort by `timestamp` field descending
	findOptions.SetSort(bson.D{{Key: "timestamp", Value: -1}})
	logs, err := db.GetLogsWithFilter(client, collection, ctx, filter, findOptions)

	// 500 INTERNAL SERVER ERROR: generic error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error while retrieving data.")
		fmt.Println(err) //DEBUG: logging!
		return
	}

	i, err := strconv.ParseInt(fmt.Sprint(logs[0].Timestamp), 10, 64)
	// 500 INTERNAL SERVER ERROR: generic error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error while retrieving data.")
		fmt.Println(err) //DEBUG: logging!
		return
	}

	err = json.NewEncoder(w).Encode(IPInfoResponse{
		Logs:         len(logs),
		LastActivity: time.Unix(i, 0),
		TopMethods:   topMethods,
		TopPaths:     topPaths,
		TopBodies:    topBodies,
	})
	if err != nil {
		fmt.Println(err) //DEBUG: logging!
	}
}

//LogsHandler >
func LogsHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "TODO")
}

//LogsDetectHandler >
func LogsDetectHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "TODO")
}

//StatsHandler >
func StatsHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "TODO")
}

//StatsDBHandler >
func StatsDBHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "TODO")
}

//----------------------------------------
// -------------- HELPERS ----------------
//----------------------------------------

//IsIntInTheRange >
func IsIntInTheRange(input string, start int, end int) (int, error) {
	intVar, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	if intVar >= start && intVar <= end {
		return intVar, nil
	}
	return 0, errors.New("integer not in the range >= " + fmt.Sprint(start) + " && <= " + fmt.Sprint(end))
}

//Top >
func Top(w http.ResponseWriter, req *http.Request, dbName string,
	client *mongo.Client, what string, howMany int, IP string) ([]string, error) {

	if what != "method" && what != "path" && what != "body" {
		return nil, errors.New("possible values for top: method / path / body.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	database := db.GetDatabase(client, dbName)
	collection := db.GetLogs(database)

	filter := []bson.M{
		{"$match": bson.M{"ip": IP}},
		{"$sortByCount": "$" + what},
		{"$limit": howMany},
	}
	logs, err := db.GetAggregatedLogs(client, collection, ctx, filter)

	// 500 INTERNAL SERVER ERROR: generic error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error while retrieving data.")
		return nil, err
	}

	// 200: but
	if len(logs) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "no stats available for the specified IP.")
		return nil, errors.New("no stats available for the specified IP.")
	}

	var result []string
	for i := 0; i < howMany; i++ {
		if i < len(logs) {
			result = append(result, logs[i].ID)
		}
	}

	return result, nil
}
