package tests

import (
	"errors"
	infrastructure "marketplace/security-api/src/users/infrastructure"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
)

type FakeUserService struct {
	mock.Mock
}

func (u *FakeUserService) CreateUser(username,password,email string) error{
	args := u.Called(username,password,email)
	return args.Error(0)
}

func TestUserController(t *testing.T){
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)
	fakeServiceUser := &FakeUserService{}
	userController := infrastructure.NewUserController(fakeServiceUser)
	t.Run("User created sucessfully, return 200",func(t *testing.T){
		// Arrenge
		fakeServiceUser.On("CreateUser","test","test","test@gmail.com").Return(nil).Once()

		// Act
		err := userController.CreateUser(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t, `{"message":"operation sucessfully"}`, string(c.Response().Body()),"Must return status 200 with message")
		utils.AssertEqual(t,200,c.Response().StatusCode())
	})
	t.Run("User already exists, return 400",func(t *testing.T){
		// Arrenge
		fakeServiceUser.On("CreateUser","test","test","test@gmail.com").Return(errors.New("username already exists, please use another")).Once()

		// Act
		err := userController.CreateUser(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t, `{"message":"username already exists, please use another"}`, string(c.Response().Body()),"Must return status 400 with message")
		utils.AssertEqual(t,400,c.Response().StatusCode())
	})
}