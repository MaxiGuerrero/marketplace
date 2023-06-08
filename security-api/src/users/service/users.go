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
	hashedPassword,err := us.encrypter.GenerateHash([]byte(password))
	if(err != nil){
		return err
	}
	return us.userRepository.Create(username,string(hashedPassword),email)
}