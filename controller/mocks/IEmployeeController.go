// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import echo "github.com/labstack/echo/v4"
import mock "github.com/stretchr/testify/mock"

// IEmployeeController is an autogenerated mock type for the IEmployeeController type
type IEmployeeController struct {
	mock.Mock
}

// AddEmployee provides a mock function with given fields: c, companyId
func (_m *IEmployeeController) AddEmployee(c echo.Context, companyId string) error {
	ret := _m.Called(c, companyId)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, string) error); ok {
		r0 = rf(c, companyId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteEmployeeByID provides a mock function with given fields: ctx, employeeId
func (_m *IEmployeeController) DeleteEmployeeByID(ctx echo.Context, employeeId string) error {
	ret := _m.Called(ctx, employeeId)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, string) error); ok {
		r0 = rf(ctx, employeeId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetEmployeeByCompanyID provides a mock function with given fields: ctx, companyID
func (_m *IEmployeeController) GetEmployeeByCompanyID(ctx echo.Context, companyID string) error {
	ret := _m.Called(ctx, companyID)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, string) error); ok {
		r0 = rf(ctx, companyID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetEmployeeByID provides a mock function with given fields: c, employeeID
func (_m *IEmployeeController) GetEmployeeByID(c echo.Context, employeeID string) error {
	ret := _m.Called(c, employeeID)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, string) error); ok {
		r0 = rf(c, employeeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEmployeeData provides a mock function with given fields: ctx, companyId, employeeId
func (_m *IEmployeeController) UpdateEmployeeData(ctx echo.Context, companyId string, employeeId string) error {
	ret := _m.Called(ctx, companyId, employeeId)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, string, string) error); ok {
		r0 = rf(ctx, companyId, employeeId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
