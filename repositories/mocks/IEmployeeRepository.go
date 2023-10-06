// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/dipay/model"
import primitive "go.mongodb.org/mongo-driver/bson/primitive"

// IEmployeeRepository is an autogenerated mock type for the IEmployeeRepository type
type IEmployeeRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *IEmployeeRepository) Create(ctx context.Context, _a1 interface{}) (*primitive.ObjectID, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *primitive.ObjectID
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) *primitive.ObjectID); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*primitive.ObjectID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, filter
func (_m *IEmployeeRepository) Delete(ctx context.Context, filter interface{}) error {
	ret := _m.Called(ctx, filter)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, filter
func (_m *IEmployeeRepository) Fetch(ctx context.Context, filter interface{}) ([]*model.Employees, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*model.Employees
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) []*model.Employees); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Employees)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchOne provides a mock function with given fields: ctx, filter, result
func (_m *IEmployeeRepository) FetchOne(ctx context.Context, filter interface{}, result interface{}) error {
	ret := _m.Called(ctx, filter, result)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}) error); ok {
		r0 = rf(ctx, filter, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, filter, update
func (_m *IEmployeeRepository) Update(ctx context.Context, filter interface{}, update interface{}) error {
	ret := _m.Called(ctx, filter, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}) error); ok {
		r0 = rf(ctx, filter, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
