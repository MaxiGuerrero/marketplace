package tests

import (
	"marketplace/security-api/src/authentication/models"
	s "marketplace/security-api/src/authentication/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FakeEncrypter struct {
	mock.Mock
}

func (fe *FakeEncrypter) Compare(hashedPassword, password []byte) bool {
	args := fe.Called(hashedPassword,password)
	return args.Bool(0)
}

type FakeAuthRepository struct {
	mock.Mock
}

func (fu *FakeAuthRepository) GetByUsername(username string) *models.User {
	args := fu.Called(username)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*models.User)
}

type FakeJwtBuilder struct {
	mock.Mock
}

func (fj *FakeJwtBuilder) BuildToken(payload *models.Payload) string {
	args := fj.Called(payload)
	return args.String(0)
}

func (fj *FakeJwtBuilder) ValidateToken(token string) (*models.Payload,error) {
	args := fj.Called(token)
	return args.Get(0).(*models.Payload),args.Error(1)
}

func TestLoginService(t *testing.T) {
	fakeEncrypter := &FakeEncrypter{}
	fakeRepository := &FakeAuthRepository{}
	fakeJwtBuilder := &FakeJwtBuilder{}
	objectId := primitive.NewObjectID()
	time := time.Time{}
	t.Log("Login Test")
	service := s.NewAuthenticationService(fakeEncrypter,fakeRepository,fakeJwtBuilder)
	t.Run("Login user successfully, return token", func(t *testing.T) {
		// Arrenge
			user := &models.User{
				ID: objectId,
				Username: "test",
				Password: "Test1234",
				CreatedAt: time,
			}
			hashedPassword := []byte(user.Password)
			payload := &models.Payload{
				UserId: objectId,
				Username: user.Username,
				CreatedAt: time,
				UpdatedAt: time,
			}
			token := &models.UserToken{
				UserId: objectId,
				Token: "token",
			}

			fakeEncrypter.On("Compare",hashedPassword,[]byte(user.Password)).Return(true).Once()
			fakeRepository.On("GetByUsername",user.Username).Return(user).Once()
			fakeJwtBuilder.On("BuildToken",payload).Return(token.Token).Once()
		// Act
			response,err := service.Login(user.Username,user.Password)
		// Assert
			require.NoError(t,err,"Login does not return an error")
			require.Equal(t,token,response,"Service return token")
	})
	t.Run("User doesnot exists, return error 'username or password is incorrect, try again'", func(t *testing.T) {
		// Arrenge
			user := &models.User{
				ID: objectId,
				Username: "test",
				Password: "Test1234",
				CreatedAt: time,
			}
			fakeRepository.On("GetByUsername",user.Username).Return(nil).Once()
		// Act
			_,err := service.Login(user.Username,user.Password)
		// Assert
			require.Error(t,err,"Login does return an error")
			assert.Equal(t,"username or password is incorrect, try again",err.Error())
	})
	t.Run("User exists but password not match, return error 'username or password is incorrect, try again'", func(t *testing.T) {
		// Arrenge
			user := &models.User{
				ID: objectId,
				Username: "test",
				Password: "Test1234",
				CreatedAt: time,
			}
			hashedPassword := []byte(user.Password)

			fakeEncrypter.On("Compare",hashedPassword,[]byte(user.Password)).Return(false).Once()
			fakeRepository.On("GetByUsername",user.Username).Return(user).Once()
		// Act
			_,err := service.Login(user.Username,user.Password)
		// Assert
			require.Error(t,err,"Login does return an error")
			assert.Equal(t,"username or password is incorrect, try again",err.Error())
	})
}