package infrastructure

import (
	server "marketplace/security-api/src/server"

	"github.com/gofiber/fiber/v2"
)

var Routes []server.Route


func GetRoutes() []server.Route{
	var route = server.Route{
		Method: "get",
		Path: "/",
		Handler:  func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		},
	}
	return append(Routes, route)
}

