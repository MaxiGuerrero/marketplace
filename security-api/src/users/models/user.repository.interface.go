package models

// Interface to implements method about user query management .
type IUserRepository interface {
	Create(username,password,email,role string)
	Update(username, email string)
	Delete(usernae string)
	Get() *Users
	GetByUsername(username string) *User
}