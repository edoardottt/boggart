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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//BuildFilter returns a bson.M object representing the
//mongoDB filter in a query.
func BuildFilter(query map[string]interface{}) bson.M {
	result := make(bson.M, len(query))
	for k, v := range query {
		result[k] = v
	}
	return result
}

//AddCondition returns the query passed as input with
//the query passed as input.
func AddCondition(query bson.M, condition string, add interface{}) bson.M {
	query[condition] = add
	return query
}

//AddMultipleCondition returns the query passed as input with
//the multiple queries passed as input.
func AddMultipleCondition(query bson.M, condition string, add []bson.M) bson.M {
	query[condition] = add
	return query
}

//GetLogsWithFilter returns a slice of logs using the
//filter taken as input.
//If the result is empty err won't be nil.
func GetLogsWithFilter(client *mongo.Client, collection *mongo.Collection,
	ctx context.Context, filter bson.M, findOptions *options.FindOptions) ([]Log, error) {
	var result []Log
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return result, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		return result, err
	}
	return result, nil
}

//AggregatedResult >
type AggregatedResult struct {
	ID    string `bson:"_id"`
	Count int    `bson:"count"`
}

//GetAggregatedLogs returns a slice of aggregated logs
//using the filter taken as input.
//If the result is empty err won't be nil.
func GetAggregatedLogs(client *mongo.Client, collection *mongo.Collection,
	ctx context.Context, filter []bson.M) ([]AggregatedResult, error) {
	var result []AggregatedResult
	cursor, err := collection.Aggregate(ctx, filter)
	if err != nil {
		return result, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		return result, err
	}
	return result, nil
}
