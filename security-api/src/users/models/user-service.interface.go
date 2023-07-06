package models

// Interface to implement all method about users management.
type IUserService interface {
	CreateUser(username,password,email,role string) error
	UpdateUser(username string, email string) error
	DeleteUser(username string) error
	GetUsers() *Users
}