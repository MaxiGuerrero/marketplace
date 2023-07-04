package infrastructure

import (
	response "marketplace/security-api/src/shared"
	utils "marketplace/security-api/src/shared/utils"
	models "marketplace/security-api/src/users/models"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
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
	businessErr := uc.service.CreateUser(req.Username,req.Password,req.Email,req.Role)
	if businessErr!=nil{
		return c.Status(400).JSON(response.BadRequest(businessErr.Error()))
	}
	return c.Status(200).JSON(response.OK())
}

func (uc UserController) UpdateUser(c *fiber.Ctx) error{
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	username := claims["username"].(string)
	req := models.UpdateUserRequest{}
	if parseError := c.BodyParser(&req); parseError!=nil{
		return c.Status(400).JSON(response.BadRequest(parseError.Error()))
	}
	if badSchema := utils.ValidateSchema(&req); badSchema != nil{
		return c.Status(400).JSON(response.BadRequest(badSchema.Error()))
	}
	businessErr := uc.service.UpdateUser(username,req.Email)
	if businessErr!=nil{
		return c.Status(400).JSON(response.BadRequest(businessErr.Error()))
	}
	return c.Status(200).JSON(response.OK())
}

func (uc UserController) DeleteUser(c *fiber.Ctx) error{
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	username := claims["username"].(string)
	businessErr := uc.service.DeleteUser(username)
	if businessErr!=nil{
		return c.Status(400).JSON(response.BadRequest(businessErr.Error()))
	}
	return c.Status(200).JSON(response.OK())
}

func (uc UserController) GetUsers(c *fiber.Ctx) error{
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	role := claims["role"].(string)
	if role != models.ADMIN.String() {
		return c.Status(401).JSON(response.Unauthorized())
	}
	return c.Status(200).JSON(uc.service.GetUsers())
}