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
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/edoardottt/boggart/db"
	net "github.com/edoardottt/boggart/internal/net"
	"github.com/edoardottt/boggart/internal/slice"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NotFoundHandler tells you if the API server is listening.
func NotFoundHandler(w http.ResponseWriter, req *http.Request) {
	// set content-type.
	w.Header().Add("Content-Type", "application/json")
	// specify status code.
	w.WriteHeader(http.StatusNotFound)
	// update response writer.
	fmt.Fprintf(w, "404 page not found")
}

// HealthHandler tells you if the API server is listening.
func HealthHandler(w http.ResponseWriter, req *http.Request) {
	// set content-type.
	w.Header().Add("Content-Type", "application/json")
	// specify status code.
	w.WriteHeader(http.StatusOK)
	// update response writer.
	fmt.Fprintf(w, "OK")
}

// IPInfoResponse.
type IPInfoResponse struct {
	Logs         int
	LastActivity time.Time
	TopMethods   []string
	TopPaths     []string
	TopBodies    []string
}

// IPInfoHandler.
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
	// 400 BAD REQUEST: top parameter not in the correct range.
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
	// Sort by `timestamp` field descending.
	findOptions.SetSort(bson.D{{Key: "timestamp", Value: -1}})
	logs, err := db.GetLogsWithFilter(client, collection, ctx, filter, findOptions)

	// 500 INTERNAL SERVER ERROR: generic error.
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

// LogsHandler.
func LogsHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "TODO")
}

// LogsDetectHandler.
func LogsDetectHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "TODO")
}

// StatsHandler.
func StatsHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "TODO")
}

// StatsDBHandler.
func StatsDBHandler(w http.ResponseWriter, req *http.Request, dbName string, client *mongo.Client) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "TODO")
}

// ---------------------------------------
// -------------- HELPERS ----------------
// ---------------------------------------

// IsIntInTheRange.
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

