package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model
type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `bson:"name,omitempty"`
	LastName string             `bson:"lastName,omitempty"`
	Age      string             `bson:"age,omitempty"`
	Email    string             `bson:"email" json:"email"`
	QRCode   string             `bson:"QRCode,omitempty"`
	Password string             `bson:"password" json:"password"`
	//CreateAt time.Time          `json:"create_at"`
	//UpdateAt time.Time          `json:"update_at,omitempty"`
}

// Users list
type Users []*User
