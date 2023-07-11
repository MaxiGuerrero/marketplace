package infrastructure

import (
	responses "marketplace/security-api/src/shared"
	middlewares "marketplace/security-api/src/shared/middlewares"

	"github.com/gofiber/fiber/v2"
)

// Register routes about management of authentication.
func RegisterRoutes(router fiber.Router, ac AuthenticationController){
	jwt := middlewares.NewAuthMiddleware()
	router.Post("/login",ac.Login)
	router.Post("/token/validate",jwt,func(c *fiber.Ctx) error {
		return c.Status(200).JSON(responses.TokenValidated())
	})
}