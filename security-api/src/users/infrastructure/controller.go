package infrastructure

import (
	response "marketplace/security-api/src/shared"
	s "marketplace/security-api/src/users/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct{
	service s.UserService
}

func NewUserController(service s.UserService) *UserController {
	return &UserController{
		service,
	}
}

func (uc UserController) createUser(c *fiber.Ctx) error{
	err := uc.service.CreateUser("test","test","test@gmail.com")
	if err!=nil{
		return c.Status(400).JSON(response.BadRequest(err.Error()))
	}
	return c.Status(200).JSON(response.OK())
}