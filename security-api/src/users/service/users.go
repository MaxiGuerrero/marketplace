package service

import (
	"errors"
	models "marketplace/security-api/src/users/models"
)

// Userservice is a service that is responsable to manage users in the system.
// All service manage the business logical.
type UserService struct{
	userRepository models.IUserRepository
	encrypter models.IEncrypter
}

// Create an instance of the UserService with injection dependencies.
func NewUserService(userRepository models.IUserRepository, encrypter models.IEncrypter) *UserService{
	return &UserService{
		userRepository,
		encrypter,
	}
}

//  Create an user in the system, is possible return an business error if exists the user with the same username.
func (us UserService) CreateUser(username,password,email,role string) error{
	userFound := us.userRepository.GetByUsername(username)
	if userFound != nil {
		return errors.New("username already exists, please use another")
	}
	hashedPassword := us.encrypter.GenerateHash([]byte(password))
	us.userRepository.Create(username,string(hashedPassword),email,role)
	return nil
}

// Update an user in the system, an user can update its email. 
// Is possible return an business error if the user doesn't exists.
func (us UserService) UpdateUser(username string, email string) error{
	userFound := us.userRepository.GetByUsername(username)
	if userFound == nil {
		return errors.New("user does not exist")
	}
	if userFound.Status == models.Inactive.String() {
		return errors.New("user does not exist")
	}
	us.userRepository.Update(username,email)
	return nil
}

// Delete an user in the system, is a logical delete.
// Is possible return an business error if the user doesn't exists.
func (us UserService) DeleteUser(username string) error{
	userFound := us.userRepository.GetByUsername(username)
	if userFound == nil {
		return errors.New("user does not exist")
	}
	if userFound.Status == models.Inactive.String() {
		return errors.New("user has already deleted")
	}
	us.userRepository.Delete(username)
	return nil
}

// Get an array of users in the system.
func (us UserService) GetUsers() *models.Users {
	return us.userRepository.Get()
}