package models

type IUserService interface {
	CreateUser(username,password,email string) error
}