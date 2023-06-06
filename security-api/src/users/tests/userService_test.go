package service

import (
	s "marketplace/security-api/src/users/service"
	"testing"

	"github.com/stretchr/testify/mock"
)

type FakeUserRepository struct {
	mock.Mock
}

func (ur FakeUserRepository) Create(username,password,email string){
	ur.Called(username,password,email)
}

func TestUserCreate(t *testing.T) {
	fakeRepo := &FakeUserRepository{}
	service := s.NewUserService(fakeRepo)
	t.Log("Create User")
	t.Run("Create user sucessfully",func(t *testing.T) {
		const username,password,email = "user","password","test@test.com"
		t.Log("Create user sucessfully")
		fakeRepo.On("Create", username, password, email).Return()
		service.CreateUser(username,password,email)
	})
}