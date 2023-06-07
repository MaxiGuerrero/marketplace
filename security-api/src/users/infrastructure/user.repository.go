package infrastructure

import "log"

type UserRepository struct{}

func (u UserRepository) Create(username,password,email string) error{
	log.Println("User created!! =D")
	return nil
}