package infrastructure

import (
	service "marketplace/security-api/src/users/service"
)

type Dependencies struct{
	UserController *UserController
}

func InitializeDependencies() *Dependencies{
	return &Dependencies{
		UserController: NewUserController(*service.NewUserService()),
	}
}