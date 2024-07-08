package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Message struct {
	Name     string `bson:"name"`
	Roll_no  string `bson:"roll_no"`
	Messages string `bson:"messages"`
}

func MessageDB() (*mongo.Client, *mongo.Collection, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	messageCollection := client.Database("lib").Collection("message")

	return client, messageCollection, err
}

func InsertMessage(message Message) error {

	client, messageCollection, err := MessageDB()
	if err != nil {
		log.Printf("connection to db failed")
		return err
	}
	defer client.Disconnect(context.Background())

	_, err = messageCollection.InsertOne(context.Background(), message)
	if err != nil {
		log.Printf("error adding message")
		return err
	}

	return err
}

func GetAllMessages() ([]Message, error) {

	client, messageCollection, err := MessageDB()
	if err != nil {
		log.Printf("connection to db failed")
		return nil, err
	}
	defer client.Disconnect(context.Background())

	var messages []Message

	cur, err := messageCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("failed to retrieve books")
		return nil, err
	}

	if err = cur.All(context.Background(), &messages); err != nil {
		return nil, err
	}

	return messages, err

}
