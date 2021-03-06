// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/fikrimohammad/ficree-api/domain"
	mock "github.com/stretchr/testify/mock"
)

// FileRepository is an autogenerated mock type for the FileRepository type
type FileRepository struct {
	mock.Mock
}

// FindByURI provides a mock function with given fields: uri
func (_m *FileRepository) FindByURI(uri string) (domain.File, error) {
	ret := _m.Called(uri)

	var r0 domain.File
	if rf, ok := ret.Get(0).(func(string) domain.File); ok {
		r0 = rf(uri)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uri)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
