// Code generated by mockery v2.53.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	system "github.com/tbe-team/raybot/internal/services/system"
)

// FakeService is an autogenerated mock type for the Service type
type FakeService struct {
	mock.Mock
}

type FakeService_Expecter struct {
	mock *mock.Mock
}

func (_m *FakeService) EXPECT() *FakeService_Expecter {
	return &FakeService_Expecter{mock: &_m.Mock}
}

// GetInfo provides a mock function with given fields: ctx
func (_m *FakeService) GetInfo(ctx context.Context) (system.Info, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetInfo")
	}

	var r0 system.Info
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (system.Info, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) system.Info); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(system.Info)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FakeService_GetInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInfo'
type FakeService_GetInfo_Call struct {
	*mock.Call
}

// GetInfo is a helper method to define mock.On call
//   - ctx context.Context
func (_e *FakeService_Expecter) GetInfo(ctx interface{}) *FakeService_GetInfo_Call {
	return &FakeService_GetInfo_Call{Call: _e.mock.On("GetInfo", ctx)}
}

func (_c *FakeService_GetInfo_Call) Run(run func(ctx context.Context)) *FakeService_GetInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *FakeService_GetInfo_Call) Return(_a0 system.Info, _a1 error) *FakeService_GetInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FakeService_GetInfo_Call) RunAndReturn(run func(context.Context) (system.Info, error)) *FakeService_GetInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetStatus provides a mock function with given fields: ctx
func (_m *FakeService) GetStatus(ctx context.Context) (system.Status, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetStatus")
	}

	var r0 system.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (system.Status, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) system.Status); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(system.Status)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FakeService_GetStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatus'
type FakeService_GetStatus_Call struct {
	*mock.Call
}

// GetStatus is a helper method to define mock.On call
//   - ctx context.Context
func (_e *FakeService_Expecter) GetStatus(ctx interface{}) *FakeService_GetStatus_Call {
	return &FakeService_GetStatus_Call{Call: _e.mock.On("GetStatus", ctx)}
}

func (_c *FakeService_GetStatus_Call) Run(run func(ctx context.Context)) *FakeService_GetStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *FakeService_GetStatus_Call) Return(_a0 system.Status, _a1 error) *FakeService_GetStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FakeService_GetStatus_Call) RunAndReturn(run func(context.Context) (system.Status, error)) *FakeService_GetStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Reboot provides a mock function with given fields: ctx
func (_m *FakeService) Reboot(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Reboot")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FakeService_Reboot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reboot'
type FakeService_Reboot_Call struct {
	*mock.Call
}

// Reboot is a helper method to define mock.On call
//   - ctx context.Context
func (_e *FakeService_Expecter) Reboot(ctx interface{}) *FakeService_Reboot_Call {
	return &FakeService_Reboot_Call{Call: _e.mock.On("Reboot", ctx)}
}

func (_c *FakeService_Reboot_Call) Run(run func(ctx context.Context)) *FakeService_Reboot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *FakeService_Reboot_Call) Return(_a0 error) *FakeService_Reboot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FakeService_Reboot_Call) RunAndReturn(run func(context.Context) error) *FakeService_Reboot_Call {
	_c.Call.Return(run)
	return _c
}

// SetStatusError provides a mock function with given fields: ctx
func (_m *FakeService) SetStatusError(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for SetStatusError")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FakeService_SetStatusError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetStatusError'
type FakeService_SetStatusError_Call struct {
	*mock.Call
}

// SetStatusError is a helper method to define mock.On call
//   - ctx context.Context
func (_e *FakeService_Expecter) SetStatusError(ctx interface{}) *FakeService_SetStatusError_Call {
	return &FakeService_SetStatusError_Call{Call: _e.mock.On("SetStatusError", ctx)}
}

func (_c *FakeService_SetStatusError_Call) Run(run func(ctx context.Context)) *FakeService_SetStatusError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *FakeService_SetStatusError_Call) Return(_a0 error) *FakeService_SetStatusError_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FakeService_SetStatusError_Call) RunAndReturn(run func(context.Context) error) *FakeService_SetStatusError_Call {
	_c.Call.Return(run)
	return _c
}

// StopEmergency provides a mock function with given fields: ctx
func (_m *FakeService) StopEmergency(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for StopEmergency")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FakeService_StopEmergency_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StopEmergency'
type FakeService_StopEmergency_Call struct {
	*mock.Call
}

// StopEmergency is a helper method to define mock.On call
//   - ctx context.Context
func (_e *FakeService_Expecter) StopEmergency(ctx interface{}) *FakeService_StopEmergency_Call {
	return &FakeService_StopEmergency_Call{Call: _e.mock.On("StopEmergency", ctx)}
}

func (_c *FakeService_StopEmergency_Call) Run(run func(ctx context.Context)) *FakeService_StopEmergency_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *FakeService_StopEmergency_Call) Return(_a0 error) *FakeService_StopEmergency_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FakeService_StopEmergency_Call) RunAndReturn(run func(context.Context) error) *FakeService_StopEmergency_Call {
	_c.Call.Return(run)
	return _c
}

// NewFakeService creates a new instance of FakeService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFakeService(t interface {
	mock.TestingT
	Cleanup(func())
}) *FakeService {
	mock := &FakeService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
