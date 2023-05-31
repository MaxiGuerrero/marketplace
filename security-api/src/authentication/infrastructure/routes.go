package infrastructure

import (
	"github.com/gofiber/fiber/v2"
)


func RegisterRoutes(router fiber.Router){
	router.Get("/",func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}

