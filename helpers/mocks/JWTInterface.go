// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	jwt "github.com/golang-jwt/jwt/v5"
	mock "github.com/stretchr/testify/mock"
)

// JWTInterface is an autogenerated mock type for the JWTInterface type
type JWTInterface struct {
	mock.Mock
}

// ExtractToken provides a mock function with given fields: token
func (_m *JWTInterface) ExtractToken(token *jwt.Token) interface{} {
	ret := _m.Called(token)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(*jwt.Token) interface{}); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// GenerateJWT provides a mock function with given fields: userID
func (_m *JWTInterface) GenerateJWT(userID string) map[string]interface{} {
	ret := _m.Called(userID)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(string) map[string]interface{}); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// GenerateToken provides a mock function with given fields: id
func (_m *JWTInterface) GenerateToken(id string) string {
	ret := _m.Called(id)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ValidateToken provides a mock function with given fields: token
func (_m *JWTInterface) ValidateToken(token string) (*jwt.Token, error) {
	ret := _m.Called(token)

	var r0 *jwt.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*jwt.Token, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *jwt.Token); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewJWTInterface creates a new instance of JWTInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJWTInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *JWTInterface {
	mock := &JWTInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
