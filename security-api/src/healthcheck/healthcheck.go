package healthcheck

import (
	"github.com/gofiber/fiber/v2"
)

type healthcheckResponse struct {
	Message string
}

func RegisterRoutes(router fiber.Router){
	router.Get("/healthcheck",func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&healthcheckResponse{Message: "OK"})
	})
}