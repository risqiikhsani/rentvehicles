// utils/validation.go
package utils

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
)

var En ut.Translator
var universalTraslator *ut.UniversalTranslator

func InitializeTranslator() {
	english := en.New()
	universalTraslator = ut.New(english, english)
	En, _ = universalTraslator.GetTranslator("en")
}
