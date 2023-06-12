package tests

import (
	models "marketplace/security-api/src/users/models"
	s "marketplace/security-api/src/users/service"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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