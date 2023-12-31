// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import api "github.com/dipay/api"

import echo "github.com/labstack/echo/v4"
import mock "github.com/stretchr/testify/mock"

// IUserAdminController is an autogenerated mock type for the IUserAdminController type
type IUserAdminController struct {
	mock.Mock
}

// Hello provides a mock function with given fields: ctx, param
func (_m *IUserAdminController) Hello(ctx echo.Context, param api.HelloParams) error {
	ret := _m.Called(ctx, param)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, api.HelloParams) error); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: ctx
func (_m *IUserAdminController) Login(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
