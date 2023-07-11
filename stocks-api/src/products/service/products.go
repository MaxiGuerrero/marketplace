package service

import (
	"errors"
	"marketplace/stocks-api/src/products/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService struct {
	ProductRepository models.IProductRepository
}

func NewProductService(productRepository models.IProductRepository) *ProductService{
	return &ProductService{productRepository}
}

func (ps *ProductService) RegisterProduct(name string, description string, price float32, stock int){
	ps.ProductRepository.RegisterProduct(name, description, price, stock)
}

func (ps *ProductService) UpdateProduct(productId primitive.ObjectID,name string, description string, price float32) error{
	product := ps.ProductRepository.GetProductById(productId)
	if product == nil {
		return errors.New("product does not exists")
	}
	ps.ProductRepository.UpdateProduct(productId,name,description,price)
	return nil
}

func (ps *ProductService) UpdateStock(productId primitive.ObjectID,stock int) error {
	product := ps.ProductRepository.GetProductById(productId)
	if product == nil {
		return errors.New("product does not exists")
	}
	ps.ProductRepository.UpdateStock(productId,stock)
	return nil
}