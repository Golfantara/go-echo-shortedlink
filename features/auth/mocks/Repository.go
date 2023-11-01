// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	auth "shortlink/features/auth"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// DeleteByID provides a mock function with given fields: userID
func (_m *Repository) DeleteByID(userID int) int64 {
	ret := _m.Called(userID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int) int64); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Insert provides a mock function with given fields: newUsers
func (_m *Repository) Insert(newUsers *auth.Users) *auth.Users {
	ret := _m.Called(newUsers)

	var r0 *auth.Users
	if rf, ok := ret.Get(0).(func(*auth.Users) *auth.Users); ok {
		r0 = rf(newUsers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.Users)
		}
	}

	return r0
}

// Login provides a mock function with given fields: email, password
func (_m *Repository) Login(email string, password string) (*auth.Users, error) {
	ret := _m.Called(email, password)

	var r0 *auth.Users
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*auth.Users, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) *auth.Users); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.Users)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Paginate provides a mock function with given fields: page, size
func (_m *Repository) Paginate(page int, size int) []auth.Users {
	ret := _m.Called(page, size)

	var r0 []auth.Users
	if rf, ok := ret.Get(0).(func(int, int) []auth.Users); ok {
		r0 = rf(page, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]auth.Users)
		}
	}

	return r0
}

// SelectByID provides a mock function with given fields: userID
func (_m *Repository) SelectByID(userID int) *auth.Users {
	ret := _m.Called(userID)

	var r0 *auth.Users
	if rf, ok := ret.Get(0).(func(int) *auth.Users); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.Users)
		}
	}

	return r0
}

// Update provides a mock function with given fields: user
func (_m *Repository) Update(user auth.Users) int64 {
	ret := _m.Called(user)

	var r0 int64
	if rf, ok := ret.Get(0).(func(auth.Users) int64); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
