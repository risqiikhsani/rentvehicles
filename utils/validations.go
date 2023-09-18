package utils

import (
	"github.com/go-playground/validator/v10"
)

// Custom validation function for password
func PasswordContainsAlphabetAndNumber(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Define regular expressions for alphabetic and numeric characters
	alphabetic := false
	numeric := false

	for _, char := range password {
		if 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' {
			alphabetic = true
		}
		if '0' <= char && char <= '9' {
			numeric = true
		}
	}

	return alphabetic && numeric
}

func InitializeValidator() {
	validate := validator.New()
	// Register custom validation tags or rules here
	validate.RegisterValidation("passwordContainsAlphabetAndNumber", PasswordContainsAlphabetAndNumber)
}
