package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockValidator struct {
}

func (mv *MockValidator) Struct(s interface{}) []ValidationError {
	return []ValidationError{{"Field1", "Error1"}, {"Field2", "Error2"}}
}

func (mv *MockValidator) RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	return nil
}

func TestCustomValidator_Struct(t *testing.T) {
	mockValidator := &MockValidator{}
	customValidator := NewCustomValidator()

	result := customValidator.Struct(mockValidator)

	var expectedErrors []ValidationError = nil

	assert.Equal(t, expectedErrors, result)
}

func TestValidateDuplicateZero(t *testing.T) {
	validate := validator.New()

	// Define a struct with the validation tag
	input := Input{
		N: []int{0, 0, 1},
	}

	validate.RegisterValidation("ValidateDuplicateZero", ValidateDuplicateZero)

	err := validate.Struct(input)
	assert.NoError(t, err)

	inputEmpty := Input{
		N: []int{},
	}

	errEmpty := validate.Struct(inputEmpty)
	assert.NoError(t, errEmpty)

	inputSingleZero := Input{
		N: []int{0},
	}

	errSingleZero := validate.Struct(inputSingleZero)
	assert.NoError(t, errSingleZero)
}

func TestValidatorUsername(t *testing.T) {
	validate := validator.New()

	type User struct {
		Username string `validate:"ValidatorUsername"`
	}

	validate.RegisterValidation("ValidatorUsername", ValidatorUsername)

	userEmpty := User{
		Username: "",
	}

	errEmpty := validate.Struct(userEmpty)
	assert.Error(t, errEmpty)

	userNotEmpty := User{
		Username: "john_doe",
	}

	errNotEmpty := validate.Struct(userNotEmpty)
	assert.NoError(t, errNotEmpty)

	userEmptyMaxLimit := User{
		Username: "companycompanycompanycompanycompanycompanycompanycompanycompanycompany",
	}

	errEmptyMaxLimit := validate.Struct(userEmptyMaxLimit)
	assert.Error(t, errEmptyMaxLimit)

}

func TestValidatorAddress(t *testing.T) {
	validate := validator.New()

	type User struct {
		Address string `validate:"ValidatorAddress"`
	}

	validate.RegisterValidation("ValidatorAddress", ValidatorAddress)

	userEmpty := User{
		Address: "",
	}
	errEmpty := validate.Struct(userEmpty)
	assert.Error(t, errEmpty)
	userNotEmpty := User{
		Address: "123 Main St",
	}
	errNotEmpty := validate.Struct(userNotEmpty)
	assert.NoError(t, errNotEmpty)

	userLessLimit := User{
		Address: "123",
	}
	errLessLimit := validate.Struct(userLessLimit)
	assert.Error(t, errLessLimit)
}

func TestValidatorCompanyName(t *testing.T) {
	validate := validator.New()
	type Company struct {
		Name string `validate:"ValidatorCompanyName"`
	}
	validate.RegisterValidation("ValidatorCompanyName", ValidatorCompanyName)

	companyEmpty := Company{
		Name: "",
	}
	errEmpty := validate.Struct(companyEmpty)
	assert.Error(t, errEmpty)
	companyNotEmpty := Company{
		Name: "ABC Inc.",
	}
	errNotEmpty := validate.Struct(companyNotEmpty)
	assert.NoError(t, errNotEmpty)

	companyLessLimit := Company{
		Name: "23",
	}
	errLessLimit := validate.Struct(companyLessLimit)
	assert.Error(t, errLessLimit)
}

func TestValidatePhoneNumber(t *testing.T) {
	validate := validator.New()
	type Phone struct {
		Number string `validate:"ValidatePhoneNumber"`
	}
	validate.RegisterValidation("ValidatePhoneNumber", ValidatePhoneNumber)
	phoneEmpty := Phone{
		Number: "",
	}
	errEmpty := validate.Struct(phoneEmpty)
	assert.Error(t, errEmpty)
	phoneEmpty1 := Phone{
		Number: "2",
	}
	errEmpty1 := validate.Struct(phoneEmpty1)
	assert.Error(t, errEmpty1)

	phoneEmpty2 := Phone{
		Number: "+62123456",
	}
	errEmpty2 := validate.Struct(phoneEmpty2)
	assert.Error(t, errEmpty2)

	phoneNotEmpty := Phone{
		Number: "+621234567890",
	}
	errNotEmpty := validate.Struct(phoneNotEmpty)
	assert.NoError(t, errNotEmpty)

}

