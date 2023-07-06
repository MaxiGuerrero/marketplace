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

// Method to get the string of the status type.
func (s Status) String() string{
	return []string{"Active","Blocked","Inactive"}[s]
}

type Role int8

const (
	ADMIN Role = iota
	USER
)

// Method to get the string of the role type.
func (s Role) String() string{
	return []string{"ADMIN","USER"}[s]
}

type User struct {
	ID primitive.ObjectID 	`bson:"_id,omitempty"`
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