package utils

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Validate if a schema from a request is correct.
func ValidateSchema(payload interface{}) error {
	err:= validate.Struct(payload)
	return err
}