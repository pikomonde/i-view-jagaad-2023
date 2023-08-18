// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	model "i-view-jagaad-2023/model"

	mock "github.com/stretchr/testify/mock"
)

// UserProviderRepository is an autogenerated mock type for the UserProviderRepository type
type UserProviderRepository struct {
	mock.Mock
}

type UserProviderRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *UserProviderRepository) EXPECT() *UserProviderRepository_Expecter {
	return &UserProviderRepository_Expecter{mock: &_m.Mock}
}

// FetchUsersFromProvider provides a mock function with given fields:
func (_m *UserProviderRepository) FetchUsersFromProvider() ([]model.User, error) {
	ret := _m.Called()

	var r0 []model.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserProviderRepository_FetchUsersFromProvider_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchUsersFromProvider'
type UserProviderRepository_FetchUsersFromProvider_Call struct {
	*mock.Call
}

// FetchUsersFromProvider is a helper method to define mock.On call
func (_e *UserProviderRepository_Expecter) FetchUsersFromProvider() *UserProviderRepository_FetchUsersFromProvider_Call {
	return &UserProviderRepository_FetchUsersFromProvider_Call{Call: _e.mock.On("FetchUsersFromProvider")}
}

func (_c *UserProviderRepository_FetchUsersFromProvider_Call) Run(run func()) *UserProviderRepository_FetchUsersFromProvider_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UserProviderRepository_FetchUsersFromProvider_Call) Return(_a0 []model.User, _a1 error) *UserProviderRepository_FetchUsersFromProvider_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserProviderRepository_FetchUsersFromProvider_Call) RunAndReturn(run func() ([]model.User, error)) *UserProviderRepository_FetchUsersFromProvider_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserProviderRepository creates a new instance of UserProviderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserProviderRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserProviderRepository {
	mock := &UserProviderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
