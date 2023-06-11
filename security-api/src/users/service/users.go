package service

import (
	"errors"
	models "marketplace/security-api/src/users/models"
) 

type UserService struct{
	userRepository models.IUserRepository
	encrypter models.IEncrypter
}

func NewUserService(userRepository models.IUserRepository, encrypter models.IEncrypter) *UserService{
	return &UserService{
		userRepository,
		encrypter,
	}
}

func (us UserService) CreateUser(username,password,email string) error{
	userFound := us.userRepository.GetByUsername(username)
	if userFound != nil {
		return errors.New("username already exists, please use another")
	}
	hashedPassword := us.encrypter.GenerateHash([]byte(password))
	us.userRepository.Create(username,string(hashedPassword),email)
	return nil
}