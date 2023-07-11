package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID primitive.ObjectID 	`json:"id" bson:"_id,omitempty"`
	Name string `json:"name"`
	Description string `json:"description"`
	Stock int `json:"stock"`
	Price float32 `json:"price"`
	CreatedAt time.Time		`json:"created_at"`
	UpdatedAt time.Time		`json:"updated_at,omitempty"`
}