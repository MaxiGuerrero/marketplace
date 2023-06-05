package infrastructure

import (
	s "marketplace/security-api/src/authentication/service"
	response "marketplace/security-api/src/shared"

	"github.com/gofiber/fiber/v2"
)

type AuthenticationController struct{
	service s.AuthenticationService
}

func NewAuthenticationController(service s.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{
		service,
	}
}

func (ac AuthenticationController) login(c *fiber.Ctx) error{
	result := ac.service.Login()
	return c.Status(200).JSON(response.Custom(result))
}