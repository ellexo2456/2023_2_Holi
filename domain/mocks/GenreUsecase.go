// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	domain "2023_2_Holi/domain"

	mock "github.com/stretchr/testify/mock"
)

// GenreUsecase is an autogenerated mock type for the GenreUsecase type
type GenreUsecase struct {
	mock.Mock
}

// GetGenres provides a mock function with given fields:
func (_m *GenreUsecase) GetGenres() ([]domain.Genre, error) {
	ret := _m.Called()

	var r0 []domain.Genre
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.Genre, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.Genre); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Genre)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGenreUsecase creates a new instance of GenreUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGenreUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *GenreUsecase {
	mock := &GenreUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
