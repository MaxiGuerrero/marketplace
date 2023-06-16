package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTBuilder interface {
	BuildToken(payload *Payload) string
}

type Payload struct{
	UserId primitive.ObjectID `json:"userId" bson:"_id,omitempty"`
	Username string			`json:"username"`
	CreatedAt time.Time		`json:"created_at"`
	UpdatedAt time.Time		`json:"updated_at,omitempty"`
	DeletedAt time.Time     `json:"deleted_at" bson:"deleted_at,omitempty"`
}