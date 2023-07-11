package main

import (
	"context"
	"marketplace/stocks-api/src/healthcheck"
	products "marketplace/stocks-api/src/products/infrastructure"
	s "marketplace/stocks-api/src/server"
	config "marketplace/stocks-api/src/shared"
	mongo "marketplace/stocks-api/src/shared/database"
)

func main(){
	ctx := context.Background()
	connector := mongo.CreateDbConnector(ctx,config.GetConfig().DbConnection,config.GetConfig().Database)
	// Dependencies' containers by each module - run injections
	productsDependencies := products.InitializeDependencies(connector)
	// Create server
	server := s.CreateServer(config.GetConfig().Port,true)
	// Register Routes
	products.RegisterRoutes(server.App,productsDependencies.ProductsController)
	healthcheck.RegisterRoutes(server.App)
	// Start Server
	server.StartServer()
}