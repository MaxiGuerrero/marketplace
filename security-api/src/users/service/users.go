package service

import models "marketplace/security-api/src/users/models" 

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
	return us.userRepository.Create(username,password,email)
}