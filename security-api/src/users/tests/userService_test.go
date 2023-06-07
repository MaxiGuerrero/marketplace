package service

import (
	s "marketplace/security-api/src/users/service"
	"testing"

	"github.com/stretchr/testify/mock"
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
	encrypter.Called(password)
	return nil,nil
}

func (encrypter FakeEncrypter) Compare(hashedPassword, password []byte) error{
	encrypter.Called(hashedPassword,password)
	return nil
}

func TestUserCreate(t *testing.T) {
	fakeRepo := &FakeUserRepository{}
	fakeEncrpypter := &FakeEncrypter{}
	service := s.NewUserService(fakeRepo, fakeEncrpypter)
	t.Log("Create User")
	t.Run("Create user sucessfully",func(t *testing.T) {
		const username,password,email = "user","password","test@test.com"
		t.Log("Test if method create has been called, user created sucessfully")
		fakeRepo.On("Create", username, password, email)
		service.CreateUser(username,password,email)
	})
}