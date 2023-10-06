package controller

import (
	"errors"
	"github.com/dipay/internal"
	"github.com/dipay/internal/validations"
	"github.com/dipay/internal/validations/mocks"
	"github.com/dipay/model"
	"github.com/dipay/usecase"
	mockUseCase "github.com/dipay/usecase/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func failedInitialEcho() echo.Context {
	e := echo.New()
	userErrJSON := `}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userErrJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c
}
func successInitialEcho() echo.Context {
	e := echo.New()
	userErrJSON := `{"password":"asdqwe1A@","phone_number":"+62"}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userErrJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c
}

func Test_companyController_AddCompany(t *testing.T) {
	type fields struct {
		CompanyUseCase usecase.ICompanyUseCase
		ContextTimeOut int
		Validator      validations.IValidator
	}
	type mockAddCompany struct {
		data string
		err  error
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantErr        bool
		mockValidator  []validations.ValidationError
		mockAddCompany mockAddCompany
	}{
		{
			name: "error bind",
			args: args{
				c: failedInitialEcho(),
			},
			wantErr: false,
		},
		{
			name: "error validator",
			args: args{
				c: successInitialEcho(),
			},
			wantErr: false,
			mockValidator: []validations.ValidationError{
				{
					Field: "Address",
					Error: "error address is empty",
				},
			},
		},
		{
			name: "error add company",
			args: args{
				c: successInitialEcho(),
			},
			wantErr:       false,
			mockValidator: nil,
			mockAddCompany: mockAddCompany{
				data: "",
				err:  errors.New(internal.ErrorInternalServer.String()),
			},
		},
		{
			name: "Happy flow",
			args: args{
				c: successInitialEcho(),
			},
			wantErr:       false,
			mockValidator: nil,
			mockAddCompany: mockAddCompany{
				data: "12334",
				err:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &companyController{}
			mockValidator := new(mocks.IValidator)
			mockValidator.On("Struct", mock.Anything).Return(tt.mockValidator)
			u.Validator = mockValidator
			mockCompanyUseCase := new(mockUseCase.ICompanyUseCase)
			mockCompanyUseCase.On("AddCompany", mock.Anything, mock.Anything).Return(tt.mockAddCompany.data, tt.mockAddCompany.err)
			u.CompanyUseCase = mockCompanyUseCase

			if err := u.AddCompany(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AddCompany() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_companyController_GetCompany(t *testing.T) {
	type fields struct {
		CompanyUseCase usecase.ICompanyUseCase
		ContextTimeOut int
		Validator      validations.IValidator
	}
	type mockGetCompany struct {
		data []*model.Companies
		err  error
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantErr        bool
		mockGetCompany mockGetCompany
	}{
		{
			name: "error get company",
			args: args{
				c: successInitialEcho(),
			},
			mockGetCompany: mockGetCompany{
				data: nil,
				err:  errors.New(internal.ErrorInternalServer.String()),
			},
			wantErr: false,
		},
		{
			name: "Happy Flow ",
			args: args{
				c: successInitialEcho(),
			},
			mockGetCompany: mockGetCompany{
				data: []*model.Companies{
					{
						CompanyName: "tester",
					},
				},
				err: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &companyController{}
			mockCompanyUseCase := new(mockUseCase.ICompanyUseCase)
			mockCompanyUseCase.On("GetCompany", mock.Anything).Return(tt.mockGetCompany.data, tt.mockGetCompany.err)
			u.CompanyUseCase = mockCompanyUseCase

			if err := u.GetCompany(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetCompany() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_companyController_UpdateCompanyStatusActive(t *testing.T) {
	type fields struct {
		CompanyUseCase usecase.ICompanyUseCase
		ContextTimeOut int
		Validator      validations.IValidator
	}
	type mockUpdateStatusActive struct {
		idAffected string
		isActive   bool
		err        error
	}
	type args struct {
		c  echo.Context
		id string
	}
	tests := []struct {
		name                   string
		fields                 fields
		args                   args
		mockUpdateStatusActive mockUpdateStatusActive
		wantErr                bool
	}{
		{
			name: "error update company status active",
			args: args{
				c: successInitialEcho(),
			},
			mockUpdateStatusActive: mockUpdateStatusActive{
				idAffected: "",
				isActive:   false,
				err:        errors.New(internal.ErrorInternalServer.String()),
			},
		},
		{
			name: "happy flow",
			args: args{
				c: successInitialEcho(),
			},
			mockUpdateStatusActive: mockUpdateStatusActive{
				idAffected: "1",
				isActive:   true,
				err:        nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &companyController{}
			mockCompanyUseCase := new(mockUseCase.ICompanyUseCase)
			mockCompanyUseCase.On("UpdateCompanyStatusActive", mock.Anything, mock.Anything).Return(tt.mockUpdateStatusActive.idAffected, tt.mockUpdateStatusActive.isActive, tt.mockUpdateStatusActive.err)
			u.CompanyUseCase = mockCompanyUseCase
			if err := u.UpdateCompanyStatusActive(tt.args.c, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCompanyStatusActive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCompanyController(t *testing.T) {
	type args struct {
		companyUseCase usecase.ICompanyUseCase
		contextTimeOut int
		validator      validations.IValidator
	}
	tests := []struct {
		name string
		args args
		want ICompanyController
	}{
		{
			name: "initial company controller",
			args: args{},
			want: &companyController{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCompanyController(tt.args.companyUseCase, tt.args.contextTimeOut, tt.args.validator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCompanyController() = %v, want %v", got, tt.want)
			}
		})
	}
}
