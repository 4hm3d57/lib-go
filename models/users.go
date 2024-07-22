package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Roll_no  string             `bson:"roll_no"`
	Password string             `bson:"password"`
	Acc_type string             `bson:"acc_type"`
}

func UserDB() (*mongo.Client, *mongo.Collection, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	userCollection := client.Database("lib").Collection("users")

	return client, userCollection, err

}

func InsertUser(user User) error {

	client, userCollection, err := UserDB()
	if err != nil {
		log.Printf("error connecting to the database", err)
		return err
	}
	defer client.Disconnect(context.Background())

	_, err = userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Printf("error inserting user", err)
		return err
	}

	return err

}

func GetUser(name, roll_no, password string) (*User, error) {

	client, userCollection, err := UserDB()
	if err != nil {
		log.Printf("error connecting to the database:", err)
		return nil, err
	}
	defer client.Disconnect(context.Background())

	var user User
	err = userCollection.FindOne(context.Background(), bson.M{"name": name, "roll_no": roll_no, "password": password}).Decode(&user)
	if err != nil {
		log.Printf("error finding user", err)
		return nil, err
	}

	return &user, nil
}

func GetUserID(userID primitive.ObjectID) (*User, error) {

	client, userCollection, err := UserDB()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	var user User
	err = userCollection.FindOne(context.Background(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func GetAllUsers() ([]User, error) {

	// connect to database
	client, userCollection, err := UserDB()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	// create a slice that has all users as its elements
	var users []User

	// collect the users from the database and append them in the slice
	cur, err := userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cur.All(context.Background(), &users); err != nil {
		return nil, err
	}

	return users, err

}
