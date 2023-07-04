package utils

import (
	"log"
	models "marketplace/security-api/src/users/models"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func RegisterValidation(){
	// validates that an enum is within the interval
	err := validate.RegisterValidation("role_enum_validation", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		switch value {
			case models.ADMIN.String(), models.USER.String():
				return true
			}
		return false
	})
	if err != nil {
		log.Panicf("Error on register validation: %v", err)
	}
}

func ValidateSchema(payload interface{}) error {
	err:= validate.Struct(payload)
	return err
}