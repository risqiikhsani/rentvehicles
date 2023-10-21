package utils

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var Validate *validator.Validate

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
	Validate = validator.New(validator.WithRequiredStructEnabled())
	// Register custom validation tags or rules here
	// Validate.RegisterValidation("passwordContainsAlphabetAndNumber", PasswordContainsAlphabetAndNumber)
	_ = en_translations.RegisterDefaultTranslations(Validate, En)
}

func TranslateError(err error, trans ut.Translator) (errs []string) {
	if err == nil {
		return nil
	}

	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := e.Translate(trans)
		errs = append(errs, translatedErr)
	}

	return errs
}
