package tests

import (
	models "marketplace/security-api/src/users/models"
	s "marketplace/security-api/src/users/service"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FakeUserRepository struct {
	mock.Mock
}

func (ur *FakeUserRepository) Create(username,password,email string){
	ur.Called(username,password,email)
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
	const username,password,email = "user","password","test@test.com"
	passwordEncrypted := []byte(password)
	fakeRepo := &FakeUserRepository{}
	fakeEncrpypter := &FakeEncrypter{}
	service := s.NewUserService(fakeRepo, fakeEncrpypter)
	t.Log("Create User")
	t.Run("user created sucessfully",func(t *testing.T) {
		// Arrenge
		fakeEncrpypter.On("GenerateHash",[]byte(password)).Return(passwordEncrypted).Once()
		fakeRepo.On("Create", username, string(passwordEncrypted), email).Once()
		fakeRepo.On("GetByUsername", username).Return(nil).Once()
		// Act
		err := service.CreateUser(username,password,email)
		// Assert
		require.NoError(t,err,"Service does not return an error")
	})
	t.Run("user already exists, return error",func(t *testing.T) {
		// Arrenge
		userFound := new(models.User)
		fakeRepo.On("GetByUsername", username).Return(userFound).Once()
		// Act
		err := service.CreateUser(username,password,email)
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
		}
		fakeRepo.On("GetByUsername", userFoundDeleted.Username).Return(userFoundDeleted).Once()
		// Act
		err := service.DeleteUser(userFoundDeleted.Username)
		// Assert
		require.Error(t,err,"User does not exists")
	})
}