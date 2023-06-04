package infrastructure

import (
	s "marketplace/security-api/src/authentication/service"

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
type Response struct{
	Message string `json:"message"`
} 

func (ac AuthenticationController) login(c *fiber.Ctx) error{
	result := ac.service.Login()
	return c.Status(200).JSON(Response{Message: result})
}