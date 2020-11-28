// Package data creates an instance of UsersDB
package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURI = "mongodb://localhost:27017"

var (
	Client  *mongo.Client
	UsersDB *mongo.Collection
)

func init() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}
	UsersDB = Client.Database("bookstore").Collection("users")
	log.Println("Successfully connected to MongoDB")
}
