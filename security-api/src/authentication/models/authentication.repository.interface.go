package models

// Interface to implements method about auth query management .
type IAuthenticationRepository interface {
	GetByUsername(username string) *User
}