package infrastructure

import (
	"marketplace/security-api/src/shared"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func NewAuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(shared.GetConfig().JWTSecret),
		ErrorHandler: func(c *fiber.Ctx,err error) error{
			return c.Status(401).JSON(shared.Unauthorized())
		},
	})
}
