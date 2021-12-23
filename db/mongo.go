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

var mongoClient *mongo.Client

const (
	collectionName = "user"
	dataBaseName   = "myFirstDatabase"
)

type MongoDb struct {
	mongoURI string
}

func NewConnection(mongoUri string) *MongoDb {
	return &MongoDb{mongoURI: mongoUri}
}

// Connect to Mongo DB
func (m MongoDb) Connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI(m.mongoURI)
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

	mongoClient = client

	log.Printf("Connected to DB")
	return client
}

func (m MongoDb) MongoUserCollection() *mongo.Collection {

	userCollection := mongoClient.Database(dataBaseName).Collection(collectionName)
	return userCollection
}

// ValidateConnection Validate Connection by PING
func ValidateConnection(client *mongo.Client) bool {
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
