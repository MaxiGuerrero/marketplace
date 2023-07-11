package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RegisterProductRequest struct {
	Name string `json:"name" validate:"required,min=4,max=32"`
	Description string `json:"description" validate:"required,min=4,max=32"`
	Price float32 `json:"price" validate:"required,numeric,gt=0"`
	Stock int `json:"stock" validate:"required,numeric,gt=0"`
}

type UpdateProductRequest struct {
	Id primitive.ObjectID `json:"productId" validate:"required"`
	Name string `json:"name" validate:"required,min=4,max=32"`
	Description string `json:"description" validate:"required,min=4,max=32"`
	Price float32 `json:"price" validate:"required,numeric,gt=0"`
}

type UpdateStockRequest struct {
	Id primitive.ObjectID `json:"productId" validate:"required"`
	Stock int `json:"stock" validate:"required,numeric,gt=0"`
}