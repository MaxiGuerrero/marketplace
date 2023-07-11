package infrastructure

import (
	middlewares "marketplace/stocks-api/src/shared/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, productController *ProductController){
	auth := middlewares.NewAuthMiddleware()
	router.Post("/products",auth,productController.RegisterProduct)
	router.Put("/products",auth,productController.UpdateProduct)
	router.Put("/products/stock",auth,productController.UpdateStock)
}