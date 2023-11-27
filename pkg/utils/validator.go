package utils

import "github.com/go-playground/validator/v10"

func Validator(data interface{}) error {
	validte := validator.New()

	err := validte.Struct(data)

	if err != nil {
		return err
	}

	return nil
}
