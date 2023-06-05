package models

type ICredentialRepository interface {
	Login(username string, password string)
}