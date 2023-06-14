package infrastructure

import (
	response "marketplace/security-api/src/shared"
	utils "marketplace/security-api/src/shared/utils"
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
	req := models.CreateUserRequest{}
	if parseError := c.BodyParser(&req); parseError!=nil{
		return c.Status(400).JSON(response.BadRequest(parseError.Error()))
	}
	if badSchema := utils.ValidateSchema(&req); badSchema != nil{
		return c.Status(400).JSON(response.BadRequest(badSchema.Error()))
	}
	businessErr := uc.service.CreateUser(req.Username,req.Password,req.Email)
	if businessErr!=nil{
		return c.Status(400).JSON(response.BadRequest(businessErr.Error()))
	}
	return c.Status(200).JSON(response.OK())
}