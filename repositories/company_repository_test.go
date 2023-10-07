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

func Test_companyRepository_FetchOne(t *testing.T) {
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
			c := &companyRepository{}
			mockModelDB := new(mocks.ICompanies)
			mockModelDB.On("GetTableName", mock.Anything).Return(tt.mockTableName)
			c.CompaniesModel = mockModelDB
			mockSingleResult := new(dbMock.SingleResult)
			mockSingleResult.On("Decode", mock.Anything).Return(tt.mockDecode)

			mockCollection := new(dbMock.Collection)
			mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult)

			mockDB := new(dbMock.Database)
			mockDB.On("Collection", mock.Anything).Return(mockCollection)

			c.MongoDatabase = mockDB

			if err := c.FetchOne(tt.args.ctx, tt.args.filter, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("FetchOne() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_companyRepository_Fetch(t *testing.T) {
	type fields struct {
		MongoDatabase  db.Database
		CompaniesModel model.ICompanies
	}
	type args struct {
		ctx    context.Context
		filter interface{}
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          []*model.Companies
		mockTableName string
		mockCursor    error
		mockFind      error
		mockNext      bool
		mockDecode    error
		wantErr       bool
	}{
		{
			name: "Error Find Collection",
			args: args{
				ctx:    context.TODO(),
				filter: bson.M{},
			},
			mockTableName: "company",
			mockCursor:    errors.New(internal.ErrorInternalServer.String()),
			mockFind:      errors.New(internal.ErrorInternalServer.String()),
			wantErr:       true,
		},
		{
			name: "Error Decode",
			args: args{
				ctx:    context.TODO(),
				filter: bson.M{},
			},
			mockTableName: "company",
			mockCursor:    errors.New(internal.ErrorInternalServer.String()),
			mockFind:      nil,
			mockNext:      true,
			mockDecode:    errors.New(internal.ErrorInternalServer.String()),
			wantErr:       true,
		},
		{
			name: "Happy Decode",
			args: args{
				ctx:    context.TODO(),
				filter: bson.M{},
			},
			mockTableName: "company",
			mockCursor:    errors.New(internal.ErrorInternalServer.String()),
			mockFind:      nil,
			mockNext:      false,
			mockDecode:    nil,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &companyRepository{}
			mockModelDB := new(mocks.ICompanies)
			mockModelDB.On("GetTableName", mock.Anything).Return(tt.mockTableName)
			c.CompaniesModel = mockModelDB
			mockCursor := new(dbMock.Cursor)
			mockCursor.On("All", mock.Anything).Return(tt.mockCursor)
			mockCursor.On("Next", mock.Anything).Return(tt.mockNext)
			mockCursor.On("Close", mock.Anything).Return(tt.mockCursor)
			mockCursor.On("Decode", mock.Anything).Return(tt.mockDecode)
			mockCollection := new(dbMock.Collection)
			mockCollection.On("Find", context.TODO(), mock.AnythingOfType("primitive.M")).Return(mockCursor, tt.mockFind)
			mockDB := new(dbMock.Database)
			mockDB.On("Collection", mock.Anything).Return(mockCollection)
			c.MongoDatabase = mockDB

			got, err := c.Fetch(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fetch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_companyRepository_Create(t *testing.T) {
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
			c := &companyRepository{}
			mockModelDB := new(mocks.ICompanies)
			mockModelDB.On("GetTableName", mock.Anything).Return(tt.mockTableName)
			c.CompaniesModel = mockModelDB
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

func Test_companyRepository_Update(t *testing.T) {
	type fields struct {
		MongoDatabase  db.Database
		CompaniesModel model.ICompanies
	}
	type mockUpdateOne struct {
		err          error
		updateResult *mongo.UpdateResult
	}
	type args struct {
		ctx    context.Context
		filter interface{}
		update interface{}
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		mockTableName string
		mockUpdateOne mockUpdateOne
		wantErr       bool
	}{
		{
			name: "Error Update",
			args: args{
				ctx:    context.TODO(),
				filter: bson.M{},
			},
			mockTableName: "company",
			mockUpdateOne: mockUpdateOne{
				err:          errors.New(internal.ErrorInternalServer.String()),
				updateResult: nil,
			},
			wantErr: true,
		},
		{
			name: "Error Update no affected",
			args: args{
				ctx:    context.TODO(),
				filter: bson.M{},
			},
			mockTableName: "company",
			mockUpdateOne: mockUpdateOne{
				err: nil,
				updateResult: &mongo.UpdateResult{
					MatchedCount: 0,
				},
			},
			wantErr: true,
		},
		{
			name: "Happy Update",
			args: args{
				ctx:    context.TODO(),
				filter: bson.M{},
			},
			mockTableName: "company",
			mockUpdateOne: mockUpdateOne{
				err: nil,
				updateResult: &mongo.UpdateResult{
					MatchedCount:  1,
					ModifiedCount: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &companyRepository{}
			mockModelDB := new(mocks.ICompanies)
			mockModelDB.On("GetTableName", mock.Anything).Return(tt.mockTableName)
			c.CompaniesModel = mockModelDB
			mockCollection := new(dbMock.Collection)
			mockCollection.On("UpdateOne", context.TODO(), mock.Anything, mock.Anything).Return(tt.mockUpdateOne.updateResult, tt.mockUpdateOne.err)
			mockDB := new(dbMock.Database)
			mockDB.On("Collection", mock.Anything).Return(mockCollection)
			c.MongoDatabase = mockDB

			if err := c.Update(tt.args.ctx, tt.args.filter, tt.args.update); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCompanyRepository(t *testing.T) {
	type args struct {
		mongoDatabase  db.Database
		companiesModel model.ICompanies
	}
	tests := []struct {
		name string
		args args
		want ICompanyRepository
	}{
		{
			name: "initial Company Repository",
			args: args{},
			want: &companyRepository{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCompanyRepository(tt.args.mongoDatabase, tt.args.companiesModel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCompanyRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
