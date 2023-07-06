package infrastructure

import (
	"context"
	"marketplace/security-api/src/authentication/models"
	response "marketplace/security-api/src/shared"
	"marketplace/security-api/src/shared/utils"

	"github.com/gofiber/fiber/v2"
)

var ctx context.Context = context.Background()

// AuthenticationController is a controller that is responsable to manage request and responses.
type AuthenticationController struct{
	service models.IAuthenticationService
}

// Create an instance of the AuthenticationController with injection dependencies.
func NewAuthenticationController(service models.IAuthenticationService) *AuthenticationController {
	return &AuthenticationController{
		service,
	}
}

// Manage request and responses about login an user in the system.
func (ac AuthenticationController) Login(c *fiber.Ctx) error{
	req := models.LoginRequest{}
	if parseError := c.BodyParser(&req); parseError!=nil{
		return c.Status(400).JSON(response.BadRequest(parseError.Error()))
	}
	if badSchema := utils.ValidateSchema(&req); badSchema != nil{
		return c.Status(400).JSON(response.BadRequest(badSchema.Error()))
	}
	userToken,businessError := ac.service.Login(req.Username,req.Password)
	if businessError != nil {
		return c.Status(401).JSON(response.BadRequest(businessError.Error()))
	}
	return c.Status(200).JSON(userToken)
}