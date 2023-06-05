package main

import (
	authentication "marketplace/security-api/src/authentication/infrastructure"
	s "marketplace/security-api/src/server"
	users "marketplace/security-api/src/users/infrastructure"
)

func main(){
	// Dependencies' containers by each module - run injections
	authDependencies := authentication.InitializeDependencies()
	usersDependencies := users.InitializeDependencies()
	// Create server
	server := s.CreateServer(8080)
	// Register Routes
	authentication.RegisterRoutes(server.App,*authDependencies.AuthenticationController)
	users.RegisterRoutes(server.App,*usersDependencies.UserController)
	// Start Server
	server.StartServer()
}