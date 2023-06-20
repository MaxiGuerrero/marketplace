package infrastructure

import (
	"github.com/gofiber/fiber/v2"
)


func RegisterRoutes(router fiber.Router, ac AuthenticationController){
	router.Post("/login",ac.Login)
}