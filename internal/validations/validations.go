package validations

import (
	"fmt"
	"github.com/dipay/internal"
	"github.com/go-playground/validator/v10"
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

	return true
}

func ValidatorAddress(fl validator.FieldLevel) bool {
	address := fl.Field().String()
	if address == "" {
		return false
	}

	return true
}

func ValidatorCompanyName(fl validator.FieldLevel) bool {
	address := fl.Field().String()
	if address == "" {
		return false
	}

	return true
}
func ValidatePhoneNumber(fl validator.FieldLevel) bool {
	address := fl.Field().String()
	if address == "" {
		return false
	}

	return true
}

func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	if email == "" {
		return false
	}

	return true
}

func ValidateJobTitle(fl validator.FieldLevel) bool {
	jobTitle := fl.Field().String()
	if jobTitle == "" {
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

	return true
}

func GetCustomMessage(msgError string, field string) string {
	if msgError == ValidUsername {
		customErrorMessage := fmt.Sprintf(internal.ErrorInvalidRequestUserName.String(), field)
		return customErrorMessage
	}

	if msgError == ValidPassword {
		customErrorMessage := fmt.Sprintf(internal.ErrorInvalidRequestPassword.String(), field)
		return customErrorMessage
	}

	return fmt.Sprintf(internal.ErrorInvalidRequest.String(), field, msgError)

}
