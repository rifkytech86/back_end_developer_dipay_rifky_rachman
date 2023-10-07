package usecase

import (
	"context"
	"errors"
	"github.com/dipay/api"
	"github.com/dipay/internal"
	"github.com/dipay/internal/jwt"
	mockJWT "github.com/dipay/internal/jwt/mocks"
	"github.com/dipay/model"
	mockModel "github.com/dipay/model/mocks"
	"github.com/dipay/repositories"
	"github.com/dipay/repositories/mocks"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"

	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_useCaseUserAdmin_Login(t *testing.T) {
	type fields struct {
		UserAdminRepository repositories.IUserAdminRepository
		JWT                 jwt.IJWTRSAToken
	}
	type mockJwt struct {
		token string
		err   error
	}
	type args struct {
		ctx context.Context
		req *api.LoginJSONBody
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		mockFetch          error
		mockJwt            mockJwt
		mockModelUserAdmin error
		want               string
		wantErr            bool
	}{
		{
			name: "error fetch data",
			args: args{
				ctx: context.TODO(),
				req: &api.LoginJSONBody{},
			},
			mockFetch: errors.New(internal.ErrorInternalServer.String()),
			wantErr:   true,
		},
		{
			name: "error fetch data with ",
			args: args{
				ctx: context.TODO(),
				req: &api.LoginJSONBody{},
			},
			mockFetch: mongo.ErrNoDocuments,
			wantErr:   true,
		},
		{
			name: "error generate jwt",
			args: args{
				ctx: context.TODO(),
				req: &api.LoginJSONBody{},
			},
			mockFetch: nil,
			wantErr:   true,
			mockJwt: mockJwt{
				token: "",
				err:   errors.New(internal.ErrorInternalServer.String()),
			},
		},
		{
			name: "error compare password",
			args: args{
				ctx: context.TODO(),
				req: &api.LoginJSONBody{},
			},
			mockFetch:          nil,
			wantErr:            true,
			mockModelUserAdmin: errors.New(internal.ErrorInvalidRequest.String()),
			mockJwt: mockJwt{
				token: "1234",
				err:   nil,
			},
			want: "",
		},
		{
			name: "happy flow",
			args: args{
				ctx: context.TODO(),
				req: &api.LoginJSONBody{},
			},
			mockFetch:          nil,
			mockModelUserAdmin: nil,
			wantErr:            false,
			mockJwt: mockJwt{
				token: "1234",
				err:   nil,
			},
			want: "1234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &useCaseUserAdmin{}
			mockUserAdminRepo := new(mocks.IUserAdminRepository)
			mockUserAdminRepo.On("Fetch", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockFetch)
			u.UserAdminRepository = mockUserAdminRepo
			mockModelUserAdmin := new(mockModel.IUserAdmin)
			mockModelUserAdmin.On("IsValidPassword", mock.Anything, mock.Anything).Return(tt.mockModelUserAdmin)
			u.ModelUserAdmin = mockModelUserAdmin

			mockUserJWT := new(mockJWT.IJWTRSAToken)
			mockUserJWT.On("GenerateToken", mock.Anything, mock.Anything).Return(tt.mockJwt.token, tt.mockJwt.err)
			u.JWT = mockUserJWT

			got, err := u.Login(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUseCaseUserAdmin(t *testing.T) {
	type args struct {
		userAdminRepository repositories.IUserAdminRepository
		jwt                 jwt.IJWTRSAToken
		modelUserAdmin      model.IUserAdmin
	}
	tests := []struct {
		name string
		args args
		want IUseCaseUserAdmin
	}{
		{
			name: "initial use case user admin",
			args: args{},
			want: &useCaseUserAdmin{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUseCaseUserAdmin(tt.args.userAdminRepository, tt.args.modelUserAdmin, tt.args.jwt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUseCaseUserAdmin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCaseUserAdmin_Register(t *testing.T) {
	objectID := primitive.NewObjectID()
	type fields struct {
		UserAdminRepository repositories.IUserAdminRepository
		JWT                 jwt.IJWTRSAToken
		ModelUserAdmin      model.IUserAdmin
	}
	type mockJwt struct {
		token string
		err   error
	}
	type mockModelUserAdmin struct {
		err       error
		encrypted string
	}
	type mockFetch struct {
		id  *primitive.ObjectID
		err error
	}
	type args struct {
		ctx context.Context
		req *api.LoginJSONBody
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		mockFetch          mockFetch
		mockJwt            mockJwt
		mockModelUserAdmin mockModelUserAdmin
		want               string
		wantErr            bool
	}{
		{
			name: "error hassed",
			args: args{
				req: &api.LoginJSONBody{
					Password: "faield",
				},
			},
			mockModelUserAdmin: mockModelUserAdmin{
				err: errors.New(internal.ErrorInternalServer.String()),
			},
			wantErr: true,
		},
		{
			name: "error invalid password",
			args: args{
				req: &api.LoginJSONBody{
					Password: "faield",
				},
			},
			mockFetch: mockFetch{
				err: errors.New(internal.ErrorInternalServer.String()),
			},
			mockModelUserAdmin: mockModelUserAdmin{
				err:       nil,
				encrypted: "xxx",
			},
			wantErr: true,
		},
		{
			name: "error no document",
			args: args{
				req: &api.LoginJSONBody{
					Password: "faield",
				},
			},
			mockFetch: mockFetch{
				err: mongo.ErrNoDocuments,
			},
			mockModelUserAdmin: mockModelUserAdmin{
				err:       nil,
				encrypted: "xxx",
			},
			wantErr: true,
		},
		{
			name: "error generate token",
			args: args{
				req: &api.LoginJSONBody{
					Password: "faield",
				},
			},
			mockFetch: mockFetch{
				err: nil,
				id:  &objectID,
			},
			mockJwt: mockJwt{
				err: errors.New(internal.ErrorInternalServer.String()),
			},
			mockModelUserAdmin: mockModelUserAdmin{
				err:       nil,
				encrypted: "xxx",
			},
			wantErr: true,
		},
		{
			name: "Happy FLow",
			args: args{
				req: &api.LoginJSONBody{
					Password: "faield",
				},
			},
			mockFetch: mockFetch{
				err: nil,
				id:  &objectID,
			},
			mockJwt: mockJwt{
				err:   nil,
				token: "xxx",
			},
			mockModelUserAdmin: mockModelUserAdmin{
				err:       nil,
				encrypted: "xxx",
			},
			wantErr: false,
			want:    "xxx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &useCaseUserAdmin{}
			mockUserAdminRepo := new(mocks.IUserAdminRepository)
			mockUserAdminRepo.On("Create", mock.Anything, mock.Anything).Return(tt.mockFetch.id, tt.mockFetch.err)
			u.UserAdminRepository = mockUserAdminRepo

			mockModelUserAdmin := new(mockModel.IUserAdmin)
			mockModelUserAdmin.On("EncryptedPassword", mock.Anything).Return(tt.mockModelUserAdmin.encrypted, tt.mockModelUserAdmin.err)
			u.ModelUserAdmin = mockModelUserAdmin

			mockUserJWT := new(mockJWT.IJWTRSAToken)
			mockUserJWT.On("GenerateToken", mock.Anything, mock.Anything).Return(tt.mockJwt.token, tt.mockJwt.err)
			u.JWT = mockUserJWT

			got, err := u.Register(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}
