// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	goly "shortlink/features/goly"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// DeleteByID provides a mock function with given fields: golyID
func (_m *Repository) DeleteByID(golyID int) int64 {
	ret := _m.Called(golyID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int) int64); ok {
		r0 = rf(golyID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// FindByGolyUrl provides a mock function with given fields: url
func (_m *Repository) FindByGolyUrl(url string) (goly.Goly, error) {
	ret := _m.Called(url)

	var r0 goly.Goly
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (goly.Goly, error)); ok {
		return rf(url)
	}
	if rf, ok := ret.Get(0).(func(string) goly.Goly); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Get(0).(goly.Goly)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: newGoly
func (_m *Repository) Insert(newGoly *goly.Goly) *goly.Goly {
	ret := _m.Called(newGoly)

	var r0 *goly.Goly
	if rf, ok := ret.Get(0).(func(*goly.Goly) *goly.Goly); ok {
		r0 = rf(newGoly)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*goly.Goly)
		}
	}

	return r0
}

// Paginate provides a mock function with given fields: page, size
func (_m *Repository) Paginate(page int, size int) []goly.Goly {
	ret := _m.Called(page, size)

	var r0 []goly.Goly
	if rf, ok := ret.Get(0).(func(int, int) []goly.Goly); ok {
		r0 = rf(page, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]goly.Goly)
		}
	}

	return r0
}

// PaginateIP provides a mock function with given fields: page, size
func (_m *Repository) PaginateIP(page int, size int) []goly.IPAdresses {
	ret := _m.Called(page, size)

	var r0 []goly.IPAdresses
	if rf, ok := ret.Get(0).(func(int, int) []goly.IPAdresses); ok {
		r0 = rf(page, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]goly.IPAdresses)
		}
	}

	return r0
}

// SearchingGoly provides a mock function with given fields: short
func (_m *Repository) SearchingGoly(short string) ([]goly.Goly, error) {
	ret := _m.Called(short)

	var r0 []goly.Goly
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]goly.Goly, error)); ok {
		return rf(short)
	}
	if rf, ok := ret.Get(0).(func(string) []goly.Goly); ok {
		r0 = rf(short)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]goly.Goly)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(short)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectByID provides a mock function with given fields: golyID
func (_m *Repository) SelectByID(golyID int) *goly.Goly {
	ret := _m.Called(golyID)

	var r0 *goly.Goly
	if rf, ok := ret.Get(0).(func(int) *goly.Goly); ok {
		r0 = rf(golyID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*goly.Goly)
		}
	}

	return r0
}

// StoreIPForGoly provides a mock function with given fields: golyID, ip
func (_m *Repository) StoreIPForGoly(golyID uint64, ip string) error {
	ret := _m.Called(golyID, ip)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, string) error); ok {
		r0 = rf(golyID, ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0
func (_m *Repository) Update(_a0 goly.Goly) int64 {
	ret := _m.Called(_a0)

	var r0 int64
	if rf, ok := ret.Get(0).(func(goly.Goly) int64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// UpdateButton provides a mock function with given fields: _a0
func (_m *Repository) UpdateButton(_a0 goly.Goly) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(goly.Goly) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
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
