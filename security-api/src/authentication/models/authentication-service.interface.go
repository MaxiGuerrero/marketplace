package models

type IAuthenticationService interface {
	Login(username,password string) (*UserToken,error)
}