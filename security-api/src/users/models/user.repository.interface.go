package models

type IUserRepository interface {
	Create(username,password,email string)
	// Update()
	// Delete()
	// GetById() *User
	GetByUsername(username string) *User
}