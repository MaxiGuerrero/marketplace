package infrastructure

import (
	"context"
	"log"
	database "marketplace/security-api/src/shared/database"
	model "marketplace/security-api/src/users/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx context.Context = context.Background()

type UserRepository struct{
	db database.DbConnector
}

func (u UserRepository) Create(username,password,email string){
	_,err := u.db.GetCollection("user").InsertOne(ctx,model.User{
		Username: username,
		Password: password,
		Email: email,
		Status: model.Status.String(model.Active),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil{
		log.Panicf("Error on create user document: %v",err.Error())
	}
	log.Printf("User %v has been created",username)
}

func (u UserRepository) GetByUsername(username string) *model.User{
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	userFound := new(model.User)
	err := u.db.GetCollection("user").FindOne(ctx,filter).Decode(&userFound)
	if err != nil { 
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Panicf("Error on GetByUsername: %v",err.Error())
	}
	return userFound
}

func (u UserRepository) Update(username, newEmail string){
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	update := bson.M{
		"$set": bson.M{
			"email": newEmail,
			"updatedat": time.Now(),
		},
	}
	_ , err := u.db.GetCollection("user").UpdateOne(ctx,filter,update)
	if err != nil {
		log.Panicf("Error on update user: %v", err)
	}
	log.Printf("User %v has been updated", username)
}

func (u UserRepository) Delete(username string){
	
}