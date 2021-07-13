package util

import "github.com/go-playground/validator"

func ValidateModel(model interface{}) error {
	validate := validator.New()
	err := validate.Struct(model)

	if err != nil {

	}

	return nil
}
