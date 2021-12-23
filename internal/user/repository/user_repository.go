package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"save-a-buddy-api/db"
	"save-a-buddy-api/model"
)

type IUserRepository interface {
	FindUsersDb() (model.Users, error)
}

type UserRepository struct {
	mongodb *db.MongoDb
}

func New(mongodb *db.MongoDb) UserRepository {
	return UserRepository{mongodb: mongodb}
}

func (ur UserRepository) FindUsersDb() (model.Users, error) {
	collection := ur.mongodb.MongoUserCollection()
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
