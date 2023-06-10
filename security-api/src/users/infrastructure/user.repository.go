package infrastructure

import (
	"context"
	"log"
	mongo "marketplace/security-api/src/shared/database"
	model "marketplace/security-api/src/users/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx context.Context = context.Background()

type UserRepository struct{
	db mongo.DbConnector
}

func (u UserRepository) Create(username,password,email string) error{
	_,err := u.db.GetCollection("user").InsertOne(ctx,model.User{
		Username: username,
		Password: password,
		Email: email,
		Status: model.Status.String(model.Active),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil{
		return err
	}
	log.Printf("User %v has been created",username)
	return nil
}

func (u UserRepository) GetByUsername(username string) (*model.User,error){
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	result := u.db.GetCollection("user").FindOne(ctx,filter)
	if result.Err() != nil {
		return nil,result.Err()
	}
	userFound := new(model.User)
	err := result.Decode(&userFound)
	if err != nil {
		return nil, err
	}
	return userFound,nil
}