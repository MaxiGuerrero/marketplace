package main

import (
	"context"
	authentication "marketplace/security-api/src/authentication/infrastructure"
	s "marketplace/security-api/src/server"
	config "marketplace/security-api/src/shared"
	mongo "marketplace/security-api/src/shared/database"
	users "marketplace/security-api/src/users/infrastructure"
)

func main(){
	ctx := context.Background()
	connector := mongo.CreateDbConnector(ctx,config.GetConfig().DbConnection,config.GetConfig().Database)
	// Dependencies' containers by each module - run injections
	authDependencies := authentication.InitializeDependencies(connector)
	usersDependencies := users.InitializeDependencies(connector)
	// Create server
	server := s.CreateServer(8080)
	// Register Routes
	authentication.RegisterRoutes(server.App,*authDependencies.AuthenticationController)
	users.RegisterRoutes(server.App,*usersDependencies.UserController)
	// Start Server
	server.StartServer()
}