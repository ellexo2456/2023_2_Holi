// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	session "2023_2_Holi/domain/grpc/session"
)

// AuthCheckerClient is an autogenerated mock type for the AuthCheckerClient type
type AuthCheckerClient struct {
	mock.Mock
}

// IsAuth provides a mock function with given fields: ctx, in, opts
func (_m *AuthCheckerClient) IsAuth(ctx context.Context, in *session.Token, opts ...grpc.CallOption) (*session.UserID, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *session.UserID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *session.Token, ...grpc.CallOption) (*session.UserID, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *session.Token, ...grpc.CallOption) *session.UserID); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*session.UserID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *session.Token, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthCheckerClient creates a new instance of AuthCheckerClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthCheckerClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthCheckerClient {
	mock := &AuthCheckerClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
