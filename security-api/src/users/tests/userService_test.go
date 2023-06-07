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

func TestUserCreate(t *testing.T) {
	fakeRepo := &FakeUserRepository{}
	service := s.NewUserService(fakeRepo)
	t.Log("Create User")
	t.Run("Create user sucessfully",func(t *testing.T) {
		const username,password,email = "user","password","test@test.com"
		t.Log("Test if method create has been called, user created sucessfully")
		fakeRepo.On("Create", username, password, email)
		service.CreateUser(username,password,email)
	})
}