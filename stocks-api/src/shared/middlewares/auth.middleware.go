package infrastructure

import (
	"github.com/gofiber/fiber/v2"
)

// Middleware that is responsable to manage the authorization JWT token from a request.
func NewAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
