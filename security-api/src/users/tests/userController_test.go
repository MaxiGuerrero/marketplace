package tests

import (
	"encoding/json"
	"errors"
	infrastructure "marketplace/security-api/src/users/infrastructure"
	models "marketplace/security-api/src/users/models"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FakeUserService struct {
	mock.Mock
}

func (u *FakeUserService) CreateUser(username,password,email string) error{
	args := u.Called(username,password,email)
	return args.Error(0)
}

func (u *FakeUserService) UpdateUser(username,email string) error{
	args := u.Called(username,email)
	return args.Error(0)
}

func (u *FakeUserService) DeleteUser(username string) error{
	args := u.Called(username)
	return args.Error(0)
}

func (u *FakeUserService) GetUsers() *models.Users {
	args := u.Called()
	return args.Get(0).(*models.Users)
}

func TestCreateUser(t *testing.T){
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)
	fakeServiceUser := &FakeUserService{}
	userController := infrastructure.NewUserController(fakeServiceUser)
	t.Run("User created sucessfully, return 200",func(t *testing.T){
		// Arrenge
		fakeServiceUser.On("CreateUser","test","test","test@gmail.com").Return(nil).Once()
		c.Request().Header.SetContentType("application/json")
        c.Request().SetBody([]byte(`{"username":"test","password":"test","email":"test@gmail.com"}`))
		// Act
		err := userController.CreateUser(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t, `{"message":"Successful operation"}`, string(c.Response().Body()),"Must return status 200 with message")
		utils.AssertEqual(t,200,c.Response().StatusCode())
	})
	t.Run("User already exists, return 400",func(t *testing.T){
		// Arrenge
		fakeServiceUser.On("CreateUser","test","test","test@gmail.com").Return(errors.New("username already exists, please use another")).Once()
		c.Request().Header.SetContentType("application/json")
        c.Request().SetBody([]byte(`{"username":"test","password":"test","email":"test@gmail.com"}`))
		// Act
		err := userController.CreateUser(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t, `{"message":"username already exists, please use another"}`, string(c.Response().Body()),"Must return status 400 with message")
		utils.AssertEqual(t,400,c.Response().StatusCode())
	})
	t.Run("Error body parser, return 400",func(t *testing.T){
		// Arrenge
		fakeServiceUser.On("CreateUser","test","test","test@gmail.com").Return(errors.New("username already exists, please use another")).Once()
		c.Request().Header.SetContentType("application/json")
        c.Request().SetBody([]byte(`{"usern$me":"test","passw0rd":"test","em4il":"test@gmail.com"}`))
		// Act
		err := userController.CreateUser(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t,400,c.Response().StatusCode())
	})
	t.Run("Error bad schema body, return 400",func(t *testing.T){
		// Arrenge
		fakeServiceUser.On("CreateUser","test","test","test@gmail.com").Return(errors.New("username already exists, please use another")).Once()
		c.Request().Header.SetContentType("application/json")
        c.Request().SetBody([]byte(`{"username":"t","password":"t","email":"test"}`))
		// Act
		err := userController.CreateUser(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t,400,c.Response().StatusCode())
	})
}

