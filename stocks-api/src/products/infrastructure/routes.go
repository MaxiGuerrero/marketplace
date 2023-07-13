package infrastructure

import (
	middlewares "marketplace/stocks-api/src/shared/middlewares"

	"github.com/gofiber/fiber/v2"
)

// Register routes about management of products.
func RegisterRoutes(router fiber.Router, productController *ProductController){
	auth := middlewares.NewAuthMiddleware()
	router.Post("/products",auth,productController.RegisterProduct)
	router.Put("/products",auth,productController.UpdateProduct)
	router.Put("/products/stock",auth,productController.UpdateStock)
	router.Get("/products",auth,productController.GetProducts)
	router.Get("/products/:id",auth,func(c *fiber.Ctx) error {
		c.Locals("id",c.Params("id"))
		return c.Next()
	},productController.GetProduct)
}