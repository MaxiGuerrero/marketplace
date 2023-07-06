package main

import (
	"context"
	"log"
	authentication "marketplace/security-api/src/authentication/infrastructure"
	"marketplace/security-api/src/healthcheck"
	s "marketplace/security-api/src/server"
	config "marketplace/security-api/src/shared"
	mongo "marketplace/security-api/src/shared/database"
	utils "marketplace/security-api/src/shared/utils"
	users "marketplace/security-api/src/users/infrastructure"
)

func main(){
	a := 1
	a = 2
	log.Printf("%v", a)
	ctx := context.Background()
	connector := mongo.CreateDbConnector(ctx,config.GetConfig().DbConnection,config.GetConfig().Database)
	// Dependencies' containers by each module - run injections
	authDependencies := authentication.InitializeDependencies(connector)
	usersDependencies := users.InitializeDependencies(connector)
	// Create server
	server := s.CreateServer(config.GetConfig().Port,true)
	// Register Routes
	authentication.RegisterRoutes(server.App,*authDependencies.AuthenticationController)
	users.RegisterRoutes(server.App,*usersDependencies.UserController)
	healthcheck.RegisterRoutes(server.App)
	// Register custom validation
	utils.RegisterValidation()
	// Start Server
	server.StartServer()
}