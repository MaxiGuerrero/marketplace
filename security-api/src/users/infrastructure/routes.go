package infrastructure

import (
	"github.com/gofiber/fiber/v2"
)


func RegisterRoutes(router fiber.Router, uc UserController){
	router.Get("/users",uc.createUser)
}

