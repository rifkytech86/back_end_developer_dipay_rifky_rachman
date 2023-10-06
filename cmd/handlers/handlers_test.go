package handlers

import (
	"github.com/dipay/api"
	"github.com/dipay/bootstrap"
	"github.com/dipay/controller"
	"github.com/dipay/controller/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func TestServerInterfaceWrapper_Hello(t *testing.T) {
	type fields struct {
		UserAdminHandler        controller.IUserAdminController
		CompanyHandler          controller.ICompanyController
		EmployeeController      controller.IEmployeeController
		CountriesController     controller.ICountriesController
		DuplicateZeroController controller.IDuplicateZeroController
	}
	type args struct {
		ctx   echo.Context
		param api.HelloParams
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		mockUserName error
		wantErr      bool
	}{
		{
			name:         "hello",
			args:         args{},
			mockUserName: nil,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &ServerInterfaceWrapper{}
			mockUserAdmin := new(mocks.IUserAdminController)
			mockUserAdmin.On("Hello", mock.Anything, mock.Anything).Return(tt.mockUserName)
			w.UserAdminHandler = mockUserAdmin
			if err := w.Hello(tt.args.ctx, tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("Hello() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServerInterfaceWrapper_Login(t *testing.T) {
	type fields struct {
		UserAdminHandler        controller.IUserAdminController
		CompanyHandler          controller.ICompanyController
		EmployeeController      controller.IEmployeeController
		CountriesController     controller.ICountriesController
		DuplicateZeroController controller.IDuplicateZeroController
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		mockUserName error
		wantErr      bool
	}{
		{
			name:         "Login",
			args:         args{},
			mockUserName: nil,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &ServerInterfaceWrapper{}
			mockUserAdmin := new(mocks.IUserAdminController)
			mockUserAdmin.On("Login", mock.Anything, mock.Anything).Return(tt.mockUserName)
			w.UserAdminHandler = mockUserAdmin
			if err := w.Login(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServerInterfaceWrapper_AddCompany(t *testing.T) {
	type fields struct {
		UserAdminHandler        controller.IUserAdminController
		CompanyHandler          controller.ICompanyController
		EmployeeController      controller.IEmployeeController
		CountriesController     controller.ICountriesController
		DuplicateZeroController controller.IDuplicateZeroController
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		mockCompanyHandler error
		wantErr            bool
	}{
		{
			name:               "AddCompany",
			args:               args{},
			mockCompanyHandler: nil,
			wantErr:            false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &ServerInterfaceWrapper{}
			mockCompanyHandler := new(mocks.ICompanyController)
			mockCompanyHandler.On("AddCompany", mock.Anything).Return(tt.mockCompanyHandler)
			w.CompanyHandler = mockCompanyHandler
			if err := w.AddCompany(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("AddCompany() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServerInterfaceWrapper_GetCompany(t *testing.T) {
	type fields struct {
		UserAdminHandler        controller.IUserAdminController
		CompanyHandler          controller.ICompanyController
		EmployeeController      controller.IEmployeeController
		CountriesController     controller.ICountriesController
		DuplicateZeroController controller.IDuplicateZeroController
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		mockCompanyHandler error
		wantErr            bool
	}{
		{
			name:               "GetCompany",
			args:               args{},
			mockCompanyHandler: nil,
			wantErr:            false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &ServerInterfaceWrapper{}
			mockCompanyHandler := new(mocks.ICompanyController)
			mockCompanyHandler.On("GetCompany", mock.Anything).Return(tt.mockCompanyHandler)
			w.CompanyHandler = mockCompanyHandler
			if err := w.GetCompany(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("GetCompany() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServerInterfaceWrapper_UpdateCompanyStatusActive(t *testing.T) {
	type fields struct {
		UserAdminHandler        controller.IUserAdminController
		CompanyHandler          controller.ICompanyController
		EmployeeController      controller.IEmployeeController
		CountriesController     controller.ICountriesController
		DuplicateZeroController controller.IDuplicateZeroController
	}
	type args struct {
		ctx echo.Context
		id  string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		mockCompanyHandler error
		wantErr            bool
	}{
		{
			name:               "UpdateCompanyStatusActive",
			args:               args{},
			mockCompanyHandler: nil,
			wantErr:            false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &ServerInterfaceWrapper{}
			mockCompanyHandler := new(mocks.ICompanyController)
			mockCompanyHandler.On("UpdateCompanyStatusActive", mock.Anything, mock.Anything).Return(tt.mockCompanyHandler)
			w.CompanyHandler = mockCompanyHandler
			if err := w.UpdateCompanyStatusActive(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCompanyStatusActive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServerInterfaceWrapper_AddEmployee(t *testing.T) {
	type fields struct {
		UserAdminHandler        controller.IUserAdminController
		CompanyHandler          controller.ICompanyController
		EmployeeController      controller.IEmployeeController
		CountriesController     controller.ICountriesController
		DuplicateZeroController controller.IDuplicateZeroController
	}
	type args struct {
		ctx echo.Context
		id  string
	}
	tests := []struct {
		name                   string
		fields                 fields
		mockAddEmployeeHandler error
		args                   args
		wantErr                bool
	}{
		{
			name:                   "AddEmployee",
			args:                   args{},
			mockAddEmployeeHandler: nil,
			wantErr:                false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &ServerInterfaceWrapper{}
			mockAddEmployeeHandler := new(mocks.IEmployeeController)
			mockAddEmployeeHandler.On("AddEmployee", mock.Anything, mock.Anything).Return(tt.mockAddEmployeeHandler)
			w.EmployeeController = mockAddEmployeeHandler
			if err := w.AddEmployee(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AddEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServerInterfaceWrapper_GetEmployeeByID(t *testing.T) {
	type fields struct {
		UserAdminHandler        controller.IUserAdminController
		CompanyHandler          controller.ICompanyController
		EmployeeController      controller.IEmployeeController
		CountriesController     controller.ICountriesController
		DuplicateZeroController controller.IDuplicateZeroController
	}
	type args struct {
		ctx echo.Context
		id  string
	}
	tests := []struct {
		name                        string
		fields                      fields
		args                        args
		mockAGetEmployeeByIDHandler error
		wantErr                     bool
	}{
		{
			name:                        "GetEmployeeByID",
			args:                        args{},
			mockAGetEmployeeByIDHandler: nil,
			wantErr:                     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &ServerInterfaceWrapper{}
			mockAddEmployeeHandler := new(mocks.IEmployeeController)
			mockAddEmployeeHandler.On("GetEmployeeByID", mock.Anything, mock.Anything).Return(tt.mockAGetEmployeeByIDHandler)
			w.EmployeeController = mockAddEmployeeHandler
			if err := w.GetEmployeeByID(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GetEmployeeByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServerInterfaceWrapper_GetEmployeeByCompanyID(t *testing.T) {
	type fields struct {
		UserAdminHandler        controller.IUserAdminController
		CompanyHandler          controller.ICompanyController
		EmployeeController      controller.IEmployeeController
		CountriesController     controller.ICountriesController
		DuplicateZeroController controller.IDuplicateZeroController
	}
	type args struct {
		ctx echo.Context
		id  string
	}
	tests := []struct {
		name                                string
		fields                              fields
		mockAGetEmployeeByCompanyIDDHandler error
		args                                args
		wantErr                             bool
	}{
		{
			name:                                "GetEmployeeByCompanyID",
			args:                                args{},
			mockAGetEmployeeByCompanyIDDHandler: nil,
			wantErr:                             false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &ServerInterfaceWrapper{}
			mockAddEmployeeHandler := new(mocks.IEmployeeController)
			mockAddEmployeeHandler.On("GetEmployeeByCompanyID", mock.Anything, mock.Anything).Return(tt.mockAGetEmployeeByCompanyIDDHandler)
			w.EmployeeController = mockAddEmployeeHandler
			if err := w.GetEmployeeByCompanyID(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GetEmployeeByCompanyID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServerInterfaceWrapper_UpdateEmployeeData(t *testing.T) {
	type fields struct {
		UserAdminHandler        controller.IUserAdminController
		CompanyHandler          controller.ICompanyController
		EmployeeController      controller.IEmployeeController
		CountriesController     controller.ICountriesController
		DuplicateZeroController controller.IDuplicateZeroController
	}
	type args struct {
		ctx        echo.Context
		companyId  string
		employeeId string
	}
	tests := []struct {
		name                          string
		fields                        fields
		mockUpdateEmployeeDataHandler error
		args                          args
		wantErr                       bool
	}{
		{
			name:                          "UpdateEmployeeData",
			args:                          args{},
			mockUpdateEmployeeDataHandler: nil,
			wantErr:                       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &ServerInterfaceWrapper{}
			mockAddEmployeeHandler := new(mocks.IEmployeeController)
			mockAddEmployeeHandler.On("UpdateEmployeeData", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockUpdateEmployeeDataHandler)
			w.EmployeeController = mockAddEmployeeHandler
			if err := w.UpdateEmployeeData(tt.args.ctx, tt.args.companyId, tt.args.employeeId); (err != nil) != tt.wantErr {
				t.Errorf("UpdateEmployeeData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServerInterfaceWrapper_DeleteEmployeeByID(t *testing.T) {
	type fields struct {
		UserAdminHandler        controller.IUserAdminController
		CompanyHandler          controller.ICompanyController
		EmployeeController      controller.IEmployeeController
		CountriesController     controller.ICountriesController
		DuplicateZeroController controller.IDuplicateZeroController
	}
	type args struct {
		ctx echo.Context
		id  string
	}
	tests := []struct {
		name                          string
		fields                        fields
		mockDeleteEmployeeByIDHandler error
		args                          args
		wantErr                       bool
	}{
		{
			name:                          "DeleteEmployeeByID",
			args:                          args{},
			mockDeleteEmployeeByIDHandler: nil,
			wantErr:                       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &ServerInterfaceWrapper{}
			mockAddEmployeeHandler := new(mocks.IEmployeeController)
			mockAddEmployeeHandler.On("DeleteEmployeeByID", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockDeleteEmployeeByIDHandler)
			w.EmployeeController = mockAddEmployeeHandler
			if err := w.DeleteEmployeeByID(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteEmployeeByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServerInterfaceWrapper_GetDataCountries(t *testing.T) {
	type fields struct {
		UserAdminHandler        controller.IUserAdminController
		CompanyHandler          controller.ICompanyController
		EmployeeController      controller.IEmployeeController
		CountriesController     controller.ICountriesController
		DuplicateZeroController controller.IDuplicateZeroController
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name                        string
		fields                      fields
		mockGetDataCountriesHandler error
		args                        args
		wantErr                     bool
	}{
		{
			name:                        "GetDataCountries",
			args:                        args{},
			mockGetDataCountriesHandler: nil,
			wantErr:                     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &ServerInterfaceWrapper{}
			mockCountiesHandler := new(mocks.ICountriesController)
			mockCountiesHandler.On("GetDataCountries", mock.Anything).Return(tt.mockGetDataCountriesHandler)
			w.CountriesController = mockCountiesHandler
			if err := w.GetDataCountries(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("GetDataCountries() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServerInterfaceWrapper_DuplicateZero(t *testing.T) {
	type fields struct {
		UserAdminHandler        controller.IUserAdminController
		CompanyHandler          controller.ICompanyController
		EmployeeController      controller.IEmployeeController
		CountriesController     controller.ICountriesController
		DuplicateZeroController controller.IDuplicateZeroController
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name                     string
		fields                   fields
		mockDuplicateZeroHandler error
		args                     args
		wantErr                  bool
	}{
		{
			name:                     "DuplicateZero",
			args:                     args{},
			mockDuplicateZeroHandler: nil,
			wantErr:                  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &ServerInterfaceWrapper{}
			mockDuplicateZeroHandler := new(mocks.IDuplicateZeroController)
			mockDuplicateZeroHandler.On("DuplicateZero", mock.Anything).Return(tt.mockDuplicateZeroHandler)
			w.DuplicateZeroController = mockDuplicateZeroHandler
			if err := w.DuplicateZero(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("DuplicateZero() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewServiceInitial(t *testing.T) {
	type args struct {
		app bootstrap.Application
	}
	tests := []struct {
		name string
		args args
		want MyHandler
	}{
		{
			name: "initial service",
			args: args{},
			want: MyHandler{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServiceInitial(tt.args.app); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServiceInitial() = %v, want %v", got, tt.want)
			}
		})
	}
}
