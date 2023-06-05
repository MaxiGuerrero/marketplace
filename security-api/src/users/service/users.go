package service

import repositories "marketplace/security-api/src/users/models" 

type UserService struct{
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) *UserService{
	return &UserService{
		userRepository,
	}
}

func (us UserService) CreateUser(){
	us.userRepository.Create()
}