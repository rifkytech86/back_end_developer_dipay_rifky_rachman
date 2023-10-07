// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IUserAdmin is an autogenerated mock type for the IUserAdmin type
type IUserAdmin struct {
	mock.Mock
}

// EncryptedPassword provides a mock function with given fields: userAdminPassword
func (_m *IUserAdmin) EncryptedPassword(userAdminPassword string) (string, error) {
	ret := _m.Called(userAdminPassword)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(userAdminPassword)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userAdminPassword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTableName provides a mock function with given fields:
func (_m *IUserAdmin) GetTableName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsValidPassword provides a mock function with given fields: userAdminPassword, userReqPassword
func (_m *IUserAdmin) IsValidPassword(userAdminPassword string, userReqPassword string) error {
	ret := _m.Called(userAdminPassword, userReqPassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userAdminPassword, userReqPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}