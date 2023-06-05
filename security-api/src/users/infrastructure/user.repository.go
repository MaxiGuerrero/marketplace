package infrastructure

import "log"

type UserRepository struct{}

func (u UserRepository) Create(username,password,email string){
	log.Println("User created!! =D")
}