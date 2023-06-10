package models

type IUserRepository interface {
	Create(username,password,email string) error
	// Update()
	// Delete()
	// GetById() *User
	GetByUsername(username string) (*User,error)
}