package infrastructure

import (
	config "marketplace/security-api/src/shared"
	mongo "marketplace/security-api/src/shared/database"
	encrpyter "marketplace/security-api/src/shared/encrypter"
	service "marketplace/security-api/src/users/service"
)

type Dependencies struct{
	UserController *UserController
}

func InitializeDependencies(db *mongo.DbConnector) *Dependencies{
	return &Dependencies{
		UserController: NewUserController(*service.NewUserService(&UserRepository{*db},encrpyter.CreateEncrypter(config.GetConfig().CostAlgorithmic))),
	}
}