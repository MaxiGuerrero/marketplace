package infrastructure

import (
	"context"
	"marketplace/security-api/src/authentication/models"
	response "marketplace/security-api/src/shared"
	"marketplace/security-api/src/shared/utils"

	"github.com/gofiber/fiber/v2"
)

var ctx context.Context = context.Background()

type AuthenticationController struct{
	service models.IAuthenticationService
}

func NewAuthenticationController(service models.IAuthenticationService) *AuthenticationController {
	return &AuthenticationController{
		service,
	}
}

func (ac AuthenticationController) login(c *fiber.Ctx) error{
	req := models.LoginRequest{}
	if parseError := c.BodyParser(&req); parseError!=nil{
		return c.Status(400).JSON(response.BadRequest(parseError.Error()))
	}
	if badSchema := utils.ValidateSchema(&req); badSchema != nil{
		return c.Status(400).JSON(response.BadRequest(badSchema.Error()))
	}
	userToken,businessError := ac.service.Login(req.Username,req.Password)
	if businessError != nil {
		return c.Status(400).JSON(response.BadRequest(businessError.Error()))
	}
	return c.Status(200).JSON(userToken)
}