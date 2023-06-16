package middleware

import (
	infrastructure "marketplace/security-api/src/authentication/infrastructure"
	shared "marketplace/security-api/src/shared"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var jwtBuilder = infrastructure.JWTBuilder{}

func ValidateToken(c *fiber.Ctx) error {
	authorizationHeader := c.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer "){
		return c.Status(fiber.StatusUnauthorized).JSON(shared.Unauthorized())
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	user,err := jwtBuilder.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(shared.Unauthorized())
	}
	c.Locals("user",user)
	return c.Next()
}
