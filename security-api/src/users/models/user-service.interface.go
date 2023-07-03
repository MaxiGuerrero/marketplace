package models

type IUserService interface {
	CreateUser(username,password,email string) error
	UpdateUser(username string, email string) error
	DeleteUser(username string) error
	GetUsers() *Users
}