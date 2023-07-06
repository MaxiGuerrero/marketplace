package models

// Struct that represent a request to create an user.
type CreateUserRequest struct {
	Username  string `json:"username" validate:"required,min=4,max=32"`
	Password string `json:"password" validate:"required,min=4,max=32"`
	Email string `json:"email" validate:"required,email,min=6,max=32"`
	Role  string `json:"role" validate:"required,role_enum_validation"`
}

// Struct that represent a request to update an user.
type UpdateUserRequest struct {
	Email string `json:"email" validate:"required,email,min=6,max=32"` 
}