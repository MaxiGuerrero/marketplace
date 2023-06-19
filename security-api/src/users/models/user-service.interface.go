package models

type IUserService interface {
	CreateUser(username,password,email string) error
	UpdateUser(username string, email string) error
}