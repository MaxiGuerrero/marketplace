package infrastructure

import (
	middlewares "marketplace/security-api/src/shared/middlewares"

	"github.com/gofiber/fiber/v2"
)

// Register routes about management of users.
func RegisterRoutes(router fiber.Router, uc UserController){
	jwt := middlewares.NewAuthMiddleware()
	router.Post("/register",uc.CreateUser)
	router.Put("/users",jwt,uc.UpdateUser)
	router.Delete("/users",jwt,uc.DeleteUser)
	router.Get("/users",jwt,uc.GetUsers)
}

