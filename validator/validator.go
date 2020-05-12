package validator

import (
	"regexp"

	"gopkg.in/go-playground/validator.v9"
)

// validation  constants
const (
	IndiaPhoneNumberTag = "indiaPhoneNumber"
)

// IsIndiaPhoneNumber checks parameter pattern for valid phone number of India
func IsIndiaPhoneNumber(fl validator.FieldLevel) bool {
	IndiaPhoneNumberRegexString := "^[0-9]{10}$"
	IndiaPhoneNumberRegex := regexp.MustCompile(IndiaPhoneNumberRegexString)

	return IndiaPhoneNumberRegex.MatchString(fl.Field().String())
}

// New return new instance of Validate
func New() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation(IndiaPhoneNumberTag, IsIndiaPhoneNumber)
	return validate
}
