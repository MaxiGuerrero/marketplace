package infrastructure

import (
	service "marketplace/security-api/src/authentication/service"
)

type Dependencies struct{
	AuthenticationController *AuthenticationController
}

func InitializeDependencies() *Dependencies{
	return &Dependencies{
		AuthenticationController: NewAuthenticationController(*service.NewAuthenticationService()),
	}
}