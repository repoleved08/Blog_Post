package validators

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(input interface{}) error {
	return validate.Struct(input)
}

func FormatValidationError(err error) string {
	var errors []string
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, err := range validationErrors {
			switch err.Tag() {
			case "required":
				errors = append(errors, fmt.Sprintf("%s is required", err.Field()))
			case "min":
				errors = append(errors, fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param()))
			case "max":
				errors = append(errors, fmt.Sprintf("%s must be at most %s characters long", err.Field(), err.Param()))
			case "email":
				errors = append(errors, fmt.Sprintf("%s must be a valid email address", err.Field()))
			// You can add more cases for other tags (e.g., "len", "gte", etc.)
			default:
				errors = append(errors, fmt.Sprintf("%s is invalid", err.Field()))
			}
		}
	}
	return strings.Join(errors, ", ")
}
