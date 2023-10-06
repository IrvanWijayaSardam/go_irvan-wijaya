// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	users "restEcho1/features/users"

	mock "github.com/stretchr/testify/mock"
)

// UserServiceInterface is an autogenerated mock type for the UserServiceInterface type
type UserServiceInterface struct {
	mock.Mock
}

// Login provides a mock function with given fields: hp, password
func (_m *UserServiceInterface) Login(hp string, password string) (*users.UserCredential, error) {
	ret := _m.Called(hp, password)

	var r0 *users.UserCredential
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*users.UserCredential, error)); ok {
		return rf(hp, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) *users.UserCredential); ok {
		r0 = rf(hp, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.UserCredential)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(hp, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: newData
func (_m *UserServiceInterface) Register(newData users.User) (*users.User, error) {
	ret := _m.Called(newData)

	var r0 *users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(users.User) (*users.User, error)); ok {
		return rf(newData)
	}
	if rf, ok := ret.Get(0).(func(users.User) *users.User); ok {
		r0 = rf(newData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	if rf, ok := ret.Get(1).(func(users.User) error); ok {
		r1 = rf(newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}


// NewUserServiceInterface creates a new instance of UserServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserServiceInterface {
	mock := &UserServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
