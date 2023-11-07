// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	domain "2023_2_Holi/domain"

	mock "github.com/stretchr/testify/mock"
)

// SessionRepository is an autogenerated mock type for the SessionRepository type
type SessionRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: session
func (_m *SessionRepository) Add(session domain.Session) error {
	ret := _m.Called(session)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Session) error); ok {
		r0 = rf(session)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByToken provides a mock function with given fields: token
func (_m *SessionRepository) DeleteByToken(token string) error {
	ret := _m.Called(token)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SessionExists provides a mock function with given fields: token
func (_m *SessionRepository) SessionExists(token string) (bool, error) {
	ret := _m.Called(token)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSessionRepository creates a new instance of SessionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSessionRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *SessionRepository {
	mock := &SessionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
