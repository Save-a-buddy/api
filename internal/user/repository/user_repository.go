package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"save-a-buddy-api/model"
)

type IUserRepository interface {
	FindUsersDb() (model.Users, error)
}

type UserRepository struct {
	mongoClient *mongo.Client
}

func New(mongoClient *mongo.Client) UserRepository {
	return UserRepository{mongoClient: mongoClient}
}

func (ur UserRepository) FindUsersDb() (model.Users, error) {
	collection := ur.mongoClient.Database("myFirstDatabase").Collection("user")
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
