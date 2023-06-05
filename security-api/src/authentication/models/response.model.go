package models

type LoginResponse struct {
	IdUser string `json:"idUser"`
	Token string `json:"token"`
}