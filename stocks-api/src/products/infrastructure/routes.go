package infrastructure

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router, productController *ProductController){
	router.Post("/products",productController.RegisterProduct)
	router.Put("/products",productController.UpdateProduct)
	router.Put("/products/stock",productController.UpdateStock)
}