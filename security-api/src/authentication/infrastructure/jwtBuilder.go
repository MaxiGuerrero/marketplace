package infrastructure

import (
	"log"
	"marketplace/security-api/src/authentication/models"
	config "marketplace/security-api/src/shared"

	"github.com/golang-jwt/jwt/v5"
)

type JWTBuilder struct {}

func (j JWTBuilder) BuildToken(payload *models.Payload) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": payload.UserId,
		"username": payload.Username,
		"createdAt": payload.CreatedAt,
		"updatedAt": payload.UpdatedAt,
		"deletedAt": payload.DeletedAt,
	})
	tokenString, err := token.SignedString(config.GetConfig().JWTSecret)
	if err != nil {
		log.Panicf("Error to generate token: %v", err.Error())
	}
	return tokenString
}