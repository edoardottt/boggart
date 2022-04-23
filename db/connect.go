package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ConnectDB creates and returns a client connected by a
//connection string to mongoDB.
//Also checks the connection if everything is ok.
func ConnectDB(connectionString string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!") //!! DEBUG !!
	return client
}

//GetDatabase returns the pointer to the database (input).
func GetDatabase(client *mongo.Client, databaseName string) *mongo.Database {
	database := client.Database(databaseName)

	return database
}

//GetLogs returns the collection of logs
func GetLogs(database *mongo.Database) *mongo.Collection {
	return database.Collection("logs")
}
