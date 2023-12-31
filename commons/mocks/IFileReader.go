// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IFileReader is an autogenerated mock type for the IFileReader type
type IFileReader struct {
	mock.Mock
}

// ReadFile provides a mock function with given fields: filename
func (_m *IFileReader) ReadFile(filename string) ([]byte, error) {
	ret := _m.Called(filename)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(filename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
