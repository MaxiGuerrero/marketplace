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
	ur.Called(username,password,email)
	return nil
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
	t.Run("Create user sucessfully",func(t *testing.T) {
		// Arrenge
		fakeRepo := &FakeUserRepository{}
		fakeEncrpypter := &FakeEncrypter{}
		service := s.NewUserService(fakeRepo, fakeEncrpypter)
		const username,password,email = "user","password","test@test.com"
		passwordEncrypted := []byte(password)
		t.Log("Validate that GenerateHash has been called, retrun an mock password")
		fakeEncrpypter.On("GenerateHash",[]byte(password)).Return(passwordEncrypted,nil)
		t.Log("Validate that Create has been called, user created sucessfully, return nil")
		fakeRepo.On("Create", username, string(passwordEncrypted), email).Return(nil)
		// Act
		err := service.CreateUser(username,password,email)
		// Assert
		require.NoError(t,err,"Service does not return an error")
	})
	t.Run("Error on encrypting",func(t *testing.T) {
		// Arrenge
		fakeRepo := &FakeUserRepository{}
		fakeEncrpypter := &FakeEncrypter{}
		service := s.NewUserService(fakeRepo, fakeEncrpypter)
		const username,password,email = "user","password","test@test.com"
		t.Log("Validate that GenerateHash has been called, retrun error")
		fakeEncrpypter.On("GenerateHash",[]byte(password)).Return([]byte{},assert.AnError)
		fakeRepo.AssertNotCalled(t,"Create","Repository must not be called")
		// Act
		err := service.CreateUser(username,password,email)
		// Assert
		assert.Error(t,err,"Service return an error")
	})
}