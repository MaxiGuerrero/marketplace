package infrastructure

import (
	config "marketplace/security-api/src/shared"
	encrpyter "marketplace/security-api/src/shared/encrypter"
	service "marketplace/security-api/src/users/service"
)

type Dependencies struct{
	UserController *UserController
}

func InitializeDependencies() *Dependencies{
	return &Dependencies{
		UserController: NewUserController(*service.NewUserService(&UserRepository{},encrpyter.CreateEncrypter(config.GetConfig().CostAlgorithmic))),
	}
}