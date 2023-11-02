// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	domain "2023_2_Holi/domain"

	mock "github.com/stretchr/testify/mock"
)

// ProfileRepository is an autogenerated mock type for the ProfileRepository type
type ProfileRepository struct {
	mock.Mock
}

// GetUser provides a mock function with given fields: userID
func (_m *ProfileRepository) GetUser(userID int) (domain.User, error) {
	ret := _m.Called(userID)

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (domain.User, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(int) domain.User); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: newUser
func (_m *ProfileRepository) UpdateUser(newUser domain.User) (domain.User, error) {
	ret := _m.Called(newUser)

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.User) (domain.User, error)); ok {
		return rf(newUser)
	}
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProfileRepository creates a new instance of ProfileRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProfileRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProfileRepository {
	mock := &ProfileRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}