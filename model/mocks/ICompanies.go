// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ICompanies is an autogenerated mock type for the ICompanies type
type ICompanies struct {
	mock.Mock
}

// GetTableName provides a mock function with given fields:
func (_m *ICompanies) GetTableName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
