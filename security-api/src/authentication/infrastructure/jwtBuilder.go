package infrastructure

import (
	"errors"
	"log"
	"marketplace/security-api/src/authentication/models"
	config "marketplace/security-api/src/shared"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Responsable to implement the JWT token.
type JWTBuilder struct {}

// Build token of a exists user.
func (j JWTBuilder) BuildToken(payload *models.Payload) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": payload.UserId,
		"username": payload.Username,
		"createdAt": payload.CreatedAt,
		"updatedAt": payload.UpdatedAt,
		"deletedAt": payload.DeletedAt,
		"role": payload.Role,
	})
	tokenString, err := token.SignedString(config.GetConfig().JWTSecret)
	if err != nil {
		log.Panicf("Error to generate token: %v", err.Error())
	}
	return tokenString
}

// Validate if a token is correct.
func (j JWTBuilder) ValidateToken(tokenString string) (*models.Payload,error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return config.GetConfig().JWTSecret, nil
	})
	if err != nil {
		return nil,err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, errors.New("")
	}
	payload := &models.Payload{
		UserId: claims["userId"].(primitive.ObjectID),
		Username: claims["username"].(string),
		CreatedAt: claims["createdAt"].(time.Time),
		UpdatedAt: claims["updatedAt"].(time.Time),
		DeletedAt: claims["deletedAt"].(time.Time),
		Role: claims["role"].(string),
	}
	return payload, nil
}