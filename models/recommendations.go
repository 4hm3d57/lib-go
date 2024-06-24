package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Recommendation struct {
	Title string
	Description string
}

func RecommDB() (*mongo.Client, *mongo.Collection, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	recommCollection := client.Database("lib").Collection("recommendation")

	return client, recommCollection, err
}



func InsertRecommendation(recommendation Recommendation) error {

	client, recommCollection, err := RecommDB()
	if err != nil {
		log.Printf("connection to db failed")
		return err
	}
	defer client.Disconnect(context.Background())

	_, err = recommCollection.InsertOne(context.Background(), recommendation)
	if err != nil {
		log.Printf("error adding message")
		return err
	}

	return err
}


func GetAllRecommendation() ([]Recommendation, error){

	client, recommCollection, err := RecommDB()
	if err != nil {
		log.Printf("connection to db failed")
		return nil, err
	}
	defer client.Disconnect(context.Background())


	var recomms []Recommendation

	cur, err := recommCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("failed to retrieve books")
		return nil, err
	}
	

	if err = cur.All(context.Background(), &recomms); err != nil {
		return nil, err
	}

	return recomms, err
	
}
