package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct[T any](payload T) []string {
	var errors []string
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element string
			element = fmt.Sprintf("%s %s %s", err.Field(), err.Tag(), err.Kind())
			errors = append(errors, element)
		}
	}
	return errors
}
