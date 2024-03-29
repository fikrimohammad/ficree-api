// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/fikrimohammad/ficree-api/domain"
	mock "github.com/stretchr/testify/mock"
)

// FileService is an autogenerated mock type for the FileService type
type FileService struct {
	mock.Mock
}

// GetFileURL provides a mock function with given fields: params
func (_m *FileService) GetFileURL(params domain.GenerateFileURLInput) (*domain.FileOutput, error) {
	ret := _m.Called(params)

	var r0 *domain.FileOutput
	if rf, ok := ret.Get(0).(func(domain.GenerateFileURLInput) *domain.FileOutput); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.FileOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.GenerateFileURLInput) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
