package infrastructure

import (
	"github.com/gofiber/fiber/v2"
)

// Register routes about management of authentication.
func RegisterRoutes(router fiber.Router, ac AuthenticationController){
	router.Post("/login",ac.Login)
}