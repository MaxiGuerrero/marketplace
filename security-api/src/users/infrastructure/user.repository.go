package infrastructure

import "log"

type UserRepository struct{}

func (u UserRepository) Create(){
	log.Println("User created!! =D")
}