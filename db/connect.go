package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ConnectDB creates and returns a client connected by a
//connection string to mongoDB.
//Also checks the connection if everything is ok.
func ConnectDB(connectionString string, dbName string, user string, pass string) *mongo.Client {
	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    dbName,
		Username:      user,
		Password:      pass,
	}
	clientOpts := options.Client().ApplyURI(connectionString).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOpts)
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
