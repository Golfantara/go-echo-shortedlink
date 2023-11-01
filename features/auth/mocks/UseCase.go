// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	dtos "shortlink/features/auth/dtos"

	mock "github.com/stretchr/testify/mock"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Create provides a mock function with given fields: newUsers
func (_m *UseCase) Create(newUsers dtos.InputUsers) *dtos.ResRegister {
	ret := _m.Called(newUsers)

	var r0 *dtos.ResRegister
	if rf, ok := ret.Get(0).(func(dtos.InputUsers) *dtos.ResRegister); ok {
		r0 = rf(newUsers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.ResRegister)
		}
	}

	return r0
}

// FindAll provides a mock function with given fields: page, size
func (_m *UseCase) FindAll(page int, size int) []dtos.ResUsers {
	ret := _m.Called(page, size)

	var r0 []dtos.ResUsers
	if rf, ok := ret.Get(0).(func(int, int) []dtos.ResUsers); ok {
		r0 = rf(page, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dtos.ResUsers)
		}
	}

	return r0
}

// FindByID provides a mock function with given fields: userID
func (_m *UseCase) FindByID(userID int) *dtos.ResUsers {
	ret := _m.Called(userID)

	var r0 *dtos.ResUsers
	if rf, ok := ret.Get(0).(func(int) *dtos.ResUsers); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.ResUsers)
		}
	}

	return r0
}

// Login provides a mock function with given fields: email, password
func (_m *UseCase) Login(email string, password string) (*dtos.ResLogin, error) {
	ret := _m.Called(email, password)

	var r0 *dtos.ResLogin
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*dtos.ResLogin, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) *dtos.ResLogin); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.ResLogin)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Modify provides a mock function with given fields: userData, userID
func (_m *UseCase) Modify(userData dtos.InputUsers, userID int) bool {
	ret := _m.Called(userData, userID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(dtos.InputUsers, int) bool); ok {
		r0 = rf(userData, userID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Remove provides a mock function with given fields: userID
func (_m *UseCase) Remove(userID int) bool {
	ret := _m.Called(userID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}