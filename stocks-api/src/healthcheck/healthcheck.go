package healthcheck

import (
	responses "marketplace/stocks-api/src/shared"

	"github.com/gofiber/fiber/v2"
)

// Register route to health check endpoint and implement its logical function response that must be {message:"OK"}.
func RegisterRoutes(router fiber.Router){
	router.Get("/healthcheck",func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&responses.Response{Message: "OK"})
	})
}