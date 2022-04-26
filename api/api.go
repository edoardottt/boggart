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
	"fmt"
	"net/http"

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

	fmt.Fprint(w, "TODO")
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

	fmt.Fprint(w, "TODO")
}

//LogsIPRangeHandler >
func LogsIPRangeHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "TODO")
}
