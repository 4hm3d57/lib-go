package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Book struct {
	Title string `bson:"title"`
	Author string `bson:"author"`
	Publisher string `bson:"publisher"`
	Year string `bson:"year"`
	Copies string `bson:"copies"`
}

func BookDB() (*mongo.Client, *mongo.Collection, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	bookCollection := client.Database("lib").Collection("books")

	return client, bookCollection, err
}



func InsertBook(book Book) error {

	client, bookCollection, err := BookDB()
	if err != nil {
		log.Printf("connection to db failed")
		return err
	}
	defer client.Disconnect(context.Background())

	_, err = bookCollection.InsertOne(context.Background(), book)
	if err != nil {
		log.Printf("error adding book")
		return err
	}

	return err
}


func GetAllBooks() ([]Book, error){

	client, bookCollection, err := BookDB()
	if err != nil {
		log.Printf("connection to db failed")
		return nil, err
	}
	defer client.Disconnect(context.Background())


	var books []Book

	cur, err := bookCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("failed to retrieve books")
		return nil, err
	}


	if err = cur.All(context.Background(), &books); err != nil {
		return nil, err
	}

	return books, err
	
}
