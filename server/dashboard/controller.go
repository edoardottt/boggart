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
	"html"
	"net/http"
	"net/url"
	"strconv"
	"text/template"
	"time"

	"github.com/edoardottt/boggart/db"
	"github.com/edoardottt/boggart/internal/slice"
	timeUtils "github.com/edoardottt/boggart/internal/time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DayTime = time.Hour * 24
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

type Result struct {
	Logs  []db.Log
	Error string
}

func dashboardResultHandler(r *http.Request, w http.ResponseWriter, client *mongo.Client, dbName string,
	tmpl *template.Template) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextBackgroundDuration*time.Second)
	defer cancel()

	database := db.GetDatabase(client, dbName)
	collection := db.GetLogs(database)
	query := r.URL.Query()
	result := Result{}

	if errorOk := returnValuesOk(query); errorOk != "" {
		result.Error = errorOk
	} else {
		filter, findOptions := buildReturnQuery(query)
		// Sort by `timestamp` field descending.
		findOptions.SetSort(bson.D{{Key: "timestamp", Value: -1}})
		logs, err := db.GetLogsWithFilter(ctx, client, collection, filter, findOptions)

		if err != nil {
			fmt.Println(err)

			return
		}

		result.Logs = logs
		if len(result.Logs) == 0 {
			result.Error = "There is no data for the specified input."
		}
	}

	buf := &bytes.Buffer{}

	err := tmpl.Execute(buf, result)
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

func buildReturnQuery(query url.Values) (bson.M, *options.FindOptions) {
	var (
		ip      = query["ip"]
		method  = query["method"]
		path    = query["path"]
		headers = query["headers"]
		body    = query["body"]
		daytime = query["day"]
		//lt        = query["lt"]
		//gt        = query["gt"]
	)

	filter := db.BuildFilter(map[string]interface{}{})

	if ip[0] != "" {
		filter = db.AddCondition(filter, "ip", ip[0])
	}

	if method[0] != "" {
		filter = db.AddCondition(filter, "method", method[0])
	}

	if path[0] != "" {
		filter = db.AddCondition(filter, "path", path[0])
	}

	if headers[0] != "" {
		filter = db.AddCondition(filter, "headers", headers[0])
	}

	if body[0] != "" {
		filter = db.AddCondition(filter, "body", body[0])
	}

	if len(daytime) == 1 && daytime[0] != "" {
		day, _ := timeUtils.GetDay(daytime[0])
		// Greater than date start and less or equal date end.
		filter = db.AddMultipleCondition(filter, "$and", []bson.M{
			{"timestamp": bson.M{"$gte": day.Unix()}},
			{"timestamp": bson.M{"$lt": day.Add(DayTime).Unix()}},
		})
	}

	findOptions := options.Find()

	if query["limit"][0] != "" {
		limit, _ := strconv.Atoi(query["limit"][0])
		findOptions = options.Find().SetLimit(int64(limit))
	}

	return filter, findOptions
}

// returnValuesOK checks if the query provided has well-formatted values.
// Empty string okay, otherwise it contains the string to be shown.
func returnValuesOk(query url.Values) string {
	inputs := []string{"ip", "method", "path", "headers", "body", "day", "limit", "lt", "gt"}
	for _, input := range inputs {
		if len(query[input]) > 1 {
			return html.EscapeString(input) + " value cannot be more than one."
		}
	}

	if query["method"][0] != "" && !slice.StringInSlice(query["method"][0], []string{
		"GET", "POST", "PUT", "DELETE", "TRACE", "HEAD", "OPTIONS", "CONNECT", "PATCH"}) {
		return html.EscapeString(query["method"][0]) + " isn't a HTTP method."
	}

	if query["limit"][0] != "" {
		limit, err := strconv.Atoi(query["limit"][0])
		if err != nil || limit < 1 {
			return html.EscapeString(query["limit"][0]) + " isn't a positive integer value."
		}
	}

	if len(query["day"]) > 0 && query["day"][0] != "" {
		if _, err := timeUtils.GetDay(query["day"][0]); err != nil {
			return html.EscapeString(err.Error())
		}
	}

	if len(query["lt"]) > 0 && query["lt"][0] != "" {
		if _, err := timeUtils.TranslateTime(query["lt"][0]); err != nil {
			return html.EscapeString(err.Error())
		}
	}

	if len(query["gt"]) > 0 && query["gt"][0] != "" {
		if _, err := timeUtils.TranslateTime(query["gt"][0]); err != nil {
			return html.EscapeString(err.Error())
		}
	}

	if query["day"][0] != "" {
		if query["lt"][0] != "" || query["gt"][0] != "" {
			return "You cannot specify both day and lt/gt."
		}
	}

	if query["lt"][0] != "" && query["gt"][0] != "" {
		ltTime, _ := timeUtils.TranslateTime(query["lt"][0])
		gtTime, _ := timeUtils.TranslateTime(query["gt"][0])
		if ltTime.Unix() >= gtTime.Unix() {
			return "lt cannot be greater or equal than gt."
		}
	}

	return ""
}
