package service

type AuthenticationService struct{}

func NewAuthenticationService() *AuthenticationService{
	return &AuthenticationService{}
}

func (as AuthenticationService) Login() string{
	return "Hello World"
}