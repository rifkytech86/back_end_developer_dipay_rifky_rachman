package usecase

import (
	"context"
	"errors"
	"github.com/dipay/api"
	"github.com/dipay/internal"
	"github.com/dipay/internal/jwt"
	mockJWT "github.com/dipay/internal/jwt/mocks"
	"github.com/dipay/repositories"
	"github.com/dipay/repositories/mocks"
	"reflect"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
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
		name      string
		fields    fields
		args      args
		mockFetch error
		mockJwt   mockJwt
		want      string
		wantErr   bool
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
			name: "Happy flow",
			args: args{
				ctx: context.TODO(),
				req: &api.LoginJSONBody{},
			},
			mockFetch: nil,
			wantErr:   false,
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
			if got := NewUseCaseUserAdmin(tt.args.userAdminRepository, tt.args.jwt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUseCaseUserAdmin() = %v, want %v", got, tt.want)
			}
		})
	}
}
