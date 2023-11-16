// Code generated by mockery v2.35.1. DO NOT EDIT.

package mocks

import (
	domain "2023_2_Holi/domain"

	mock "github.com/stretchr/testify/mock"
)

// ProfileUsecase is an autogenerated mock type for the ProfileUsecase type
type ProfileUsecase struct {
	mock.Mock
}

// GetUserData provides a mock function with given fields: userID
func (_m *ProfileUsecase) GetUserData(userID int) (domain.User, error) {
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
func (_m *ProfileUsecase) UpdateUser(newUser domain.User) (domain.User, error) {
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

// UploadImage provides a mock function with given fields: userID, imageData
func (_m *ProfileUsecase) UploadImage(userID int, imageData []byte) (string, error) {
	ret := _m.Called(userID, imageData)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int, []byte) (string, error)); ok {
		return rf(userID, imageData)
	}
	if rf, ok := ret.Get(0).(func(int, []byte) string); ok {
		r0 = rf(userID, imageData)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int, []byte) error); ok {
		r1 = rf(userID, imageData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProfileUsecase creates a new instance of ProfileUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProfileUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProfileUsecase {
	mock := &ProfileUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
