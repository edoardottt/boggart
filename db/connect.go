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
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ContextBackgroundDuration = 5
)

// ConnectDB creates and returns a client connected by a
// connection string to mongoDB.
// Also checks the connection if everything is ok.
func ConnectDB(connectionString string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), ContextBackgroundDuration*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
		return client, nil
	}
	// Check the connection.
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
		return client, nil
	}

	return client, nil
}

// GetDatabase returns the pointer to the database (input).
func GetDatabase(client *mongo.Client, databaseName string) *mongo.Database {
	database := client.Database(databaseName)

	return database
}

// GetLogs returns the collection of logs.
func GetLogs(database *mongo.Database) *mongo.Collection {
	return database.Collection("logs")
}
