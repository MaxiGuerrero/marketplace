package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Struct that represent a request to create a product.
type RegisterProductRequest struct {
	Name string `json:"name" validate:"required,min=4,max=32"`
	Description string `json:"description" validate:"required,min=4,max=32"`
	Price float32 `json:"price" validate:"required,numeric,gt=0"`
	Stock int `json:"stock" validate:"required,numeric,gt=0"`
}

// Struct that represent a request to update a product.
type UpdateProductRequest struct {
	Id primitive.ObjectID `json:"productId" validate:"required"`
	Name string `json:"name" validate:"required,min=4,max=32"`
	Description string `json:"description" validate:"required,min=4,max=32"`
	Price float32 `json:"price" validate:"required,numeric,gt=0"`
}

// Struct that represent a request to update the stock of a product.
type UpdateStockRequest struct {
	Id primitive.ObjectID `json:"productId" validate:"required"`
	Stock int `json:"stock" validate:"required,numeric,gt=0"`
}