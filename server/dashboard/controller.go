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

package dashboard

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/edoardottt/boggart/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func dashboardIndexHandler(w http.ResponseWriter, tmpl *template.Template) {
	buf := &bytes.Buffer{}

	err := tmpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func dashboardLatestHandler(w http.ResponseWriter, client *mongo.Client, dbName string,
	tmpl *template.Template) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextBackgroundDuration*time.Second)
	defer cancel()

	database := db.GetDatabase(client, dbName)
	collection := db.GetLogs(database)

	logs, err := db.GetLatestNLogs(ctx, client, collection, LatestNLogs)
	if err != nil {
		fmt.Println(err)

		return
	}

	buf := &bytes.Buffer{}

	err = tmpl.Execute(buf, logs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func dashboardIDHandler(w http.ResponseWriter, client *mongo.Client, dbName string,
	tmpl *template.Template, id string) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextBackgroundDuration*time.Second)
	defer cancel()

	database := db.GetDatabase(client, dbName)
	collection := db.GetLogs(database)
	logID, err := db.GetLogByID(ctx, client, collection, id)

	if err != nil {
		fmt.Println(err)
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, logID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func dashboardOverviewHandler(w http.ResponseWriter, client *mongo.Client, dbName string,
	tmpl *template.Template) {
	type Result struct {
		Requests int64
	}

	ctx, cancel := context.WithTimeout(context.Background(), ContextBackgroundDuration*time.Second)
	defer cancel()

	database := db.GetDatabase(client, dbName)
	collection := db.GetLogs(database)

	logs, err := db.GetNumberOfLogs(ctx, client, collection)
	if err != nil {
		fmt.Println(err)

		return
	}

	result := Result{Requests: logs}
	buf := &bytes.Buffer{}

	err = tmpl.Execute(buf, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func dashboardQueryHandler(w http.ResponseWriter, tmpl *template.Template) {
	buf := &bytes.Buffer{}

	err := tmpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func dashboardDetectionHandler(w http.ResponseWriter, tmpl *template.Template) {
	buf := &bytes.Buffer{}

	err := tmpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func dashboardStatusHandler(w http.ResponseWriter, tmpl *template.Template) {
	buf := &bytes.Buffer{}

	err := tmpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
