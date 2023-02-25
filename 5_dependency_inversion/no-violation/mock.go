// Code generated by mockery v2.13.1. DO NOT EDIT.

package dependencyinversion

import (
	mock "github.com/stretchr/testify/mock"
)

// UserRepositoryInterface is an autogenerated mock type for the UserRepositoryInterface type
type UserRepositoryMock struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: id
func (_m *UserRepositoryMock) GetUserByID(id uint) (*User, error) {
	ret := _m.Called(id)

	var r0 *User
	if rf, ok := ret.Get(0).(func(uint) *User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepositoryInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepositoryInterface creates a new instance of UserRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepositoryMock(t mockConstructorTestingTNewUserRepositoryInterface) *UserRepositoryMock {
	mock := &UserRepositoryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
