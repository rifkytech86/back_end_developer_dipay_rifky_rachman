// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import validations "github.com/dipay/internal/validations"
import validator "github.com/go-playground/validator/v10"

// IValidator is an autogenerated mock type for the IValidator type
type IValidator struct {
	mock.Mock
}

// RegisterValidation provides a mock function with given fields: tag, fn, callValidationEvenIfNull
func (_m *IValidator) RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	_va := make([]interface{}, len(callValidationEvenIfNull))
	for _i := range callValidationEvenIfNull {
		_va[_i] = callValidationEvenIfNull[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, tag, fn)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, validator.Func, ...bool) error); ok {
		r0 = rf(tag, fn, callValidationEvenIfNull...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Struct provides a mock function with given fields: s
func (_m *IValidator) Struct(s interface{}) []validations.ValidationError {
	ret := _m.Called(s)

	var r0 []validations.ValidationError
	if rf, ok := ret.Get(0).(func(interface{}) []validations.ValidationError); ok {
		r0 = rf(s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]validations.ValidationError)
		}
	}

	return r0
}
