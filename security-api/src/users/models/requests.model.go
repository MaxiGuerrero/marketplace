package models

type CreateUserRequest struct {
	Username  string `json:"username" validate:"required,min=4,max=32"`
	Password string `json:"password" validate:"required,min=4,max=32"`
	Email string `json:"email" validate:"required,email,min=6,max=32"` 
}

type UpdateUserRequest struct {
	Email string `json:"email" validate:"required,email,min=6,max=32"` 
}