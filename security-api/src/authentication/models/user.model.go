package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status int8

const (
	Active Status = iota
	Blocked
	Inactive
)

func (s Status) String() string{
	return []string{"Active","Blocked","Inactive"}[s]
}

type User struct {
	ID primitive.ObjectID 	`json:"userId" bson:"_id,omitempty"`
	Username string 		`json:"username"`
	Password string 		`json:"password"`
	Email string			`json:"email"`
	Status string           `json:"status"`
	CreatedAt time.Time		`json:"created_at"`
	UpdatedAt time.Time		`json:"updated_at,omitempty"`
	DeletedAt time.Time     `json:"deleted_at" bson:"deleted_at,omitempty"`
	Role string				`json:"role"`
}

type Users []User