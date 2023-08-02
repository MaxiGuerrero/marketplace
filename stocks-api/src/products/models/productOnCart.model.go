package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductOnCart struct {
	ProductId primitive.ObjectID `json:"productId" bson:"_id,omitempty"`
	Amount int `json:"amount"`
}