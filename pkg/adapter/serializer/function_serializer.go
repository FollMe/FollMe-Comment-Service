package serializer

import (
	validator "github.com/go-playground/validator/v10"
)

func Validate(obj interface{}) error {
	validate := validator.New()
	return validate.Struct(obj)
}
