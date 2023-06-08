package service

import (
	s "marketplace/security-api/src/users/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type FakeUserRepository struct {
	mock.Mock
}

func (ur FakeUserRepository) Create(username,password,email string) error{
	args := ur.Called(username,password,email)
	return args.Error(0)
}

type FakeEncrypter struct {
	mock.Mock
}

func (encrypter FakeEncrypter) GenerateHash(password []byte) ([]byte, error){
	args := encrypter.Called(password)
	return args.Get(0).([]byte), args.Error(1)
}

func (encrypter FakeEncrypter) Compare(hashedPassword, password []byte) error{
	args := encrypter.Called(hashedPassword,password)
	return args.Error(0)
}

func TestUserCreate(t *testing.T) {
	t.Log("Create User")
	t.Run("user created sucessfully",func(t *testing.T) {
		// Arrenge
		const username,password,email = "user","password","test@test.com"
		fakeRepo := &FakeUserRepository{}
		fakeEncrpypter := &FakeEncrypter{}
		service := s.NewUserService(fakeRepo, fakeEncrpypter)
		passwordEncrypted := []byte(password)
		fakeEncrpypter.On("GenerateHash",[]byte(password)).Return(passwordEncrypted,nil)
		fakeRepo.On("Create", username, string(passwordEncrypted), email).Return(nil)
		// Act
		err := service.CreateUser(username,password,email)
		// Assert
		require.NoError(t,err,"Service does not return an error")
	})
	t.Run("Error on encrypting, return error",func(t *testing.T) {
		// Arrenge
		const username,password,email = "user","password","test@test.com"
		fakeRepo := &FakeUserRepository{}
		fakeEncrpypter := &FakeEncrypter{}
		service := s.NewUserService(fakeRepo, fakeEncrpypter)
		fakeEncrpypter.On("GenerateHash",[]byte(password)).Return([]byte{},assert.AnError)
		fakeRepo.AssertNotCalled(t,"Create","Repository must not be called")
		// Act
		err := service.CreateUser(username,password,email)
		// Assert
		assert.Error(t,err,"Service return an error")
	})
	t.Run("Error on repository, return error",func(t *testing.T) {
		// Arrenge
		const username,password,email = "user","password","test@test.com"
		fakeRepo := &FakeUserRepository{}
		fakeEncrpypter := &FakeEncrypter{}
		service := s.NewUserService(fakeRepo, fakeEncrpypter)
		passwordEncrypted := []byte(password)
		fakeEncrpypter.On("GenerateHash",[]byte(password)).Return(passwordEncrypted,nil)
		fakeRepo.On("Create", username, string(passwordEncrypted), email).Return(assert.AnError)
		// Act
		err := service.CreateUser(username,password,email)
		// Assert
		assert.Error(t,err,"Service return an error")
	})
}