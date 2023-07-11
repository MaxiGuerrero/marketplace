package infrastructure

import (
	"marketplace/stocks-api/src/products/service"
	mongo "marketplace/stocks-api/src/shared/database"
)

// This struct is responsable of manage injection dendencies of the system.
type Dependencies struct{
	ProductsController *ProductController
}


// Initialize dependencies injecting objects and configurations.
func InitializeDependencies(db *mongo.DbConnector) *Dependencies{
	return &Dependencies{
		NewProductController(service.NewProductService(NewProductRepository(db))),
	}
}