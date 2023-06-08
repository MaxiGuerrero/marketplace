package service

import (
	s "marketplace/security-api/src/users/service"
	"testing"

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
	fakeRepo := &FakeUserRepository{}
	fakeEncrpypter := &FakeEncrypter{}
	service := s.NewUserService(fakeRepo, fakeEncrpypter)
	t.Log("Create User")
	t.Run("Create user sucessfully",func(t *testing.T) {
		const username,password,email = "user","password","test@test.com"
		passwordEncrypted := []byte(password)
		t.Log("Test if method GenerateHash has been called, retrun an mock password")
		fakeEncrpypter.On("GenerateHash",[]byte(password)).Return(passwordEncrypted,nil)
		t.Log("Test if method create has been called, user created sucessfully, return nil")
		fakeRepo.On("Create", username, string(passwordEncrypted), email).Return(nil)
		err := service.CreateUser(username,password,email)
		require.Nil(t,err,"Service does not return an error")
	})
}