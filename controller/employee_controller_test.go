package controller

import (
	"errors"
	"github.com/dipay/internal"
	"github.com/dipay/internal/validations"
	"github.com/dipay/internal/validations/mocks"
	"github.com/dipay/model"
	"github.com/dipay/usecase"
	useCaseMock "github.com/dipay/usecase/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func Test_employeeController_AddEmployee(t *testing.T) {
	type fields struct {
		EmployeeUseCase usecase.IEmployeeUseCase
		ContextTimeOut  int
		Validator       validations.IValidator
	}
	type mockAddEmployee struct {
		employeeID string
		companyID  string
		err        error
	}
	type args struct {
		c         echo.Context
		companyId string
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		mockValidator   []validations.ValidationError
		mockAddEmployee mockAddEmployee
		wantErr         bool
	}{
		{
			name: "error bind add employee",
			args: args{
				c:         failedInitialEcho(),
				companyId: "1",
			},
			wantErr: false,
		},
		{
			name: "error validator",
			args: args{
				c:         successInitialEcho(),
				companyId: "1",
			},
			mockValidator: []validations.ValidationError{
				{
					Field: "email",
					Error: "error email is empty",
				},
			},
			wantErr: false,
		},
		{
			name: "error add employee",
			args: args{
				c:         successInitialEcho(),
				companyId: "1",
			},
			mockValidator: nil,
			mockAddEmployee: mockAddEmployee{
				err: errors.New(internal.ErrorInternalServer.String()),
			},
			wantErr: false,
		},
		{
			name: "happy flow",
			args: args{
				c:         successInitialEcho(),
				companyId: "1",
			},
			mockValidator: nil,
			mockAddEmployee: mockAddEmployee{
				err: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &employeeController{}
			mockValidator := new(mocks.IValidator)
			mockValidator.On("Struct", mock.Anything).Return(tt.mockValidator)
			e.Validator = mockValidator
			mockEmployeeUseCase := new(useCaseMock.IEmployeeUseCase)
			mockEmployeeUseCase.On("AddEmployee", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockAddEmployee.employeeID, tt.mockAddEmployee.companyID, tt.mockAddEmployee.err)
			e.EmployeeUseCase = mockEmployeeUseCase

			if err := e.AddEmployee(tt.args.c, tt.args.companyId); (err != nil) != tt.wantErr {
				t.Errorf("AddEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_employeeController_GetEmployeeByID(t *testing.T) {
	type fields struct {
		EmployeeUseCase usecase.IEmployeeUseCase
		ContextTimeOut  int
		Validator       validations.IValidator
	}
	type mockGetEmployeeByID struct {
		data *model.Employees
		err  error
	}
	type args struct {
		c          echo.Context
		employeeID string
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		mockGetEmployeeByID mockGetEmployeeByID
		wantErr             bool
	}{
		{
			name: "error employee id empty",
			args: args{
				c:          successInitialEcho(),
				employeeID: "",
			},
		},
		{
			name: "error employee id empty",
			args: args{
				c:          successInitialEcho(),
				employeeID: "1",
			},
			mockGetEmployeeByID: mockGetEmployeeByID{
				err: errors.New(internal.ErrorInternalServer.String()),
			},
		},
		{
			name: "Happy Flow",
			args: args{
				c:          successInitialEcho(),
				employeeID: "1",
			},
			mockGetEmployeeByID: mockGetEmployeeByID{
				data: &model.Employees{},
				err:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &employeeController{}
			mockEmployeeUseCase := new(useCaseMock.IEmployeeUseCase)
			mockEmployeeUseCase.On("GetEmployeeByID", mock.Anything, mock.Anything).Return(tt.mockGetEmployeeByID.data, tt.mockGetEmployeeByID.err)
			e.EmployeeUseCase = mockEmployeeUseCase
			if err := e.GetEmployeeByID(tt.args.c, tt.args.employeeID); (err != nil) != tt.wantErr {
				t.Errorf("GetEmployeeByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_employeeController_GetEmployeeByCompanyID(t *testing.T) {
	type fields struct {
		EmployeeUseCase usecase.IEmployeeUseCase
		ContextTimeOut  int
		Validator       validations.IValidator
	}
	type mockGetEmployeeByCompanyID struct {
		data []*model.Employees
		err  error
	}
	type args struct {
		c         echo.Context
		companyID string
	}
	tests := []struct {
		name                       string
		fields                     fields
		args                       args
		mockGetEmployeeByCompanyID mockGetEmployeeByCompanyID
		wantErr                    bool
	}{
		{
			name: "error parameter company empty",
			args: args{
				c:         successInitialEcho(),
				companyID: "",
			},
		},
		{
			name: "error get employee by company",
			args: args{
				c:         successInitialEcho(),
				companyID: "1",
			},
			mockGetEmployeeByCompanyID: mockGetEmployeeByCompanyID{
				data: nil,
				err:  errors.New(internal.ErrorInternalServer.String()),
			},
		},
		{
			name: "Happy GetEmployeeByCompanyID",
			args: args{
				c:         successInitialEcho(),
				companyID: "1",
			},
			mockGetEmployeeByCompanyID: mockGetEmployeeByCompanyID{
				data: []*model.Employees{
					{
						Name: "tester",
					},
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &employeeController{}
			mockEmployeeUseCase := new(useCaseMock.IEmployeeUseCase)
			mockEmployeeUseCase.On("GetEmployeeByCompanyID", mock.Anything, mock.Anything).Return(tt.mockGetEmployeeByCompanyID.data, tt.mockGetEmployeeByCompanyID.err)
			e.EmployeeUseCase = mockEmployeeUseCase
			if err := e.GetEmployeeByCompanyID(tt.args.c, tt.args.companyID); (err != nil) != tt.wantErr {
				t.Errorf("GetEmployeeByCompanyID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_employeeController_UpdateEmployeeData(t *testing.T) {
	type fields struct {
		EmployeeUseCase usecase.IEmployeeUseCase
		ContextTimeOut  int
		Validator       validations.IValidator
	}
	type mockUpdateEmployeeData struct {
		employeeID string
		companyID  string
		err        error
	}
	type args struct {
		c          echo.Context
		companyID  string
		employeeId string
	}
	tests := []struct {
		name                   string
		fields                 fields
		mockValidator          []validations.ValidationError
		args                   args
		mockUpdateEmployeeData mockUpdateEmployeeData
		wantErr                bool
	}{
		{
			name: "error bind ",
			args: args{
				c:          failedInitialEcho(),
				companyID:  "",
				employeeId: "",
			},
		},
		{
			name: "error validator",
			args: args{
				c:          successInitialEcho(),
				companyID:  "",
				employeeId: "",
			},
			mockValidator: []validations.ValidationError{
				{
					Field: "Email",
					Error: "error Email is empty",
				},
			},
		},
		{
			name: "error update employee",
			args: args{
				c:          successInitialEcho(),
				companyID:  "",
				employeeId: "",
			},
			mockValidator: nil,
			mockUpdateEmployeeData: mockUpdateEmployeeData{
				err: errors.New(internal.ErrorInternalServer.String()),
			},
		},
		{
			name: "Happy Flow",
			args: args{
				c:          successInitialEcho(),
				companyID:  "",
				employeeId: "",
			},
			mockValidator: nil,
			mockUpdateEmployeeData: mockUpdateEmployeeData{
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &employeeController{}
			mockValidator := new(mocks.IValidator)
			mockValidator.On("Struct", mock.Anything).Return(tt.mockValidator)
			e.Validator = mockValidator
			mockEmployeeUseCase := new(useCaseMock.IEmployeeUseCase)
			mockEmployeeUseCase.On("UpdateEmployeeData", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.mockUpdateEmployeeData.employeeID, tt.mockUpdateEmployeeData.companyID, tt.mockUpdateEmployeeData.err)
			e.EmployeeUseCase = mockEmployeeUseCase
			if err := e.UpdateEmployeeData(tt.args.c, tt.args.companyID, tt.args.employeeId); (err != nil) != tt.wantErr {
				t.Errorf("UpdateEmployeeData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_employeeController_DeleteEmployeeByID(t *testing.T) {
	type fields struct {
		EmployeeUseCase usecase.IEmployeeUseCase
		ContextTimeOut  int
		Validator       validations.IValidator
	}
	type args struct {
		c          echo.Context
		employeeId string
	}
	tests := []struct {
		name                   string
		fields                 fields
		args                   args
		mockDeleteEmployeeByID error
		wantErr                bool
	}{
		{
			name: "err delete employee id empty",
			args: args{
				c:          successInitialEcho(),
				employeeId: "",
			},
		},
		{
			name: "error delete employee id",
			args: args{
				c:          successInitialEcho(),
				employeeId: "1",
			},
			mockDeleteEmployeeByID: errors.New(internal.ErrorInternalServer.String()),
		},
		{
			name: "Happy Flow",
			args: args{
				c:          successInitialEcho(),
				employeeId: "1",
			},
			mockDeleteEmployeeByID: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &employeeController{}
			mockEmployeeUseCase := new(useCaseMock.IEmployeeUseCase)
			mockEmployeeUseCase.On("DeleteEmployeeByID", mock.Anything, mock.Anything).Return(tt.mockDeleteEmployeeByID)
			e.EmployeeUseCase = mockEmployeeUseCase
			if err := e.DeleteEmployeeByID(tt.args.c, tt.args.employeeId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteEmployeeByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewEmployeeController(t *testing.T) {
	type args struct {
		employeeUseCase usecase.IEmployeeUseCase
		contextTimeOut  int
		validator       validations.IValidator
	}
	tests := []struct {
		name string
		args args
		want IEmployeeController
	}{
		{
			name: "initial employee controller",
			args: args{},
			want: &employeeController{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmployeeController(tt.args.employeeUseCase, tt.args.contextTimeOut, tt.args.validator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmployeeController() = %v, want %v", got, tt.want)
			}
		})
	}
}
