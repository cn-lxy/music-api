package db

import (
	"context"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestMongoInsert(t *testing.T) {
	collection := Client.Database("test").Collection("demo")

	doc := bson.D{bson.E{
		Key:   "name",
		Value: "dave",
	}}
	res, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(res.InsertedID)
}
