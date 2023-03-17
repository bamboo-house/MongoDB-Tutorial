package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:mongo@mongo:27017/sample"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("fetch start.")

	collection := client.Database("sample").Collection("sample_collection")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.Background()) {
		result := struct {
			ID         primitive.ObjectID
			sampleDate string
			username   string
		}{}

		err := cursor.Decode(&result)

		if err != nil {
			log.Fatal(err)
		}
		raw := cursor.Current
		log.Println(raw)
	}
	log.Println(collection)
	// Close connection
	client.Disconnect(ctx)
}
