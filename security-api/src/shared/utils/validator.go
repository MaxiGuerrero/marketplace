package utils

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func ValidateSchema(payload interface{}) error {
	err:= validate.Struct(payload)
	return err
}