// Top.
func Top(w http.ResponseWriter, req *http.Request, dbName string,
	client *mongo.Client, what string, howMany int, ip string) ([]string, error) {
	if what != "method" && what != "path" && what != "body" {
		return nil, ErrPossibleTopValue
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	database := db.GetDatabase(client, dbName)
	collection := db.GetLogs(database)

	filter := []bson.M{
		{"$match": bson.M{"ip": ip}},
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

		return nil, ErrNoStatsIP
	}

	var result []string

	for i := 0; i < howMany; i++ {
		if i < len(logs) {
			result = append(result, logs[i].ID)
		}
	}

	return result, nil
}

// GetAPILogsQuery >
// - id
// - ip
// - method
// - header
// - path
// - date (YYYY-MM-DD)
// - lt (less than YYYY-MM-DD-HH-MM-SS)
// - gt (greater than YYYY-MM-DD-HH-MM-SS).
func GetAPILogsQuery(req *http.Request) (bson.M, error) {
	id := req.URL.Query().Get("id")
	ip := req.URL.Query().Get("ip")
	method := req.URL.Query().Get("method")
	header := req.URL.Query().Get("header")
	path, err := url.QueryUnescape(req.URL.Query().Get("path"))

	if err != nil {
		return bson.M{}, err
	}

	date := req.URL.Query().Get("date")
	lt := req.URL.Query().Get("lt")
	gt := req.URL.Query().Get("gt")
	err = CheckAPILogsParams(id, ip, method, header, path, date, lt, gt)

	if err != nil {
		return bson.M{}, err
	}

	// build query.
	return bson.M{}, nil
}

// CheckAPILogsParams.
func CheckAPILogsParams(id, ip, method, header, path, date, lt, gt string) error {
	// if id is present, the others are blank.
	if id != "" {
		if ip != "" || method != "" || header != "" || path != "" || date != "" || lt != "" || gt != "" {
			return ErrIDDefined
		}
	}
	// if date is present, lt and gt are blank.
	if date != "" {
		if lt != "" || gt != "" {
			return ErrDateDefined
		}

		_, err := TranslateTime(date)
		if err != nil {
			return err
		}
	}

	if lt != "" {
		_, err := TranslateTime(lt)

		if err != nil {
			return err
		}
	}

	if gt != "" {
		_, err := TranslateTime(gt)

		if err != nil {
			return err
		}
	}

	if lt != "" && gt != "" {
		ltT, _ := TranslateTime(lt)
		gtT, _ := TranslateTime(gt)

		if ltT.Unix() < gtT.Unix() {
			return ErrLtBeforeGt
		}
	}

	if method != "" {
		method = strings.ToUpper(method)
		if !slice.StringInSlice(method,
			[]string{
				"GET",
				"HEAD",
				"POST",
				"PUT",
				"PATCH",
				"DELETE",
				"CONNECT",
				"OPTIONS",
				"TRACE"}) {
			return ErrHTTPMethodUnknown
		}
	}

	if ip != "" {
		if !net.ValidIPAddress(ip) {
			return ErrInvalidIP
		}
	}

	return nil
}

// BuildAPILogsQuery.
func BuildAPILogsQuery(id, ip, method, header, path, date, lt, gt string) bson.M {
	var filter bson.M
	if id != "" {
		filter = db.BuildFilter(map[string]interface{}{"_id": id})

		return filter
	}

	if ip != "" {
		filter = db.BuildFilter(map[string]interface{}{"ip": ip})
	}

	if method != "" {
		if len(filter) != 0 {
			filter = db.AddCondition(filter, "method", method)
		} else {
			filter = db.BuildFilter(map[string]interface{}{"method": method})
		}
	}

	// DEBUG: map[string][]string.
	if header != "" {
		if len(filter) != 0 {
			filter = db.AddCondition(filter, "header", header)
		} else {
			filter = db.BuildFilter(map[string]interface{}{"header": header})
		}
	}

	if path != "" {
		if len(filter) != 0 {
			filter = db.AddCondition(filter, "path", path)
		} else {
			filter = db.BuildFilter(map[string]interface{}{"path": path})
		}
	}

	if path != "" {
		if len(filter) != 0 {
			filter = db.AddCondition(filter, "path", path)
		} else {
			filter = db.BuildFilter(map[string]interface{}{"path": path})
		}
	}

	if date != "" {
		dateT, _ := TranslateTime(date)

		if len(filter) == 0 {
			filter = db.BuildFilter(map[string]interface{}{})
		}

		filter = db.AddMultipleCondition(filter, "$and", []bson.M{
			{"timestamp": bson.M{"$gte": dateT.Unix()}},
			{"timestamp": bson.M{"$lt": dateT.Add(time.Hour * 24).Unix()}},
		})
	}

	filter = AddTimestampToQuery(lt, gt, filter)

	return filter
}

// GetAPIDetectQuery >
// - regex (Go format)
// - attack (use a list of well known regex)
// - target (where to apply the regex)
// - ip
// - method
// - header
// - path
// - date
// - lt (less than YYYY-MM-DD-HH-MM-SS)
// - gt (greater than YYYY-MM-DD-HH-MM-SS).
func GetAPIDetectQuery(req *http.Request) (bson.M, error) {
	regex := req.URL.Query().Get("regex")
	attack := req.URL.Query().Get("attack")
	target := req.URL.Query().Get("target")
	ip := req.URL.Query().Get("ip")
	method := req.URL.Query().Get("method")
	header := req.URL.Query().Get("header")
	path, err := url.QueryUnescape(req.URL.Query().Get("path"))

	if err != nil {
		return bson.M{}, err
	}

	date := req.URL.Query().Get("date")
	lt := req.URL.Query().Get("lt")
	gt := req.URL.Query().Get("gt")
	err = CheckAPIDetectParams(regex, attack, target, ip, method, header, path, date, lt, gt)

	if err != nil {
		return bson.M{}, err
	}

	// build query.
	return bson.M{}, nil
}

// CheckAPIDetectParams.
func CheckAPIDetectParams(regex, attack, target, ip, method, header, path, date, lt, gt string) error {
	// if date is present, lt and gt are blank.
	if date != "" {
		if lt != "" || gt != "" {
			return ErrDateDefined
		}

		_, err := TranslateTime(date)
		if err != nil {
			return err
		}
	}

	if lt != "" {
		_, err := TranslateTime(lt)
		if err != nil {
			return err
		}
	}

	if gt != "" {
		_, err := TranslateTime(gt)
		if err != nil {
			return err
		}
	}

	if lt != "" && gt != "" {
		ltT, _ := TranslateTime(lt)
		gtT, _ := TranslateTime(gt)

		if ltT.Unix() < gtT.Unix() {
			return ErrLtBeforeGt
		}
	}

	if method != "" {
		method = strings.ToUpper(method)
		if !slice.StringInSlice(method,
			[]string{
				"GET",
				"HEAD",
				"POST",
				"PUT",
				"PATCH",
				"DELETE",
				"CONNECT",
				"OPTIONS",
				"TRACE"}) {
			return ErrHTTPMethodUnknown
		}
	}

	if ip != "" {
		if !net.ValidIPAddress(ip) {
			return ErrInvalidIP
		}
	}

	return nil
}

// BuildAPIDetectQuery.
func BuildAPIDetectQuery(regex, attack, target, ip, method, header, path, date, lt, gt string) bson.M {
	var filter bson.M
	/*
		DEBUG
		Implement regex, attack, target.
	*/
	if ip != "" {
		filter = db.BuildFilter(map[string]interface{}{"ip": ip})
	}

	if method != "" {
		if len(filter) != 0 {
			filter = db.AddCondition(filter, "method", method)
		} else {
			filter = db.BuildFilter(map[string]interface{}{"method": method})
		}
	}

	// DEBUG: map[string][]string.
	if header != "" {
		if len(filter) != 0 {
			filter = db.AddCondition(filter, "header", header)
		} else {
			filter = db.BuildFilter(map[string]interface{}{"header": header})
		}
	}

	if path != "" {
		if len(filter) != 0 {
			filter = db.AddCondition(filter, "path", path)
		} else {
			filter = db.BuildFilter(map[string]interface{}{"path": path})
		}
	}

	if path != "" {
		if len(filter) != 0 {
			filter = db.AddCondition(filter, "path", path)
		} else {
			filter = db.BuildFilter(map[string]interface{}{"path": path})
		}
	}

	if date != "" {
		dateT, _ := TranslateTime(date)

		if len(filter) == 0 {
			filter = db.BuildFilter(map[string]interface{}{})
		}

		filter = db.AddMultipleCondition(filter, "$and", []bson.M{
			{"timestamp": bson.M{"$gte": dateT.Unix()}},
			{"timestamp": bson.M{"$lt": dateT.Add(time.Hour * 24).Unix()}},
		})
	}

	filter = AddTimestampToQuery(lt, gt, filter)

	return filter
}

// TranslateTime.
func TranslateTime(input string) (time.Time, error) {
	t, err := time.Parse("2006-01-02T15:04:05-0700", input)
	if err != nil {
		return time.Time{}, ErrDatetimeFormat
	}

	return t, nil
}

// AddTimestampToQuery.
func AddTimestampToQuery(lt, gt string, filter bson.M) bson.M {
	if lt != "" && gt != "" {
		ltT, _ := TranslateTime(lt)
		gtT, _ := TranslateTime(gt)

		if len(filter) == 0 {
			filter = db.BuildFilter(map[string]interface{}{})
		}

		filter = db.AddMultipleCondition(filter, "$and", []bson.M{
			{"timestamp": bson.M{"$gte": gtT.Unix()}},
			{"timestamp": bson.M{"$lt": ltT.Add(time.Hour * 24).Unix()}},
		})
	} else if lt != "" {
		ltT, _ := TranslateTime(lt)
		if len(filter) == 0 {
			filter = db.BuildFilter(map[string]interface{}{})
		}
		filter = db.AddMultipleCondition(filter, "timestamp", []bson.M{
			{"$lt": ltT.Unix()},
		})
	} else if gt != "" {
		gtT, _ := TranslateTime(gt)
		if len(filter) == 0 {
			filter = db.BuildFilter(map[string]interface{}{})
		}
		filter = db.AddMultipleCondition(filter, "timestamp", []bson.M{
			{"$gte": gtT.Unix()},
		})
	}

	return filter
}
