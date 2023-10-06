package controller

import (
	"errors"
	"github.com/dipay/api"
	"github.com/dipay/internal"
	"github.com/dipay/internal/validations"
	"github.com/dipay/internal/validations/mocks"
	"github.com/dipay/usecase"
	useCaseMock "github.com/dipay/usecase/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func Test_userAdminController_Hello(t *testing.T) {
	type fields struct {
		UserAdminUseCase usecase.IUseCaseUserAdmin
		ContextTimeOut   int
		Validator        validations.IValidator
	}
	type args struct {
		ctx   echo.Context
		param api.HelloParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test hello",
			args: args{
				ctx: successInitialEcho(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userAdminController{
				UserAdminUseCase: tt.fields.UserAdminUseCase,
				ContextTimeOut:   tt.fields.ContextTimeOut,
				Validator:        tt.fields.Validator,
			}
			if err := u.Hello(tt.args.ctx, tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("Hello() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userAdminController_Login(t *testing.T) {
	type fields struct {
		UserAdminUseCase usecase.IUseCaseUserAdmin
		ContextTimeOut   int
		Validator        validations.IValidator
	}
	type mockLogin struct {
		data string
		err  error
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		mockValidator []validations.ValidationError
		mockLogin     mockLogin
		wantErr       bool
	}{
		{
			name: "error bind",
			args: args{
				c: failedInitialEcho(),
			},
		},
		{
			name: "error validator",
			args: args{
				c: successInitialEcho(),
			},
			mockValidator: []validations.ValidationError{
				{
					Field: "Password",
					Error: "error Password is empty",
				},
			},
		},
		{
			name: "error login",
			args: args{
				c: successInitialEcho(),
			},
			mockValidator: nil,
			mockLogin: mockLogin{
				err: errors.New(internal.ErrorInternalServer.String()),
			},
		},
		{
			name: "Happy Flow",
			args: args{
				c: successInitialEcho(),
			},
			mockValidator: nil,
			mockLogin: mockLogin{
				data: "",
				err:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userAdminController{}
			mockValidator := new(mocks.IValidator)
			mockValidator.On("Struct", mock.Anything).Return(tt.mockValidator)
			u.Validator = mockValidator

			mockUserAdminUseCase := new(useCaseMock.IUseCaseUserAdmin)
			mockUserAdminUseCase.On("Login", mock.Anything, mock.Anything).Return(tt.mockLogin.data, tt.mockLogin.err)
			u.UserAdminUseCase = mockUserAdminUseCase

			if err := u.Login(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewUserAdminController(t *testing.T) {
	type args struct {
		userAdminUseCase usecase.IUseCaseUserAdmin
		contextTimeOut   int
		validator        validations.IValidator
	}
	tests := []struct {
		name string
		args args
		want IUserAdminController
	}{
		{
			name: "initial user admin controller",
			args: args{},
			want: &userAdminController{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserAdminController(tt.args.userAdminUseCase, tt.args.contextTimeOut, tt.args.validator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserAdminController() = %v, want %v", got, tt.want)
			}
		})
	}
}
