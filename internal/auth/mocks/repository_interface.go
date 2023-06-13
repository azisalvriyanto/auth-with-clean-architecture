// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	auth "auth-with-clean-architecture/internal/auth"

	mock "github.com/stretchr/testify/mock"

	user "auth-with-clean-architecture/internal/user"
)

// RepositoryInterface is an autogenerated mock type for the RepositoryInterface type
type RepositoryInterface struct {
	mock.Mock
}

type RepositoryInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *RepositoryInterface) EXPECT() *RepositoryInterface_Expecter {
	return &RepositoryInterface_Expecter{mock: &_m.Mock}
}

// FindByUsername provides a mock function with given fields: username
func (_m *RepositoryInterface) FindByUsername(username string) (*user.User, error) {
	ret := _m.Called(username)

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*user.User, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) *user.User); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoryInterface_FindByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByUsername'
type RepositoryInterface_FindByUsername_Call struct {
	*mock.Call
}

// FindByUsername is a helper method to define mock.On call
//   - username string
func (_e *RepositoryInterface_Expecter) FindByUsername(username interface{}) *RepositoryInterface_FindByUsername_Call {
	return &RepositoryInterface_FindByUsername_Call{Call: _e.mock.On("FindByUsername", username)}
}

func (_c *RepositoryInterface_FindByUsername_Call) Run(run func(username string)) *RepositoryInterface_FindByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RepositoryInterface_FindByUsername_Call) Return(_a0 *user.User, _a1 error) *RepositoryInterface_FindByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RepositoryInterface_FindByUsername_Call) RunAndReturn(run func(string) (*user.User, error)) *RepositoryInterface_FindByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// Login provides a mock function with given fields: _a0
func (_m *RepositoryInterface) Login(_a0 *auth.Payload) (*auth.ProfileItemWithToken, error) {
	ret := _m.Called(_a0)

	var r0 *auth.ProfileItemWithToken
	var r1 error
	if rf, ok := ret.Get(0).(func(*auth.Payload) (*auth.ProfileItemWithToken, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*auth.Payload) *auth.ProfileItemWithToken); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.ProfileItemWithToken)
		}
	}

	if rf, ok := ret.Get(1).(func(*auth.Payload) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoryInterface_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type RepositoryInterface_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - _a0 *auth.Payload
func (_e *RepositoryInterface_Expecter) Login(_a0 interface{}) *RepositoryInterface_Login_Call {
	return &RepositoryInterface_Login_Call{Call: _e.mock.On("Login", _a0)}
}

func (_c *RepositoryInterface_Login_Call) Run(run func(_a0 *auth.Payload)) *RepositoryInterface_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*auth.Payload))
	})
	return _c
}

func (_c *RepositoryInterface_Login_Call) Return(_a0 *auth.ProfileItemWithToken, _a1 error) *RepositoryInterface_Login_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RepositoryInterface_Login_Call) RunAndReturn(run func(*auth.Payload) (*auth.ProfileItemWithToken, error)) *RepositoryInterface_Login_Call {
	_c.Call.Return(run)
	return _c
}

// ShowProfile provides a mock function with given fields: tokenSigned
func (_m *RepositoryInterface) ShowProfile(tokenSigned string) (*user.User, error) {
	ret := _m.Called(tokenSigned)

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*user.User, error)); ok {
		return rf(tokenSigned)
	}
	if rf, ok := ret.Get(0).(func(string) *user.User); ok {
		r0 = rf(tokenSigned)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenSigned)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoryInterface_ShowProfile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShowProfile'
type RepositoryInterface_ShowProfile_Call struct {
	*mock.Call
}

// ShowProfile is a helper method to define mock.On call
//   - tokenSigned string
func (_e *RepositoryInterface_Expecter) ShowProfile(tokenSigned interface{}) *RepositoryInterface_ShowProfile_Call {
	return &RepositoryInterface_ShowProfile_Call{Call: _e.mock.On("ShowProfile", tokenSigned)}
}

func (_c *RepositoryInterface_ShowProfile_Call) Run(run func(tokenSigned string)) *RepositoryInterface_ShowProfile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RepositoryInterface_ShowProfile_Call) Return(_a0 *user.User, _a1 error) *RepositoryInterface_ShowProfile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RepositoryInterface_ShowProfile_Call) RunAndReturn(run func(string) (*user.User, error)) *RepositoryInterface_ShowProfile_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewRepositoryInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepositoryInterface creates a new instance of RepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepositoryInterface(t mockConstructorTestingTNewRepositoryInterface) *RepositoryInterface {
	mock := &RepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
