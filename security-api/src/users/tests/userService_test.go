package tests

import (
	"marketplace/security-api/src/shared/utils"
	models "marketplace/security-api/src/users/models"
	s "marketplace/security-api/src/users/service"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMain(m *testing.M) {
	utils.RegisterValidation()
	os.Exit(m.Run())
}

type FakeUserRepository struct {
	mock.Mock
}

func (ur *FakeUserRepository) Create(username,password,email,role string){
	ur.Called(username,password,email,role)
}

func (ur *FakeUserRepository) GetByUsername(username string) *models.User{
	args := ur.Called(username)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*models.User)
}

func (ur *FakeUserRepository) Update(username string, email string){
	ur.Called(username,email)
}

func (ur *FakeUserRepository) Delete(username string){
	ur.Called(username)
}

func (ur *FakeUserRepository) Get() *models.Users{
	args := ur.Called()
	return args.Get(0).(*models.Users)
}

type FakeEncrypter struct {
	mock.Mock
}

func (encrypter *FakeEncrypter) GenerateHash(password []byte) []byte{
	args := encrypter.Called(password)
	return args.Get(0).([]byte)
}

func (encrypter *FakeEncrypter) Compare(hashedPassword, password []byte) bool{
	args := encrypter.Called(hashedPassword,password)
	return args.Get(0).(bool)
}

func TestUserCreate(t *testing.T) {
	const username,password,email,role = "user","password","test@test.com","USER"
	passwordEncrypted := []byte(password)
	fakeRepo := &FakeUserRepository{}
	fakeEncrpypter := &FakeEncrypter{}
	service := s.NewUserService(fakeRepo, fakeEncrpypter)
	t.Log("Create User")
	t.Run("user created sucessfully",func(t *testing.T) {
		// Arrenge
		fakeEncrpypter.On("GenerateHash",[]byte(password)).Return(passwordEncrypted).Once()
		fakeRepo.On("Create", username, string(passwordEncrypted), email,role).Once()
		fakeRepo.On("GetByUsername", username).Return(nil).Once()
		// Act
		err := service.CreateUser(username,password,email,role)
		// Assert
		require.NoError(t,err,"Service does not return an error")
	})
	t.Run("user already exists, return error",func(t *testing.T) {
		// Arrenge
		userFound := new(models.User)
		fakeRepo.On("GetByUsername", username).Return(userFound).Once()
		// Act
		err := service.CreateUser(username,password,email,role)
		// Assert
		require.Error(t,err,"Username already exists")
	})
}

func TestUserUpdate(t *testing.T) {
	userFound := &models.User{
		ID: primitive.NewObjectID(),
		Username: "user",
		Password: "password",
		Email: "test@test.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status: models.Active.String(),
		Role: models.USER.String(),
	}
	fakeRepo := &FakeUserRepository{}
	fakeEncrpypter := &FakeEncrypter{}
	service := s.NewUserService(fakeRepo, fakeEncrpypter)
	t.Log("Update User")
	t.Run("user updated sucessfully",func(t *testing.T) {
		// Arrenge
		fakeRepo.On("Update", userFound.Username, userFound.Email).Once()
		fakeRepo.On("GetByUsername", userFound.Username).Return(userFound).Once()
		// Act
		err := service.UpdateUser(userFound.Username,userFound.Email)
		// Assert
		require.NoError(t,err,"Service does not return an error")
	})
	t.Run("user does not exists, return error",func(t *testing.T) {
		// Arrenge
		fakeRepo.On("GetByUsername", "userNotExists").Return(nil).Once()
		// Act
		err := service.UpdateUser("userNotExists","email@email.com")
		// Assert
		require.Error(t,err,"User does not exists")
	})
}

func TestUserDelete(t *testing.T) {
	userFound := &models.User{
		ID: primitive.NewObjectID(),
		Username: "user",
		Password: "password",
		Email: "test@test.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status: models.Active.String(),
		Role: models.USER.String(),
	}
	fakeRepo := &FakeUserRepository{}
	fakeEncrpypter := &FakeEncrypter{}
	service := s.NewUserService(fakeRepo, fakeEncrpypter)
	t.Log("Delete User")
	t.Run("user deleted sucessfully",func(t *testing.T) {
		// Arrenge
		fakeRepo.On("Delete", userFound.Username).Once()
		fakeRepo.On("GetByUsername", userFound.Username).Return(userFound).Once()
		// Act
		err := service.DeleteUser(userFound.Username)
		// Assert
		require.NoError(t,err,"Service does not return an error")
	})
	t.Run("user does not exists, return error",func(t *testing.T) {
		// Arrenge
		fakeRepo.On("GetByUsername", "userNotExists").Return(nil).Once()
		// Act
		err := service.DeleteUser("userNotExists")
		// Assert
		require.Error(t,err,"User does not exists")
	})

	t.Run("user has already deleted, return error",func(t *testing.T) {
		// Arrenge
		userFoundDeleted := &models.User{
			ID: primitive.NewObjectID(),
			Username: "user",
			Password: "password",
			Email: "test@test.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: time.Now(),
			Status: models.Inactive.String(),
			Role: models.USER.String(),
		}
		fakeRepo.On("GetByUsername", userFoundDeleted.Username).Return(userFoundDeleted).Once()
		// Act
		err := service.DeleteUser(userFoundDeleted.Username)
		// Assert
		require.Error(t,err,"User does not exists")
	})
}

func TestGetUsers(t *testing.T) {
	fakeRepo := &FakeUserRepository{}
	fakeEncrpypter := &FakeEncrypter{}
	service := s.NewUserService(fakeRepo, fakeEncrpypter)
	t.Log("Get Users")
	t.Run("Get list of users found",func(t *testing.T) {
		// Arrenge
		usersExpected := &models.Users{
			{	
				ID: primitive.NewObjectID(),
				Username: "user1",
				Password: "password1",
				Email: "test1@test.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Role: models.USER.String(),
				Status: models.Active.String(),
			},
			{
				ID: primitive.NewObjectID(),
				Username: "user2",
				Password: "password2",
				Email: "test2@test.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Role: models.USER.String(),
				Status: models.Active.String(),
			},
		}
		fakeRepo.On("Get").Return(usersExpected).Once()
		// Act
		users := service.GetUsers()
		// Assert
		require.Equal(t, usersExpected, users, "List of users must be equal")
	})
}