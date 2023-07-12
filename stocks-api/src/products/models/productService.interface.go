package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type IProductService interface {
	RegisterProduct(name string, description string, price float32, stock int)
	UpdateProduct(productId primitive.ObjectID,name string, description string, price float32) error
	UpdateStock(productId primitive.ObjectID,stock int) error
	GetAll() *[]Product
}