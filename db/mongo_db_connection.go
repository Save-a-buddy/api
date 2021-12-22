package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type IMongoDb interface {
	NewConnection(mongoUri string) *MongoDb
	ConnectToDB() *mongo.Client
	ValidateConnection(client *mongo.Client) bool
}

type MongoDb struct {
	MongoURI string
}

func NewConnection(mongoUri string) *MongoDb {
	return &MongoDb{MongoURI: mongoUri}
}

// ConnectToDB Connect to Mongo DB
func (m MongoDb) ConnectToDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI(m.MongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Printf("Connected to DB")
	return client
}

// ValidateConnection Validate Connection by PING
func ValidateConnection(client *mongo.Client) bool {
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
