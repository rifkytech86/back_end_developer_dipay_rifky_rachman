package repositories

import (
	"context"
	"errors"
	"github.com/dipay/internal"
	"github.com/dipay/internal/db"
	dbMock "github.com/dipay/internal/db/mocks"
	"github.com/dipay/model"
	"github.com/dipay/model/mocks"
	"github.com/dipay/pkg/httpClient"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func Test_employeeRepository_Fetch(t *testing.T) {
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
			want:          []*model.Companies{},
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
			c := &employeeRepository{}
			mockModelDB := new(mocks.ICompanies)
			mockModelDB.On("GetTableName", mock.Anything).Return(tt.mockTableName)
			c.EmployeeModel = mockModelDB
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

			_, err := c.Fetch(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_employeeRepository_FetchOne(t *testing.T) {
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
			c := &employeeRepository{}
			mockModelDB := new(mocks.ICompanies)
			mockModelDB.On("GetTableName", mock.Anything).Return(tt.mockTableName)
			c.EmployeeModel = mockModelDB
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

func Test_employeeRepository_Create(t *testing.T) {
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
			c := &employeeRepository{}
			mockModelDB := new(mocks.ICompanies)
			mockModelDB.On("GetTableName", mock.Anything).Return(tt.mockTableName)
			c.EmployeeModel = mockModelDB
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

func Test_employeeRepository_Update(t *testing.T) {
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
			c := &employeeRepository{}
			mockModelDB := new(mocks.ICompanies)
			mockModelDB.On("GetTableName", mock.Anything).Return(tt.mockTableName)
			c.EmployeeModel = mockModelDB
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

func Test_employeeRepository_Delete(t *testing.T) {
	type fields struct {
		MongoDatabase db.Database
		EmployeeModel model.IEmployees
	}
	type mockDeleteOne struct {
		id  int64
		err error
	}
	type args struct {
		ctx    context.Context
		filter interface{}
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantErr       bool
		mockTableName string
		mockDeleteOne mockDeleteOne
	}{
		{
			name:          "error delete",
			mockTableName: "company",
			args: args{
				ctx:    context.TODO(),
				filter: bson.M{},
			},
			mockDeleteOne: mockDeleteOne{
				err: errors.New(internal.ErrorInternalServer.String()),
				id:  0,
			},
			wantErr: true,
		},
		{
			name:          "Happy delete",
			mockTableName: "company",
			args: args{
				ctx:    context.TODO(),
				filter: bson.M{},
			},
			mockDeleteOne: mockDeleteOne{
				err: nil,
				id:  1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &employeeRepository{}
			mockModelDB := new(mocks.ICompanies)
			mockModelDB.On("GetTableName", mock.Anything).Return(tt.mockTableName)
			c.EmployeeModel = mockModelDB
			mockCollection := new(dbMock.Collection)
			mockCollection.On("DeleteOne", context.TODO(), mock.Anything).Return(tt.mockDeleteOne.id, tt.mockDeleteOne.err)
			mockDB := new(dbMock.Database)
			mockDB.On("Collection", mock.Anything).Return(mockCollection)
			c.MongoDatabase = mockDB

			if err := c.Delete(tt.args.ctx, tt.args.filter); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewEmployeeRepository(t *testing.T) {
	type args struct {
		mongoDatabase db.Database
		employeeModel model.IEmployees
		clientHttp    httpClient.IClientHttp
		hostEmail     string
	}
	tests := []struct {
		name string
		args args
		want IEmployeeRepository
	}{
		{
			name: "initial employee",
			args: args{},
			want: &employeeRepository{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmployeeRepository(tt.args.mongoDatabase, tt.args.employeeModel, tt.args.clientHttp, tt.args.hostEmail); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmployeeRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
