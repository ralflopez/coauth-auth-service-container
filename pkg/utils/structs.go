package utils

import (
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(s interface{}) error {
	return validator.New().Struct(s)
}
