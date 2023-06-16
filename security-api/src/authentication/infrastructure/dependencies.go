package infrastructure

import (
	service "marketplace/security-api/src/authentication/service"
	config "marketplace/security-api/src/shared"
	mongo "marketplace/security-api/src/shared/database"
	encrpyter "marketplace/security-api/src/shared/encrypter"
)

type Dependencies struct{
	AuthenticationController *AuthenticationController
}

func InitializeDependencies(db *mongo.DbConnector) *Dependencies{
	return &Dependencies{
		AuthenticationController: NewAuthenticationController(*service.NewAuthenticationService(encrpyter.CreateEncrypter(config.GetConfig().CostAlgorithmic),AuthenticationRepository{*db},&JWTBuilder{})),
	}
}