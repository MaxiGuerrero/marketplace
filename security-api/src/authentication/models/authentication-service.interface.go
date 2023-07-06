package models

// Interface to implement all method about auth management.
type IAuthenticationService interface {
	Login(username,password string) (*UserToken,error)
}