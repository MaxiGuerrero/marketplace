package infrastructure

import (
	"context"
	"log"
	mongo "marketplace/security-api/src/shared/database"
)

var ctx context.Context = context.Background()

type UserRepository struct{
	db mongo.DbConnector
}

func (u UserRepository) Create(username,password,email string) error{
	// u.db.GetCollection("user").InsertOne(ctx,)
	log.Println("User created!! =D")
	return nil
}