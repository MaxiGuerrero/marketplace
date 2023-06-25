package models

type IUserRepository interface {
	Create(username,password,email string)
	Update(username, email string)
	Delete(usernae string)
	// GetById() *User
	GetByUsername(username string) *User
}