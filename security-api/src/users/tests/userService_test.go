package service

import (
	models "marketplace/security-api/src/users/models"
	s "marketplace/security-api/src/users/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type FakeUserRepository struct {
	mock.Mock
}

func (ur *FakeUserRepository) Create(username,password,email string) error{
	args := ur.Called(username,password,email)
	return args.Error(0)
}

func (ur *FakeUserRepository) GetByUsername(username string) (*models.User,error){
	args := ur.Called(username)
	if args.Get(0) == nil {
		return nil,args.Error(1)
	}
	return args.Get(0).(*models.User),args.Error(1)
}

type FakeEncrypter struct {
	mock.Mock
}

func (encrypter *FakeEncrypter) GenerateHash(password []byte) ([]byte, error){
	args := encrypter.Called(password)
	return args.Get(0).([]byte), args.Error(1)
}

func (encrypter *FakeEncrypter) Compare(hashedPassword, password []byte) error{
	args := encrypter.Called(hashedPassword,password)
	return args.Error(0)
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
		fakeEncrpypter.On("GenerateHash",[]byte(password)).Return(passwordEncrypted,nil).Once()
		fakeRepo.On("Create", username, string(passwordEncrypted), email).Return(nil).Once()
		fakeRepo.On("GetByUsername", username).Return(nil,nil).Once()
		// Act
		err := service.CreateUser(username,password,email)
		// Assert
		require.NoError(t,err,"Service does not return an error")
	})
	t.Run("user already exists, return error",func(t *testing.T) {
		// Arrenge
		userFound := new(models.User)
		fakeRepo.On("GetByUsername", username).Return(userFound,nil).Once()
		// Act
		err := service.CreateUser(username,password,email)
		// Assert
		require.Error(t,err,"Username already exists")
	})
	t.Run("Error on encrypting, return error",func(t *testing.T) {
		// Arrenge
		fakeEncrpypter.On("GenerateHash",[]byte(password)).Return([]byte{},assert.AnError).Once()
		fakeRepo.On("GetByUsername", username).Return(nil,nil).Once()
		// Act
		err := service.CreateUser(username,password,email)
		// Assert
		fakeRepo.AssertNotCalled(t,"Create","Repository must not be called")
		assert.Error(t,err,"Service return an error")
	})
	t.Run("Error on repository, return error",func(t *testing.T) {
		// Arrenge
		fakeEncrpypter.On("GenerateHash",[]byte(password)).Return(passwordEncrypted,nil).Once()
		fakeRepo.On("Create", username, string(passwordEncrypted), email).Return(assert.AnError).Once()
		fakeRepo.On("GetByUsername", username).Return(nil,nil).Once()
		// Act
		err := service.CreateUser(username,password,email)
		// Assert
		assert.Error(t,err,"Service return an error")
	})
}