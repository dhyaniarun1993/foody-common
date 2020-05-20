package validator

import (
	"regexp"

	"gopkg.in/go-playground/validator.v9"
)

// validation  constants
const (
	IndiaPhoneNumberTag = "indiaPhoneNumber"
	OtpTag              = "otp"
)

// IsIndiaPhoneNumber checks parameter pattern for valid phone number of India
func IsIndiaPhoneNumber(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}
	IndiaPhoneNumberRegexString := "^[0-9]{10}$"
	IndiaPhoneNumberRegex := regexp.MustCompile(IndiaPhoneNumberRegexString)

	return IndiaPhoneNumberRegex.MatchString(fl.Field().String())
}

// IsOtp checks parameter pattern for valid otp of 6 digital
func IsOtp(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}
	OtpRegexString := "^[0-9]{6}$"
	OtpRegex := regexp.MustCompile(OtpRegexString)

	return OtpRegex.MatchString(fl.Field().String())
}

// New return new instance of Validate
func New() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation(IndiaPhoneNumberTag, IsIndiaPhoneNumber)
	validate.RegisterValidation(OtpTag, IsOtp)
	return validate
}
