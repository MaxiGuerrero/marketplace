package tests

import (
	"errors"
	"fmt"
	"marketplace/security-api/src/authentication/infrastructure"
	"marketplace/security-api/src/authentication/models"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FakeAuthService struct {
	mock.Mock
}

func (fa *FakeAuthService) Login(username,password string) (*models.UserToken,error){
	args := fa.Called(username,password)
	if args.Get(0) == nil{
		return nil,args.Error(1)
	}
	return args.Get(0).(*models.UserToken),args.Error(1)
}

func TestLoginController(t *testing.T){
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)
	fakeServiceUser := &FakeAuthService{}
	userController := infrastructure.NewAuthenticationController(fakeServiceUser)
	t.Run("Login sucessfully, return 200",func(t *testing.T){
		token := &models.UserToken{
			UserId: primitive.NewObjectID(),
			Token: "123455",
		}
		// Arrenge
		fakeServiceUser.On("Login","test","test").Return(token,nil).Once()
		c.Request().Header.SetContentType("application/json")
        c.Request().SetBody([]byte(`{"username":"test","password":"test"}`))
		expected := fmt.Sprintf(`{"userId":"%s","token":"%s"}`,token.UserId.Hex(),token.Token)
		// Act
		err := userController.Login(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t, expected, string(c.Response().Body()),"Must return status 200 with message")
		utils.AssertEqual(t,200,c.Response().StatusCode())
	})
	t.Run("Unauthorized, return 401",func(t *testing.T){
		// Arrenge
		fakeServiceUser.On("Login","test","test").Return(nil,errors.New("username or password is incorrect, try again")).Once()
		c.Request().Header.SetContentType("application/json")
        c.Request().SetBody([]byte(`{"username":"test","password":"test"}`))
		// Act
		err := userController.Login(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t,`{"message":"username or password is incorrect, try again"}`, string(c.Response().Body()))
		utils.AssertEqual(t,401,c.Response().StatusCode())
	})
}