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

//ApiServer > to be filled
func ApiServer() {

	// DB setup
	connString := os.Getenv("MONGO_CONN") // "mongodb://hostname:27017"
	dbName := os.Getenv("MONGO_DB")
	client := db.ConnectDB(connString)

	// Routes setup
	r := mux.NewRouter()

	//NotFound
	r.HandleFunc(NotFound, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 page not found")
	}).Methods("GET")

	//Health
	r.HandleFunc(Health, func(w http.ResponseWriter, r *http.Request) {
		HealthCheck(w, r)
	}).Methods("GET")

	//LogsDate
	r.HandleFunc(LogsDate, func(w http.ResponseWriter, r *http.Request) {
		LogsDateHandler(w, r, dbName, client)
	}).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    ":8094",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
