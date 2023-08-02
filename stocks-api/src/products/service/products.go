package service

import (
	"errors"
	"log"
	"marketplace/stocks-api/src/products/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProductService is a service that is responsable to manage products in the system.
// All service manage the business logical.
type ProductService struct {
	ProductRepository models.IProductRepository
}

// Create an instance of the ProductService with injection dependencies.
func NewProductService(productRepository models.IProductRepository) *ProductService{
	return &ProductService{productRepository}
}

// Create a product in the system.
func (ps *ProductService) RegisterProduct(name string, description string, price float32, stock int){
	ps.ProductRepository.RegisterProduct(name, description, price, stock)
}

// Update a product in the system. 
// Is possible return an business error if the product doesn't exists.
func (ps *ProductService) UpdateProduct(productId primitive.ObjectID,name string, description string, price float32) error{
	product := ps.ProductRepository.GetProductById(productId)
	if product == nil {
		return errors.New("product does not exists")
	}
	ps.ProductRepository.UpdateProduct(productId,name,description,price)
	return nil
}


// Update stock of one product in the system. 
// Is possible return an business error if the product doesn't exists.
func (ps *ProductService) UpdateStock(productId primitive.ObjectID,stock int) error {
	product := ps.ProductRepository.GetProductById(productId)
	if product == nil {
		return errors.New("product does not exists")
	}
	ps.ProductRepository.UpdateStock(productId,stock)
	return nil
}

// Get an array of product in the system.
func (ps *ProductService) GetAll() *[]models.Product{
	return ps.ProductRepository.GetAll()
}

// Get product by ID.
// Is possible return an business error if the product doesn't exists.
func (ps *ProductService) GetProductById(productId primitive.ObjectID) *models.Product {
	return ps.ProductRepository.GetProductById(productId)
}

func (ps *ProductService) ReciveCheckout(products *[]models.ProductOnCart){
	for _,product := range *products {
		productFound := ps.ProductRepository.GetProductById(product.ProductId)
		newStock := productFound.Stock - product.Amount
		if (newStock < 0){
			log.Printf("Stock on product %v not reach", productFound.Name);
			return
		}
		ps.ProductRepository.UpdateStock(product.ProductId,newStock)
	}
}