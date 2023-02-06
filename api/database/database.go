package database

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	Client *mongo.Client
}

const uri = "mongodb://admin:password@localhost:27017/?maxPoolSize=20&w=majority"

func init() {
	fmt.Println("Initializing Database with Data from local JSON File")
	count := LoadCertifications()
	fmt.Println("Initialized DB with Size: ", count)
}


func Connect() (*DB, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected & pinged")
	return &DB{
		Client:client,
	}, nil
}

func (db *DB) Disconnect() error {
	client := db.Client
	if client != nil {
		return client.Disconnect(context.TODO())
	}
	return errors.New("no client to disconnect")
}

func (db *DB) GetDatabase() *mongo.Database{
	return db.Client.Database("db")
}

func (db *DB) GetClient() *mongo.Client {
	return db.Client
}