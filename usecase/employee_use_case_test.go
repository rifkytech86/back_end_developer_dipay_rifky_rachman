package usecase

import (
	"context"
	"errors"
	"github.com/dipay/api"
	"github.com/dipay/internal"
	"github.com/dipay/model"
	"github.com/dipay/repositories"
	"github.com/dipay/repositories/mocks"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func Test_employeeUserAdmin_AddEmployee(t *testing.T) {
	type fields struct {
		EmployeeRepository repositories.IEmployeeRepository
		CompanyRepository  repositories.ICompanyRepository
	}
	type args struct {
		ctx       context.Context
		companyId string
		req       *api.AddEmployeeJSONBody
	}
	type mockCreate struct {
		id  *primitive.ObjectID
		err error
	}
	tests := []struct {
		name                 string
		fields               fields
		args                 args
		want                 string
		want1                string
		wantErr              bool
		mockFetchOne         error
		mockFetchOneEmployee error
		mockCreate           mockCreate
	}{
		{
			name: "error object from hex",
			args: args{
				ctx: context.TODO(),
			},
			wantErr:      true,
			mockFetchOne: errors.New(internal.ErrorInternalServer.String()),
			mockCreate: mockCreate{
				id:  nil,
				err: errors.New(internal.ErrorInternalServer.String()),
			},
			mockFetchOneEmployee: errors.New(internal.ErrorInternalServer.String()),
		},
		{
			name: "error fetch one",
			args: args{
				companyId: "651fb48744d0b172ae3ff0fa",
			},
			wantErr:      true,
			mockFetchOne: errors.New(internal.ErrorInternalServer.String()),
			mockCreate: mockCreate{
				id:  nil,
				err: errors.New(internal.ErrorInternalServer.String()),
			},
			mockFetchOneEmployee: errors.New(internal.ErrorInternalServer.String()),
		},
		{
			name: "error fetch one no document update",
			args: args{
				companyId: "651fb48744d0b172ae3ff0fa",
			},
			wantErr:      true,
			mockFetchOne: mongo.ErrNoDocuments,
			mockCreate: mockCreate{
				id:  nil,
				err: errors.New(internal.ErrorInternalServer.String()),
			},
			mockFetchOneEmployee: errors.New(internal.ErrorInternalServer.String()),
		},
		{
			name: "error create employee",
			args: args{
				companyId: "651fb48744d0b172ae3ff0fa",
				req: &api.AddEmployeeJSONBody{
					Name:        "tester",
					Email:       "rmail.com",
					PhoneNumber: "+6209809890",
					Jobtitle:    "jobtitle",
				},
			},
			wantErr:      true,
			mockFetchOne: nil,
			mockCreate: mockCreate{
				id:  nil,
				err: errors.New(internal.ErrorInternalServer.String()),
			},
			mockFetchOneEmployee: errors.New(internal.ErrorInternalServer.String()),
		},
		{
			name: "error create employee not affected",
			args: args{
				companyId: "651fb48744d0b172ae3ff0fa",
				req: &api.AddEmployeeJSONBody{
					Name:        "tester",
					Email:       "rmail.com",
					PhoneNumber: "+6209809890",
					Jobtitle:    "jobtitle",
				},
			},
			wantErr:      true,
			mockFetchOne: nil,
			mockCreate: mockCreate{
				id:  nil,
				err: mongo.ErrNoDocuments,
			},
			mockFetchOneEmployee: errors.New(internal.ErrorInternalServer.String()),
		},
		{
			name: "error fetch one",
			args: args{
				companyId: "651fb48744d0b172ae3ff0fa",
				req: &api.AddEmployeeJSONBody{
					Name:        "tester",
					Email:       "rmail.com",
					PhoneNumber: "+6209809890",
					Jobtitle:    "jobtitle",
				},
			},
			wantErr:      true,
			mockFetchOne: nil,
			mockCreate: mockCreate{
				id:  nil,
				err: nil,
			},
			mockFetchOneEmployee: errors.New(internal.ErrorInternalServer.String()),
		},
		{
			name: "Happy flow",
			args: args{
				companyId: "651fb48744d0b172ae3ff0fa",
				req: &api.AddEmployeeJSONBody{
					Name:        "tester",
					Email:       "rmail.com",
					PhoneNumber: "+6209809890",
					Jobtitle:    "jobtitle",
				},
			},
			wantErr:      false,
			mockFetchOne: nil,
			mockCreate: mockCreate{
				id:  nil,
				err: nil,
			},
			mockFetchOneEmployee: nil,
			want:                 "000000000000000000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &employeeUserAdmin{}
			mockCompanyRepo := new(mocks.ICompanyRepository)
			mockCompanyRepo.On("FetchOne", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockFetchOne)
			e.CompanyRepository = mockCompanyRepo

			mockEmployeeRepo := new(mocks.IEmployeeRepository)
			mockEmployeeRepo.On("Create", mock.Anything, mock.Anything).Return(tt.mockCreate.id, tt.mockCreate.err)
			mockEmployeeRepo.On("FetchOne", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockFetchOneEmployee)
			mockEmployeeRepo.On("SendEmail", mock.Anything, mock.Anything).Return(nil)
			e.EmployeeRepository = mockEmployeeRepo

			got, got1, err := e.AddEmployee(tt.args.ctx, tt.args.companyId, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddEmployee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddEmployee() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AddEmployee() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_employeeUserAdmin_GetEmployeeByID(t *testing.T) {
	type fields struct {
		EmployeeRepository repositories.IEmployeeRepository
		CompanyRepository  repositories.ICompanyRepository
	}
	type args struct {
		ctx        context.Context
		employeeID string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		want         *model.Employees
		wantErr      bool
		mockFetchOne error
	}{
		{
			name: "error primitive object",
			args: args{
				ctx:        context.TODO(),
				employeeID: "1",
			},
			wantErr: true,
		},
		{
			name: "error fetch one",
			args: args{
				ctx:        context.TODO(),
				employeeID: "651fb48744d0b172ae3ff0fa",
			},
			wantErr:      true,
			mockFetchOne: errors.New(internal.ErrorInternalServer.String()),
		},
		{
			name: "error fetch one error no document",
			args: args{
				ctx:        context.TODO(),
				employeeID: "651fb48744d0b172ae3ff0fa",
			},
			wantErr:      true,
			mockFetchOne: mongo.ErrNoDocuments,
		},
		{
			name: "Happy flow",
			args: args{
				ctx:        context.TODO(),
				employeeID: "651fb48744d0b172ae3ff0fa",
			},
			wantErr:      false,
			mockFetchOne: nil,
			want:         &model.Employees{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &employeeUserAdmin{}
			mockEmployeeRepo := new(mocks.IEmployeeRepository)
			mockEmployeeRepo.On("FetchOne", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockFetchOne)

			e.EmployeeRepository = mockEmployeeRepo

			got, err := e.GetEmployeeByID(tt.args.ctx, tt.args.employeeID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEmployeeByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEmployeeByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_employeeUserAdmin_GetEmployeeByCompanyID(t *testing.T) {
	type fields struct {
		EmployeeRepository repositories.IEmployeeRepository
		CompanyRepository  repositories.ICompanyRepository
	}
	type args struct {
		ctx       context.Context
		companyID string
	}
	type mockFetch struct {
		data []*model.Employees
		err  error
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      []*model.Employees
		wantErr   bool
		mockFetch mockFetch
	}{
		{
			name: "error fetch employee",
			args: args{
				ctx:       context.TODO(),
				companyID: "1",
			},
			mockFetch: mockFetch{
				data: nil,
				err:  errors.New(internal.ErrorInternalServer.String()),
			},
			wantErr: true,
		},
		{
			name: "error fetch employee",
			args: args{
				ctx:       context.TODO(),
				companyID: "1",
			},
			mockFetch: mockFetch{
				data: nil,
				err:  mongo.ErrNoDocuments,
			},
			wantErr: true,
		},
		{
			name: "Happy Flow",
			args: args{
				ctx:       context.TODO(),
				companyID: "1",
			},
			mockFetch: mockFetch{
				data: []*model.Employees{},
				err:  nil,
			},
			wantErr: false,
			want:    []*model.Employees{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &employeeUserAdmin{}
			mockEmployeeRepo := new(mocks.IEmployeeRepository)
			mockEmployeeRepo.On("Fetch", mock.Anything, mock.Anything).Return(tt.mockFetch.data, tt.mockFetch.err)
			e.EmployeeRepository = mockEmployeeRepo

			got, err := e.GetEmployeeByCompanyID(tt.args.ctx, tt.args.companyID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEmployeeByCompanyID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEmployeeByCompanyID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_employeeUserAdmin_UpdateEmployeeData(t *testing.T) {
	type fields struct {
		EmployeeRepository repositories.IEmployeeRepository
		CompanyRepository  repositories.ICompanyRepository
	}
	type args struct {
		ctx        context.Context
		companyID  string
		employeeID string
		req        *api.UpdateEmployeeDataJSONRequestBody
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		want         string
		want1        string
		wantErr      bool
		mockUpdate   error
		mockFetchOne error
	}{
		{
			name: "error data objectID from hex",
			args: args{
				ctx:        context.TODO(),
				companyID:  "1",
				employeeID: "1",
				req:        &api.UpdateEmployeeDataJSONRequestBody{},
			},
			mockUpdate: errors.New(internal.ErrorInternalServer.String()),
			wantErr:    true,
		},
		{
			name: "error update employee",
			args: args{
				ctx:        context.TODO(),
				companyID:  "651fb40e44d0b172ae3ff0f9",
				employeeID: "651fb40e44d0b172ae3ff0f9",
				req:        &api.UpdateEmployeeDataJSONRequestBody{},
			},
			mockUpdate: errors.New(internal.ErrorInternalServer.String()),
			wantErr:    true,
		},
		{
			name: "error update employee no modify",
			args: args{
				ctx:        context.TODO(),
				companyID:  "651fb40e44d0b172ae3ff0f9",
				employeeID: "651fb40e44d0b172ae3ff0f9",
				req:        &api.UpdateEmployeeDataJSONRequestBody{},
			},
			mockUpdate: errors.New(internal.ErrNoModifyUpdate.String()),
			wantErr:    true,
		},
		{
			name: "error get data employee",
			args: args{
				ctx:        context.TODO(),
				companyID:  "651fb40e44d0b172ae3ff0f9",
				employeeID: "651fb40e44d0b172ae3ff0f9",
				req:        &api.UpdateEmployeeDataJSONRequestBody{},
			},
			mockUpdate:   nil,
			mockFetchOne: errors.New(internal.ErrorInternalServer.String()),
			wantErr:      true,
		},
		{
			name: "Happy Flow",
			args: args{
				ctx:        context.TODO(),
				companyID:  "651fb40e44d0b172ae3ff0f9",
				employeeID: "651fb40e44d0b172ae3ff0f9",
				req:        &api.UpdateEmployeeDataJSONRequestBody{},
			},
			mockUpdate:   nil,
			mockFetchOne: nil,
			wantErr:      false,
			want:         "000000000000000000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &employeeUserAdmin{}
			mockEmployeeRepo := new(mocks.IEmployeeRepository)
			mockEmployeeRepo.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockUpdate)
			mockEmployeeRepo.On("FetchOne", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockFetchOne)
			e.EmployeeRepository = mockEmployeeRepo

			got, got1, err := e.UpdateEmployeeData(tt.args.ctx, tt.args.companyID, tt.args.employeeID, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateEmployeeData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateEmployeeData() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("UpdateEmployeeData() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_employeeUserAdmin_DeleteEmployeeByID(t *testing.T) {
	type fields struct {
		EmployeeRepository repositories.IEmployeeRepository
		CompanyRepository  repositories.ICompanyRepository
	}
	type args struct {
		ctx        context.Context
		employeeID string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		mockDelete error
	}{
		{
			name: "error object id from hex",
			args: args{
				ctx:        context.TODO(),
				employeeID: "1",
			},
			wantErr:    true,
			mockDelete: errors.New(internal.ErrorInternalServer.String()),
		},
		{
			name: "error delete employee",
			args: args{
				ctx:        context.TODO(),
				employeeID: "651fb40e44d0b172ae3ff0f9",
			},
			wantErr:    true,
			mockDelete: errors.New(internal.ErrorInternalServer.String()),
		},
		{
			name: "error delete employee no document",
			args: args{
				ctx:        context.TODO(),
				employeeID: "651fb40e44d0b172ae3ff0f9",
			},
			wantErr:    true,
			mockDelete: mongo.ErrNoDocuments,
		},
		{
			name: "Happy Flow",
			args: args{
				ctx:        context.TODO(),
				employeeID: "651fb40e44d0b172ae3ff0f9",
			},
			wantErr:    false,
			mockDelete: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &employeeUserAdmin{}
			mockEmployeeRepo := new(mocks.IEmployeeRepository)
			mockEmployeeRepo.On("Delete", mock.Anything, mock.Anything).Return(tt.mockDelete)
			e.EmployeeRepository = mockEmployeeRepo

			if err := e.DeleteEmployeeByID(tt.args.ctx, tt.args.employeeID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteEmployeeByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewEmployeeUseCase(t *testing.T) {
	type args struct {
		employeeRepository repositories.IEmployeeRepository
		companyRepository  repositories.ICompanyRepository
	}
	tests := []struct {
		name string
		args args
		want IEmployeeUseCase
	}{
		{
			name: "initial employee use case",
			args: args{},
			want: &employeeUserAdmin{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmployeeUseCase(tt.args.employeeRepository, tt.args.companyRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmployeeUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
