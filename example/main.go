package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	ai "github.com/feixiao/mongo-ai"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mmc:mmc007@172.20.99.13:27017"))
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("example-db").Collection("counters")

	ai := ai.Create(collection)

	for i := 0; i < 10; i++ {
		client.Database("example-db").Collection("users").InsertOne(ctx, bson.M{
			"_id":   ai.Next("sequenc"),
			"login": "test",
			"age":   32,
		})
	}
}
