package database

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb://admin:password@localhost:27017/?maxPoolSize=20&w=majority"

func init() {
	fmt.Println("Initializing Database with Data from local JSON File")
	count := LoadCertifications()
	fmt.Println("Initialized DB with Size: ", count)
}


func Connect() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic(err)
		return nil
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Panic(err)
		return nil
	}
	fmt.Println("Successfully connected & pinged")
	return client
}

func Disconnect(client *mongo.Client) error {
	if client != nil {
		return client.Disconnect(context.TODO())
	}
	return errors.New("No client to disconnect.")
}