func TestUpdateUser(t *testing.T){
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)
	fakeServiceUser := &FakeUserService{}
	userController := infrastructure.NewUserController(fakeServiceUser)
	token := &jtoken.Token{
		Claims: jtoken.MapClaims{
			"username":"test",
		},
	}
	t.Run("User updated sucessfully, return 200",func(t *testing.T){
		// Arrenge
		fakeServiceUser.On("UpdateUser","test","test@gmail.com").Return(nil).Once()
		c.Request().Header.SetContentType("application/json")
        c.Request().SetBody([]byte(`{"email":"test@gmail.com"}`))
		c.Locals("user",token)
		// Act
		err := userController.UpdateUser(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t, `{"message":"Successful operation"}`, string(c.Response().Body()),"Must return status 200 with message")
		utils.AssertEqual(t,200,c.Response().StatusCode())
	})
		t.Run("User does not exists, return 400",func(t *testing.T){
		// Arrenge
		fakeServiceUser.On("UpdateUser","test","test@gmail.com").Return(errors.New("user does not exist")).Once()
		c.Request().Header.SetContentType("application/json")
        c.Request().SetBody([]byte(`{"email":"test@gmail.com"}`))
		c.Locals("user",token)
		// Act
		err := userController.UpdateUser(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t, `{"message":"user does not exist"}`, string(c.Response().Body()),"Must return status 400 with message")
		utils.AssertEqual(t,400,c.Response().StatusCode())
	})
	t.Run("Error body parser, return 400",func(t *testing.T){
		// Arrenge
		fakeServiceUser.On("UpdateUser","test","test@gmail.com").Return(nil).Once()
		c.Request().Header.SetContentType("application/json")
        c.Request().SetBody([]byte(`{"em$il":"test@gmail.com"}`))
		c.Locals("user",token)
		// Act
		err := userController.UpdateUser(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t,400,c.Response().StatusCode())
	})
	t.Run("Error bad schema body, return 400",func(t *testing.T){
		// Arrenge
		fakeServiceUser.On("UpdateUser","test","test@gmail.com").Return(nil).Once()
		c.Request().Header.SetContentType("application/json")
        c.Request().SetBody([]byte(`{"email":"test"}`))
		c.Locals("user",token)
		// Act
		err := userController.UpdateUser(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t,400,c.Response().StatusCode())
	})
}

func TestDeleteUser(t *testing.T){
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)
	fakeServiceUser := &FakeUserService{}
	userController := infrastructure.NewUserController(fakeServiceUser)
	token := &jtoken.Token{
		Claims: jtoken.MapClaims{
			"username":"test",
		},
	}
	t.Run("User deleted sucessfully, return 200",func(t *testing.T){
		// Arrenge
		fakeServiceUser.On("DeleteUser","test").Return(nil).Once()
		c.Locals("user",token)
		// Act
		err := userController.DeleteUser(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t, `{"message":"Successful operation"}`, string(c.Response().Body()),"Must return status 200 with message")
		utils.AssertEqual(t,200,c.Response().StatusCode())
	})
	t.Run("User does not exists, return 400",func(t *testing.T){
		// Arrenge
		fakeServiceUser.On("DeleteUser","test").Return(errors.New("user does not exist")).Once()
		c.Locals("user",token)
		// Act
		err := userController.DeleteUser(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t, `{"message":"user does not exist"}`, string(c.Response().Body()),"Must return status 400 with message")
		utils.AssertEqual(t,400,c.Response().StatusCode())
	})
	t.Run("User has already deleted, return 400",func(t *testing.T){
		// Arrenge
		fakeServiceUser.On("DeleteUser","test").Return(errors.New("user has already deleted")).Once()
		c.Locals("user",token)
		// Act
		err := userController.DeleteUser(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t, `{"message":"user has already deleted"}`, string(c.Response().Body()),"Must return status 400 with message")
		utils.AssertEqual(t,400,c.Response().StatusCode())
	})
}

func TestGetUsersC(t *testing.T){
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)
	fakeServiceUser := &FakeUserService{}
	userController := infrastructure.NewUserController(fakeServiceUser)
	token := &jtoken.Token{
		Claims: jtoken.MapClaims{
			"username":"test",
		},
	}
	t.Run("Get list of user successfully, return 200",func(t *testing.T){
		// Arrenge
		c.Locals("user",token)
		usersExpected := &models.Users{
			{	
				ID: primitive.NewObjectID(),
				Username: "user1",
				Password: "password1",
				Email: "test1@test.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				ID: primitive.NewObjectID(),
				Username: "user2",
				Password: "password2",
				Email: "test2@test.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
		fakeServiceUser.On("GetUsers").Return(usersExpected).Once()
		// Act
		err := userController.GetUsers(c)
		usersJson, _ := json.Marshal(usersExpected)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t,200,c.Response().StatusCode())
		utils.AssertEqual(t, string(usersJson), string(c.Response().Body()),"Must return list of users found with status 200")
	})
}