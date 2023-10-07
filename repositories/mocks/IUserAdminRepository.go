// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import primitive "go.mongodb.org/mongo-driver/bson/primitive"

// IUserAdminRepository is an autogenerated mock type for the IUserAdminRepository type
type IUserAdminRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, model
func (_m *IUserAdminRepository) Create(ctx context.Context, model interface{}) (*primitive.ObjectID, error) {
	ret := _m.Called(ctx, model)

	var r0 *primitive.ObjectID
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) *primitive.ObjectID); ok {
		r0 = rf(ctx, model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*primitive.ObjectID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, filter, result
func (_m *IUserAdminRepository) Fetch(ctx context.Context, filter interface{}, result interface{}) error {
	ret := _m.Called(ctx, filter, result)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}) error); ok {
		r0 = rf(ctx, filter, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
