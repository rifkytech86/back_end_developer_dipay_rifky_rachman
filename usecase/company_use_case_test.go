package usecase

import (
	"context"
	"errors"
	"github.com/dipay/api"
	"github.com/dipay/internal"
	"github.com/dipay/internal/jwt"
	"github.com/dipay/model"
	"github.com/dipay/repositories"
	"github.com/dipay/repositories/mocks"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func Test_companyUseCase_AddCompany(t *testing.T) {
	insertedID := primitive.NewObjectID()
	type fields struct {
		CompanyRepository repositories.ICompanyRepository
		JWT               jwt.IJWTRSAToken
	}
	type mockCreate struct {
		id  *primitive.ObjectID
		err error
	}
	type args struct {
		ctx context.Context
		req *api.AddCompanyJSONBody
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       string
		mockCreate mockCreate
		wantErr    bool
	}{
		{
			name: "error add company",
			args: args{
				ctx: context.TODO(),
				req: &api.AddCompanyJSONBody{
					CompanyName:     "tester",
					TelephoneNumber: "+6290809809",
					Address:         "address",
				},
			},
			mockCreate: mockCreate{
				id:  &insertedID,
				err: errors.New(internal.ErrorInvalidRequest.String()),
			},
			wantErr: true,
		},
		{
			name: "Happy Flow add company",
			args: args{
				ctx: context.TODO(),
				req: &api.AddCompanyJSONBody{
					CompanyName:     "tester",
					TelephoneNumber: "+6290809809",
					Address:         "address",
				},
			},
			mockCreate: mockCreate{
				id:  &insertedID,
				err: nil,
			},
			wantErr: false,
			want:    insertedID.Hex(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &companyUseCase{}
			mockRepo := new(mocks.ICompanyRepository)
			mockRepo.On("Create", mock.Anything, mock.Anything).Return(tt.mockCreate.id, tt.mockCreate.err)
			u.CompanyRepository = mockRepo

			got, err := u.AddCompany(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddCompany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddCompany() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_companyUseCase_GetCompany(t *testing.T) {
	type fields struct {
		CompanyRepository repositories.ICompanyRepository
	}
	type mockGetCompany struct {
		data []*model.Companies
		err  error
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		want           []*model.Companies
		wantErr        bool
		mockGetCompany mockGetCompany
	}{
		{
			name: "error fetch",
			args: args{
				ctx: context.TODO(),
			},
			mockGetCompany: mockGetCompany{
				data: []*model.Companies{},
				err:  errors.New(internal.ErrorInvalidRequest.String()),
			},
			wantErr: true,
		},
		{
			name: "error fetch, but expected error no document",
			args: args{
				ctx: context.TODO(),
			},
			mockGetCompany: mockGetCompany{
				data: []*model.Companies{},
				err:  mongo.ErrNoDocuments,
			},
			wantErr: false,
			want:    []*model.Companies{},
		},
		{
			name: "happy flow fetch",
			args: args{
				ctx: context.TODO(),
			},
			mockGetCompany: mockGetCompany{
				data: []*model.Companies{},
				err:  nil,
			},
			wantErr: false,
			want:    []*model.Companies{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &companyUseCase{}
			mockRepo := new(mocks.ICompanyRepository)
			mockRepo.On("Fetch", mock.Anything, mock.Anything).Return(tt.mockGetCompany.data, tt.mockGetCompany.err)
			u.CompanyRepository = mockRepo

			got, err := u.GetCompany(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCompany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCompany() got = %v, want %v", got, tt.want)
			}
		})
	}
}
