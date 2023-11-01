// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	donate "shortlink/features/donate"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CheckTransaction provides a mock function with given fields: orderID
func (_m *Repository) CheckTransaction(orderID string) (donate.Status, error) {
	ret := _m.Called(orderID)

	var r0 donate.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (donate.Status, error)); ok {
		return rf(orderID)
	}
	if rf, ok := ret.Get(0).(func(string) donate.Status); ok {
		r0 = rf(orderID)
	} else {
		r0 = ret.Get(0).(donate.Status)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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

// Insert provides a mock function with given fields: newData
func (_m *Repository) Insert(newData *donate.Transaction) *donate.Transaction {
	ret := _m.Called(newData)

	var r0 *donate.Transaction
	if rf, ok := ret.Get(0).(func(*donate.Transaction) *donate.Transaction); ok {
		r0 = rf(newData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*donate.Transaction)
		}
	}

	return r0
}

// Paginate provides a mock function with given fields: page, size
func (_m *Repository) Paginate(page int, size int) []donate.Transaction {
	ret := _m.Called(page, size)

	var r0 []donate.Transaction
	if rf, ok := ret.Get(0).(func(int, int) []donate.Transaction); ok {
		r0 = rf(page, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]donate.Transaction)
		}
	}

	return r0
}

// SelectByID provides a mock function with given fields: userID
func (_m *Repository) SelectByID(userID int) *donate.Transaction {
	ret := _m.Called(userID)

	var r0 *donate.Transaction
	if rf, ok := ret.Get(0).(func(int) *donate.Transaction); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*donate.Transaction)
		}
	}

	return r0
}

// SelectByOrderID provides a mock function with given fields: orderID
func (_m *Repository) SelectByOrderID(orderID string) (*donate.Transaction, error) {
	ret := _m.Called(orderID)

	var r0 *donate.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*donate.Transaction, error)); ok {
		return rf(orderID)
	}
	if rf, ok := ret.Get(0).(func(string) *donate.Transaction); ok {
		r0 = rf(orderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*donate.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SnapRequest provides a mock function with given fields: orderID, amount
func (_m *Repository) SnapRequest(orderID string, amount int64) (string, string) {
	ret := _m.Called(orderID, amount)

	var r0 string
	var r1 string
	if rf, ok := ret.Get(0).(func(string, int64) (string, string)); ok {
		return rf(orderID, amount)
	}
	if rf, ok := ret.Get(0).(func(string, int64) string); ok {
		r0 = rf(orderID, amount)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, int64) string); ok {
		r1 = rf(orderID, amount)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// UpdateStatusTransaction provides a mock function with given fields: id, status
func (_m *Repository) UpdateStatusTransaction(id uint, status string) error {
	ret := _m.Called(id, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, string) error); ok {
		r0 = rf(id, status)
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