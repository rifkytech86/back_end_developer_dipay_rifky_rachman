package validations

import (
	"fmt"
	"github.com/dipay/internal"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

const (
	ValidPassword      = "validatorPassword"
	ValidUsername      = "validatorUserName"
	ValidAddress       = "validatorAddress"
	ValidCompanyName   = "validatorCompanyName"
	ValidPhoneNumber   = "validatorTelephoneNumber"
	ValidEmail         = "validatorEmail"
	ValidJobTitle      = "validatorJobTitle"
	ValidEmployee      = "validatorEmployee"
	ValidDuplicateZero = "validatorDuplicateZero"
)

//go:generate mockery --name IValidator
type IValidator interface {
	Struct(s interface{}) []ValidationError
	RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error
}

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// MyValidator is an implementation of the Validator interface that wraps the validations.Validate
type customValidator struct {
	validate *validator.Validate
}

func (mv *customValidator) Struct(s interface{}) []ValidationError {
	var validationErrors []ValidationError
	err := mv.validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ValidationError{
				Field: err.Field(),
				Error: err.Tag(),
			})
		}
	}

	return validationErrors
}

func (mv *customValidator) RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	return mv.validate.RegisterValidation(tag, fn, callValidationEvenIfNull...)
}

func NewCustomValidator() IValidator {
	return &customValidator{
		validate: validator.New(),
	}
}

func ValidatorUsername(fl validator.FieldLevel) bool {
	userName := fl.Field().String()
	if userName == "" {
		return false
	}
	if len(userName) > 30 {
		return false
	}

	return true
}

func ValidatorAddress(fl validator.FieldLevel) bool {
	address := fl.Field().String()
	if address == "" {
		return false
	}

	if len(address) < 10 || len(address) > 50 {
		return false
	}
	return true
}

func ValidatorCompanyName(fl validator.FieldLevel) bool {
	companyName := fl.Field().String()
	if companyName == "" {
		return false
	}
	if len(companyName) < 3 || len(companyName) > 50 {
		return false
	}
	return true
}
func ValidatePhoneNumber(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()
	if phoneNumber == "" {
		return false
	}

	if len(phoneNumber) < 2 {
		return false
	}
	minChar := 8
	maxChar := 16
	if len(phoneNumber) > 3 && string(phoneNumber[0:3]) == "+62" {
		minChar = 10
		maxChar = 18
	}
	if len(phoneNumber) < minChar || len(phoneNumber) > maxChar {
		return false
	}

	return isValidPhoneNumber(phoneNumber)
}

func isValidPhoneNumber(input string) bool {
	phoneNumberRegex := regexp.MustCompile(`^(\+62|0)\d+$`)
	return phoneNumberRegex.MatchString(input)
}

func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	if email == "" {
		return false
	}
	if len(email) < 5 || len(email) > 255 {
		return false
	}
	return isValidEmail(email)

}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func ValidateJobTitle(fl validator.FieldLevel) bool {
	jobTitle := fl.Field().String()
	if jobTitle == "" {
		return false
	}
	convJobTitle := strings.ToLower(jobTitle)
	statusJobTitle := internal.JobTittle(convJobTitle)
	switch statusJobTitle {
	case internal.JobTitleManager:
		return true
	case internal.JobTitleDirector:
		return true
	case internal.JobTitleStaff:
		return true
	default:
		return false
	}

	return true
}
func ValidateEmployee(fl validator.FieldLevel) bool {
	employee := fl.Field().String()
	if employee == "" {
		return false
	}

	return true
}

type Input struct {
	N []int `json:"n"`
}

func ValidateDuplicateZero(fl validator.FieldLevel) bool {
	n := fl.Field().String()
	if n == "" {
		return false // Empty string is considered valid
	}

	return true
}

func ValidatorPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if password == "" {
		return false
	}

	if len(password) > 30 {
		return false
	}

	return true
}

func GetCustomMessage(msgError string, field string) string {
	if msgError == ValidUsername {
		customErrorMessage := fmt.Sprintf("%s %s", internal.ErrorInvalidRequestUserName.String(), field)
		return customErrorMessage
	}

	if msgError == ValidPassword {
		customErrorMessage := fmt.Sprintf("%s %s", internal.ErrorInvalidRequestPassword.String(), field)
		return customErrorMessage
	}
	if msgError == ValidCompanyName {
		customErrorMessage := fmt.Sprintf("%s %s", internal.ErrorInvalidRequestCompanyName.String(), field)
		return customErrorMessage
	}

	if msgError == ValidPhoneNumber {
		customErrorMessage := fmt.Sprintf("%s %s", internal.ErrorInvalidRequestPhoneNumber.String(), field)
		return customErrorMessage
	}

	if msgError == ValidAddress {
		customErrorMessage := fmt.Sprintf("%s %s", internal.ErrorInvalidRequestAddress.String(), field)
		return customErrorMessage
	}

	if msgError == ValidJobTitle {
		customErrorMessage := fmt.Sprintf("%s %s", internal.ErrorInvalidRequestJobTitle.String(), field)
		return customErrorMessage
	}
	if msgError == ValidEmail {
		customErrorMessage := fmt.Sprintf("%s %s", internal.ErrorInvalidRequestEmail.String(), field)
		return customErrorMessage
	}

	return fmt.Sprintf("%s %s", internal.ErrorInvalidRequest.String(), field)

}
