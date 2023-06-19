package models

type IUserRepository interface {
	Create(username,password,email string)
	Update(username, email string)
	// Delete()
	// GetById() *User
	GetByUsername(username string) *User
}