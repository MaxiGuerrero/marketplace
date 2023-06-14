package models

type IAuthenticationRepository interface {
	GetByUsername(username string) *User
}