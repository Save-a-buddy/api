package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"save-a-buddy-api/model"
)

type UserService struct {
	mongoClient *mongo.Client
}

func New(mongoClient *mongo.Client) UserService {
	return UserService{mongoClient: mongoClient}
}

func (us UserService) FindUsers() (model.Users, error) {
	collection := us.mongoClient.Database("myFirstDatabase").Collection("user")
	var users model.Users
	filter := bson.D{}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var user model.User
		err = cur.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}
