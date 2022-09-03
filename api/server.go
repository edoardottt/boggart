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
	"log"
	"net/http"
	"os"
	"time"

	"github.com/edoardottt/boggart/db"
	"github.com/gorilla/mux"
)

const (
	WriteTimeoutDuration = 15
	ReadTimeoutDuration  = 15
)

// Server > to be filled.
func Server() {
	// DB setup.
	connString := os.Getenv("MONGO_CONN") // "mongodb://hostname:27017".
	dbName := os.Getenv("MONGO_DB")
	client, err := db.ConnectDB(connString)
	// ------- debug -------.
	if err != nil {
		fmt.Println("API: Connected to MongoDB!")
	}

	// Routes setup.
	router := mux.NewRouter()

	// NotFound.
	router.HandleFunc(NotFound, func(w http.ResponseWriter, r *http.Request) {
		NotFoundHandler(w, r)
	}).Methods("GET")

	// Health.
	router.HandleFunc(Health, func(w http.ResponseWriter, r *http.Request) {
		HealthHandler(w, r)
	}).Methods("GET")

	// IPInfo.
	router.HandleFunc(IPInfo, func(w http.ResponseWriter, r *http.Request) {
		IPInfoHandler(w, r, dbName, client)
	}).Methods("GET")

	// ApiLogs.
	router.HandleFunc(APILogs, func(w http.ResponseWriter, r *http.Request) {
		LogsHandler(w, r, dbName, client)
	}).Methods("GET")

	// ApiDetect.
	router.HandleFunc(APIDetect, func(w http.ResponseWriter, r *http.Request) {
		LogsDetectHandler(w, r, dbName, client)
	}).Methods("GET")

	// ApiStats.
	router.HandleFunc(APIStats, func(w http.ResponseWriter, r *http.Request) {
		StatsHandler(w, r, dbName, client)
	}).Methods("GET")

	// ApiStatsDB.
	router.HandleFunc(APIStatsDB, func(w http.ResponseWriter, r *http.Request) {
		StatsDBHandler(w, r, dbName, client)
	}).Methods("GET")

	srv := &http.Server{
		Handler: router,
		Addr:    ":8094",
		// Good practice: enforce timeouts for servers you create.
		WriteTimeout: WriteTimeoutDuration * time.Second,
		ReadTimeout:  ReadTimeoutDuration * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
