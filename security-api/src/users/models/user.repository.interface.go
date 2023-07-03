package models

type IUserRepository interface {
	Create(username,password,email string)
	Update(username, email string)
	Delete(usernae string)
	Get() *Users
	GetByUsername(username string) *User
}