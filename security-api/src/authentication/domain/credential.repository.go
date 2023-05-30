package domain

type CredentialRepository interface {
	Login(username string, password string)
}