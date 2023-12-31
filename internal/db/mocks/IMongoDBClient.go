// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import mongo "go.mongodb.org/mongo-driver/mongo"

// IMongoDBClient is an autogenerated mock type for the IMongoDBClient type
type IMongoDBClient struct {
	mock.Mock
}

// Disconnection provides a mock function with given fields: client
func (_m *IMongoDBClient) Disconnection(client *mongo.Client) error {
	ret := _m.Called(client)

	var r0 error
	if rf, ok := ret.Get(0).(func(*mongo.Client) error); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitConnection provides a mock function with given fields:
func (_m *IMongoDBClient) InitConnection() (*mongo.Client, error) {
	ret := _m.Called()

	var r0 *mongo.Client
	if rf, ok := ret.Get(0).(func() *mongo.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PingConnection provides a mock function with given fields: client
func (_m *IMongoDBClient) PingConnection(client *mongo.Client) error {
	ret := _m.Called(client)

	var r0 error
	if rf, ok := ret.Get(0).(func(*mongo.Client) error); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDatabase provides a mock function with given fields: client, databaseName
func (_m *IMongoDBClient) SetDatabase(client *mongo.Client, databaseName string) *mongo.Database {
	ret := _m.Called(client, databaseName)

	var r0 *mongo.Database
	if rf, ok := ret.Get(0).(func(*mongo.Client, string) *mongo.Database); ok {
		r0 = rf(client, databaseName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Database)
		}
	}

	return r0
}
