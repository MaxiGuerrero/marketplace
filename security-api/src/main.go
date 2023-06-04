package main

import (
	authentication "marketplace/security-api/src/authentication/infrastructure"
	s "marketplace/security-api/src/server"
)

func main(){
	authDependencies := authentication.InitializeDependencies()
	server := s.CreateServer(8080)
	authentication.RegisterRoutes(server.App,*authDependencies.AuthenticationController)
	server.StartServer()
}