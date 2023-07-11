package infrastructure

import (
	"context"
	"log"
	"marketplace/stocks-api/src/products/models"
	"marketplace/stocks-api/src/shared/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	db *database.DbConnector
}

func NewProductRepository(db *database.DbConnector)*ProductRepository{
	return &ProductRepository{db}
}

var ctx context.Context = context.Background()

func (pr *ProductRepository) RegisterProduct(name string, description string, price float32, stock int) {
	_,err := pr.db.GetCollection("product").InsertOne(ctx,models.Product{
		Name: name,
		Description: description,
		Stock: stock,
		Price: price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Panicf("Error on create product document: %v",err.Error())
	}
	log.Printf("Product %v has been registered",name)
}

func (pr *ProductRepository) UpdateProduct(productId primitive.ObjectID,name string, description string, price float32){
	filter := bson.D{primitive.E{Key: "_id", Value: productId}}
	update := bson.M{
		"$set": bson.M{
			"name": name,
			"description": description,
			"price": price,
			"updatedat": time.Now(),
		},
	}
	_ , err := pr.db.GetCollection("product").UpdateOne(ctx,filter,update)
	if err != nil {
		log.Panicf("Error on update product: %v", err)
	}
	log.Printf("Product %v has been updated", name)
}

func (pr *ProductRepository) GetProductById(productId primitive.ObjectID) *models.Product{
	product := &models.Product{}
	filter := bson.D{primitive.E{Key: "_id", Value: productId}}
	result := pr.db.GetCollection("product").FindOne(ctx,filter)
	if result.Err() != nil {
		if result.Err()  == mongo.ErrNoDocuments {
			return nil
		}
		log.Panicf("Error on get product by id: %v", result.Err().Error())
	}
	err := result.Decode(product)
	if err != nil {
		log.Panicf("Error on decode product in get product by id: %v",err)
	}
	return product
}

func (pr *ProductRepository) UpdateStock(productId primitive.ObjectID, stock int){
	filter := bson.D{primitive.E{Key: "_id", Value: productId}}
	update := bson.M{
		"$set": bson.M{
			"stock": stock,
			"updatedat": time.Now(),
		},
	}
	_ , err := pr.db.GetCollection("product").UpdateOne(ctx,filter,update)
	if err != nil {
		log.Panicf("Error on update product: %v", err)
	}
	log.Printf("Product id %v has been updated its stock", productId.Hex())
}