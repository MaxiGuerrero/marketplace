package infrastructure

import (
	"context"
	"log"
	database "marketplace/security-api/src/shared/database"
	models "marketplace/security-api/src/users/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx context.Context = context.Background()

// Repository that is responsable to manage queries and database connection about users.
// Also is responsable to manage database error connections, throw panic if a error exists.
type UserRepository struct{
	db database.DbConnector
}

// Register an user in the mongo database.
func (u UserRepository) Create(username,password,email,role string){
	_,err := u.db.GetCollection("user").InsertOne(ctx,models.User{
		Username: username,
		Password: password,
		Email: email,
		Status: models.Status.String(models.Active),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Role: role,
	})
	if err != nil{
		log.Panicf("Error on create user document: %v",err.Error())
	}
	log.Printf("User %v has been created",username)
}

// Get an user from the mongo database.
func (u UserRepository) GetByUsername(username string) *models.User{
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	userFound := new(models.User)
	err := u.db.GetCollection("user").FindOne(ctx,filter).Decode(&userFound)
	if err != nil { 
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Panicf("Error on GetByUsername: %v",err.Error())
	}
	return userFound
}

// Update an exists user collection from the mongo database.
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

// Logical delete an exists and active user updating the user collection setting delete date and status inactive.
func (u UserRepository) Delete(username string){
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	update := bson.M{
		"$set": bson.M{
			"deletedat": time.Now(),
			"status": models.Inactive.String(),
		},
	}
	_ , err := u.db.GetCollection("user").UpdateOne(ctx,filter,update)
	if err != nil {
		log.Panicf("Error on delete user: %v", err)
	}
	log.Printf("User %v has been deleted", username)
}

// Get list of users from the mongo database.
func (u UserRepository) Get() *models.Users{
	var users models.Users
	cursor , err := u.db.GetCollection("user").Find(ctx,bson.D{})
	if err != nil {
		log.Panicf("Error on get users: %v", err)
	}
	if err = cursor.All(ctx,&users); err != nil {
		log.Panicf("Error on get users: %v", err)
	}
	return &users
}