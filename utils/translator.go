// utils/validation.go
package utils

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
)

var En ut.Translator

// only one instance as translators within are shared + goroutine safe
var universalTraslator *ut.UniversalTranslator

func InitializeTranslator() {
	e := en.New()
	universalTraslator = ut.New(e, e)
	En, _ = universalTraslator.GetTranslator("en")
}
