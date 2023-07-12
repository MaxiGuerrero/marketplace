package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type IProductRepository interface {
	RegisterProduct(name string, description string, price float32, stock int)
	UpdateProduct(productId primitive.ObjectID,name string, description string, price float32)
	GetProductById(productId primitive.ObjectID) *Product
	UpdateStock(productId primitive.ObjectID, stock int)
	GetAll() *[]Product
}