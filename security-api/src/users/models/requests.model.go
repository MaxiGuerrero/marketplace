package models

type CreateUserRequest struct {
	Username  string `json:"username" validate:"required,min=4,max=32"`
	Password string `json:"password" validate:"required,min=4,max=32"`
	Email string `json:"email" validate:"required,email,min=6,max=32"`
	Role  string `json:"role" validate:"required,role_enum_validation"`
}

type UpdateUserRequest struct {
	Email string `json:"email" validate:"required,email,min=6,max=32"` 
}