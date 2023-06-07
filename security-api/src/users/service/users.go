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

func (us UserService) CreateUser(username,password,email string) error{
	return us.userRepository.Create(username,password,email)
}