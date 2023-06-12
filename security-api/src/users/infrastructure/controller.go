package infrastructure

import (
	response "marketplace/security-api/src/shared"
	models "marketplace/security-api/src/users/models"

	"github.com/gofiber/fiber/v2"
)

type UserController struct{
	service models.IUserService
}

func NewUserController(service models.IUserService) *UserController {
	return &UserController{
		service,
	}
}

func (uc UserController) CreateUser(c *fiber.Ctx) error{
	err := uc.service.CreateUser("test","test","test@gmail.com")
	if err!=nil{
		return c.Status(400).JSON(response.BadRequest(err.Error()))
	}
	return c.Status(200).JSON(response.OK())
}