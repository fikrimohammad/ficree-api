// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// All provides a mock function with given fields: params
func (_m *UserService) All(params map[string]interface{}) ([]map[string]interface{}, error) {
	ret := _m.Called(params)

	var r0 []map[string]interface{}
	if rf, ok := ret.Get(0).(func(map[string]interface{}) []map[string]interface{}); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]interface{}) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: params
func (_m *UserService) Create(params map[string]interface{}) (map[string]interface{}, error) {
	ret := _m.Called(params)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(map[string]interface{}) map[string]interface{}); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]interface{}) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Destroy provides a mock function with given fields: id
func (_m *UserService) Destroy(id int) (map[string]interface{}, error) {
	ret := _m.Called(id)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(int) map[string]interface{}); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Show provides a mock function with given fields: id
func (_m *UserService) Show(id int) (map[string]interface{}, error) {
	ret := _m.Called(id)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(int) map[string]interface{}); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, params
func (_m *UserService) Update(id int, params map[string]interface{}) (map[string]interface{}, error) {
	ret := _m.Called(id, params)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(int, map[string]interface{}) map[string]interface{}); ok {
		r0 = rf(id, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, map[string]interface{}) error); ok {
		r1 = rf(id, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}