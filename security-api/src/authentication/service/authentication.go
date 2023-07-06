package service

import (
	"errors"
	"marketplace/security-api/src/authentication/models"
)

// AuthenticationService is a service that is responsable to manage the authentication user in the system.
// All service manage the business logical.
type AuthenticationService struct{
	encrypter models.IEncrypter
	authenticationRepository models.IAuthenticationRepository
	jwtBuilder models.JWTBuilder
}

// Create an instance of the AuthenticationService with injection dependencies
func NewAuthenticationService(encrypter models.IEncrypter, authenticationRepository models.IAuthenticationRepository, jwtBuilder models.JWTBuilder) *AuthenticationService{
	return &AuthenticationService{encrypter, authenticationRepository, jwtBuilder}
}

// Implementation about login a user in the system.
func (as AuthenticationService) Login(username, password string) (*models.UserToken,error){
	userFound := as.authenticationRepository.GetByUsername(username)
	if userFound == nil {
		return nil,errors.New("username or password is incorrect, try again")
	}
	if userFound.Status == models.Inactive.String() {
		return nil,errors.New("username or password is incorrect, try again")
	}
	if !as.encrypter.Compare([]byte(userFound.Password),[]byte(password)) {
		return nil,errors.New("username or password is incorrect, try again")
	}
	token := as.jwtBuilder.BuildToken(&models.Payload{
		UserId: userFound.ID,
		Username: userFound.Username,
		CreatedAt: userFound.CreatedAt,
		UpdatedAt: userFound.UpdatedAt,
		DeletedAt: userFound.DeletedAt,
		Role: userFound.Role,
	})
	return &models.UserToken{UserId: userFound.ID,Token: token},nil
}