package models

// Struct that represent a request to login an user.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}