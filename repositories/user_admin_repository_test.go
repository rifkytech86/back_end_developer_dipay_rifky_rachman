package repositories

import (
	"context"
	"errors"
	"github.com/dipay/internal"
	"github.com/dipay/internal/db"
	dbMock "github.com/dipay/internal/db/mocks"
	"github.com/dipay/model"
	"github.com/dipay/model/mocks"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func Test_userAdminRepository_Fetch(t *testing.T) {
	var result interface{}
	type fields struct {
		MongoDatabase  *mongo.Database
		CompaniesModel model.ICompanies
	}
	type args struct {
		ctx    context.Context
		filter interface{}
		result interface{}
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		mockTableName string
		mockDecode    error
		wantErr       bool
	}{
		{
			name: "Error Fetch one",
			args: args{
				ctx:    context.TODO(),
				filter: bson.M{},
				result: result,
			},
			mockTableName: "company",
			mockDecode:    errors.New(internal.ErrorInternalServer.String()),
			wantErr:       true,
		},
		{
			name: "Happy Fetch one",
			args: args{
				ctx:    context.TODO(),
				filter: bson.M{},
				result: result,
			},
			mockTableName: "company",
			mockDecode:    nil,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &userAdminRepository{}
			mockModelDB := new(mocks.IUserAdmin)
			mockModelDB.On("GetTableName", mock.Anything).Return(tt.mockTableName)
			c.UserAdminModel = mockModelDB

			mockSingleResult := new(dbMock.SingleResult)
			mockSingleResult.On("Decode", mock.Anything).Return(tt.mockDecode)

			mockCollection := new(dbMock.Collection)
			mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult)

			mockDB := new(dbMock.Database)
			mockDB.On("Collection", mock.Anything).Return(mockCollection)
			c.MongoDatabase = mockDB
			if err := c.Fetch(tt.args.ctx, tt.args.filter, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("FetchOne() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewUserAdminRepository(t *testing.T) {
	type args struct {
		mongoDatabase  db.Database
		userAdminModel model.IUserAdmin
	}
	tests := []struct {
		name string
		args args
		want IUserAdminRepository
	}{
		{
			name: "initial user admin repository",
			args: args{},
			want: &userAdminRepository{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserAdminRepository(tt.args.mongoDatabase, tt.args.userAdminModel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserAdminRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userAdminRepository_Create(t *testing.T) {
	insertedID := primitive.NewObjectID()

	type fields struct {
		MongoDatabase  db.Database
		CompaniesModel model.ICompanies
	}
	type mockErrorInertOne struct {
		err error
		id  primitive.ObjectID
	}
	type args struct {
		ctx   context.Context
		model interface{}
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		want              *primitive.ObjectID
		wantErr           bool
		mockErrorInertOne mockErrorInertOne
		mockTableName     string
	}{
		{
			name: "Error InsertOne",
			args: args{
				ctx:   context.TODO(),
				model: model.Companies{},
			},
			mockTableName: "company",
			mockErrorInertOne: mockErrorInertOne{
				err: errors.New(internal.ErrorInternalServer.String()),
				id:  insertedID,
			},
			wantErr: true,
		},
		{
			name: "Happy InsertOne",
			args: args{
				ctx:   context.TODO(),
				model: model.Companies{},
			},
			mockTableName: "company",
			mockErrorInertOne: mockErrorInertOne{
				err: nil,
				id:  insertedID,
			},
			wantErr: false,
			want:    &insertedID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &userAdminRepository{}
			mockModelDB := new(mocks.IUserAdmin)
			mockModelDB.On("GetTableName", mock.Anything).Return(tt.mockTableName)
			c.UserAdminModel = mockModelDB
			mockCollection := new(dbMock.Collection)
			mockCollection.On("InsertOne", context.TODO(), mock.Anything).Return(tt.mockErrorInertOne.id, tt.mockErrorInertOne.err)
			mockDB := new(dbMock.Database)
			mockDB.On("Collection", mock.Anything).Return(mockCollection)
			c.MongoDatabase = mockDB

			got, err := c.Create(tt.args.ctx, tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
