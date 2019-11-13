// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/alunegov/k3archive/models"
	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// DeleteFile provides a mock function with given fields: id
func (_m *Usecase) DeleteFile(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDir provides a mock function with given fields:
func (_m *Usecase) GetDir() ([]*models.FileInfo, error) {
	ret := _m.Called()

	var r0 []*models.FileInfo
	if rf, ok := ret.Get(0).(func() []*models.FileInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFile provides a mock function with given fields: id
func (_m *Usecase) GetFile(id string) (*models.FileInfo, error) {
	ret := _m.Called(id)

	var r0 *models.FileInfo
	if rf, ok := ret.Get(0).(func(string) *models.FileInfo); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}