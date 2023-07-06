package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Struct that represent the token of one user.
type UserToken struct {
	UserId primitive.ObjectID 	`json:"userId" bson:"_id,omitempty"`
	Token string 		`json:"token"`
}