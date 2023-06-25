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

func (us UserService) UpdateUser(username string, email string) error{
	userFound := us.userRepository.GetByUsername(username)
	if userFound == nil {
		return errors.New("user does not exist")
	}
	us.userRepository.Update(username,email)
	return nil
}

func (us UserService) DeleteUser(username string) error{
	userFound := us.userRepository.GetByUsername(username)
	if userFound == nil {
		return errors.New("user does not exist")
	}
	if !userFound.DeletedAt.IsZero() {
		return errors.New("user has already deleted")
	}
	us.userRepository.Delete(username)
	return nil
}