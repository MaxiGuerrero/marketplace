package infrastructure

import (
	"log"
	"marketplace/security-api/src/authentication/models"
	"marketplace/security-api/src/shared/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthenticationRepository struct{
	db database.DbConnector
}

func (ar AuthenticationRepository) GetByUsername(username string) *models.User {
	log.Panicf("Test panic")
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	userFound := new(models.User)
	err := ar.db.GetCollection("user").FindOne(ctx,filter).Decode(&userFound)
	if err != nil { 
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Panicf("Error on database: %v",err.Error())
	}
	return userFound
}