func TestValidateEmail(t *testing.T) {
	validate := validator.New()
	type Email struct {
		Address string `validate:"ValidateEmail"`
	}
	validate.RegisterValidation("ValidateEmail", ValidateEmail)
	emailEmpty := Email{
		Address: "",
	}
	errEmpty := validate.Struct(emailEmpty)
	assert.Error(t, errEmpty)
	emailNotEmpty := Email{
		Address: "test@example.com",
	}
	errNotEmpty := validate.Struct(emailNotEmpty)
	assert.NoError(t, errNotEmpty)
}

func TestValidateJobTitle(t *testing.T) {
	validate := validator.New()
	type Job struct {
		Title string `validate:"ValidateJobTitle"`
	}
	validate.RegisterValidation("ValidateJobTitle", ValidateJobTitle)
	jobEmpty := Job{
		Title: "",
	}
	errEmpty := validate.Struct(jobEmpty)
	assert.Error(t, errEmpty)
	jobNotEmpty := Job{
		Title: "manager",
	}
	errNotEmpty := validate.Struct(jobNotEmpty)
	assert.NoError(t, errNotEmpty)
	jobNotEmptyDir := Job{
		Title: "director",
	}
	errNotEmptyDir := validate.Struct(jobNotEmptyDir)
	assert.NoError(t, errNotEmptyDir)

	jobNotEmptyStaff := Job{
		Title: "staff",
	}
	errNotEmptyStaff := validate.Struct(jobNotEmptyStaff)
	assert.NoError(t, errNotEmptyStaff)

	jobNotEmptyUnKnown := Job{
		Title: "staffxxx",
	}
	errNotEmptyUnknown := validate.Struct(jobNotEmptyUnKnown)
	assert.Error(t, errNotEmptyUnknown)
}

func TestValidateEmployee(t *testing.T) {
	validate := validator.New()
	type EmployeeStruct struct {
		Name string `validate:"ValidateEmployee"`
	}
	validate.RegisterValidation("ValidateEmployee", ValidateEmployee)
	employeeEmpty := EmployeeStruct{
		Name: "",
	}
	errEmpty := validate.Struct(employeeEmpty)
	assert.Error(t, errEmpty)
	employeeNotEmpty := EmployeeStruct{
		Name: "John Doe",
	}
	errNotEmpty := validate.Struct(employeeNotEmpty)
	assert.NoError(t, errNotEmpty)
}

func TestValidatorPassword(t *testing.T) {
	validate := validator.New()
	type User struct {
		Password string `validate:"ValidatorPassword"`
	}
	validate.RegisterValidation("ValidatorPassword", ValidatorPassword)
	userEmpty := User{
		Password: "",
	}

	errEmpty := validate.Struct(userEmpty)
	assert.Error(t, errEmpty)

	userNotEmpty := User{
		Password: "P@ssw0rd",
	}

	errNotEmpty := validate.Struct(userNotEmpty)
	assert.NoError(t, errNotEmpty)
}

func TestGetCustomMessage(t *testing.T) {
	usernameMessage := GetCustomMessage("ValidUsername", "username")
	assert.Equal(t, "invalid request username", usernameMessage)
	passwordMessage := GetCustomMessage("ValidPassword", "password")
	assert.Equal(t, "invalid request password", passwordMessage)
	validateCompanyName := GetCustomMessage("ValidatorCompanyName", "password")
	assert.Equal(t, "invalid request password", validateCompanyName)
	otherMessage := GetCustomMessage("OtherError", "field")
	assert.Equal(t, "invalid request field", otherMessage)
}
