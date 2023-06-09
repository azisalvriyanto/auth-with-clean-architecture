// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	user "auth-with-clean-architecture/internal/user"

	mock "github.com/stretchr/testify/mock"
)

// UseCaseInterface is an autogenerated mock type for the UseCaseInterface type
type UseCaseInterface struct {
	mock.Mock
}

type UseCaseInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *UseCaseInterface) EXPECT() *UseCaseInterface_Expecter {
	return &UseCaseInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0
func (_m *UseCaseInterface) Create(_a0 *user.User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*user.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UseCaseInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type UseCaseInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 *user.User
func (_e *UseCaseInterface_Expecter) Create(_a0 interface{}) *UseCaseInterface_Create_Call {
	return &UseCaseInterface_Create_Call{Call: _e.mock.On("Create", _a0)}
}

func (_c *UseCaseInterface_Create_Call) Run(run func(_a0 *user.User)) *UseCaseInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*user.User))
	})
	return _c
}

func (_c *UseCaseInterface_Create_Call) Return(_a0 error) *UseCaseInterface_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UseCaseInterface_Create_Call) RunAndReturn(run func(*user.User) error) *UseCaseInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Destroy provides a mock function with given fields: ID
func (_m *UseCaseInterface) Destroy(ID string) (*user.User, error) {
	ret := _m.Called(ID)

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*user.User, error)); ok {
		return rf(ID)
	}
	if rf, ok := ret.Get(0).(func(string) *user.User); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UseCaseInterface_Destroy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Destroy'
type UseCaseInterface_Destroy_Call struct {
	*mock.Call
}

// Destroy is a helper method to define mock.On call
//   - ID string
func (_e *UseCaseInterface_Expecter) Destroy(ID interface{}) *UseCaseInterface_Destroy_Call {
	return &UseCaseInterface_Destroy_Call{Call: _e.mock.On("Destroy", ID)}
}

func (_c *UseCaseInterface_Destroy_Call) Run(run func(ID string)) *UseCaseInterface_Destroy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *UseCaseInterface_Destroy_Call) Return(_a0 *user.User, _a1 error) *UseCaseInterface_Destroy_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UseCaseInterface_Destroy_Call) RunAndReturn(run func(string) (*user.User, error)) *UseCaseInterface_Destroy_Call {
	_c.Call.Return(run)
	return _c
}

// Show provides a mock function with given fields: ID
func (_m *UseCaseInterface) Show(ID string) (*user.User, error) {
	ret := _m.Called(ID)

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*user.User, error)); ok {
		return rf(ID)
	}
	if rf, ok := ret.Get(0).(func(string) *user.User); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UseCaseInterface_Show_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Show'
type UseCaseInterface_Show_Call struct {
	*mock.Call
}

// Show is a helper method to define mock.On call
//   - ID string
func (_e *UseCaseInterface_Expecter) Show(ID interface{}) *UseCaseInterface_Show_Call {
	return &UseCaseInterface_Show_Call{Call: _e.mock.On("Show", ID)}
}

func (_c *UseCaseInterface_Show_Call) Run(run func(ID string)) *UseCaseInterface_Show_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *UseCaseInterface_Show_Call) Return(_a0 *user.User, _a1 error) *UseCaseInterface_Show_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UseCaseInterface_Show_Call) RunAndReturn(run func(string) (*user.User, error)) *UseCaseInterface_Show_Call {
	_c.Call.Return(run)
	return _c
}

// ShowAll provides a mock function with given fields:
func (_m *UseCaseInterface) ShowAll() ([]user.User, error) {
	ret := _m.Called()

	var r0 []user.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]user.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []user.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UseCaseInterface_ShowAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShowAll'
type UseCaseInterface_ShowAll_Call struct {
	*mock.Call
}

// ShowAll is a helper method to define mock.On call
func (_e *UseCaseInterface_Expecter) ShowAll() *UseCaseInterface_ShowAll_Call {
	return &UseCaseInterface_ShowAll_Call{Call: _e.mock.On("ShowAll")}
}

func (_c *UseCaseInterface_ShowAll_Call) Run(run func()) *UseCaseInterface_ShowAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UseCaseInterface_ShowAll_Call) Return(_a0 []user.User, _a1 error) *UseCaseInterface_ShowAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UseCaseInterface_ShowAll_Call) RunAndReturn(run func() ([]user.User, error)) *UseCaseInterface_ShowAll_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ID, _a1
func (_m *UseCaseInterface) Update(ID string, _a1 user.User) (*user.User, error) {
	ret := _m.Called(ID, _a1)

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, user.User) (*user.User, error)); ok {
		return rf(ID, _a1)
	}
	if rf, ok := ret.Get(0).(func(string, user.User) *user.User); ok {
		r0 = rf(ID, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string, user.User) error); ok {
		r1 = rf(ID, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UseCaseInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type UseCaseInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ID string
//   - _a1 user.User
func (_e *UseCaseInterface_Expecter) Update(ID interface{}, _a1 interface{}) *UseCaseInterface_Update_Call {
	return &UseCaseInterface_Update_Call{Call: _e.mock.On("Update", ID, _a1)}
}

func (_c *UseCaseInterface_Update_Call) Run(run func(ID string, _a1 user.User)) *UseCaseInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(user.User))
	})
	return _c
}

func (_c *UseCaseInterface_Update_Call) Return(_a0 *user.User, _a1 error) *UseCaseInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UseCaseInterface_Update_Call) RunAndReturn(run func(string, user.User) (*user.User, error)) *UseCaseInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewUseCaseInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCaseInterface creates a new instance of UseCaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCaseInterface(t mockConstructorTestingTNewUseCaseInterface) *UseCaseInterface {
	mock := &UseCaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
