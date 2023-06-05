package main

import (
	authentication "marketplace/security-api/src/authentication/infrastructure"
	s "marketplace/security-api/src/server"
	users "marketplace/security-api/src/users/infrastructure"
)

func main(){
	authDependencies := authentication.InitializeDependencies()
	usersDependencies := users.InitializeDependencies()
	server := s.CreateServer(8080)
	authentication.RegisterRoutes(server.App,*authDependencies.AuthenticationController)
	users.RegisterRoutes(server.App,*usersDependencies.UserController)
	server.StartServer()
}