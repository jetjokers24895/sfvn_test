// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"
	DTOS "sfvn_test/dtos/requests"

	dtos "sfvn_test/dtos/responses"

	mock "github.com/stretchr/testify/mock"
)

// ServiceWrapper is an autogenerated mock type for the ServiceWrapper type
type ServiceWrapper struct {
	mock.Mock
}

type ServiceWrapper_Expecter struct {
	mock *mock.Mock
}

func (_m *ServiceWrapper) EXPECT() *ServiceWrapper_Expecter {
	return &ServiceWrapper_Expecter{mock: &_m.Mock}
}

// GetHistoriesOfSymbol provides a mock function with given fields: _a0, _a1, _a2
func (_m *ServiceWrapper) GetHistoriesOfSymbol(_a0 context.Context, _a1 *DTOS.GetHistories, _a2 string) ([]*dtos.DTOGetHistoryPriceResponse, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for GetHistoriesOfSymbol")
	}

	var r0 []*dtos.DTOGetHistoryPriceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *DTOS.GetHistories, string) ([]*dtos.DTOGetHistoryPriceResponse, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *DTOS.GetHistories, string) []*dtos.DTOGetHistoryPriceResponse); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dtos.DTOGetHistoryPriceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *DTOS.GetHistories, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceWrapper_GetHistoriesOfSymbol_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHistoriesOfSymbol'
type ServiceWrapper_GetHistoriesOfSymbol_Call struct {
	*mock.Call
}

// GetHistoriesOfSymbol is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *DTOS.GetHistories
//   - _a2 string
func (_e *ServiceWrapper_Expecter) GetHistoriesOfSymbol(_a0 interface{}, _a1 interface{}, _a2 interface{}) *ServiceWrapper_GetHistoriesOfSymbol_Call {
	return &ServiceWrapper_GetHistoriesOfSymbol_Call{Call: _e.mock.On("GetHistoriesOfSymbol", _a0, _a1, _a2)}
}

func (_c *ServiceWrapper_GetHistoriesOfSymbol_Call) Run(run func(_a0 context.Context, _a1 *DTOS.GetHistories, _a2 string)) *ServiceWrapper_GetHistoriesOfSymbol_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*DTOS.GetHistories), args[2].(string))
	})
	return _c
}

func (_c *ServiceWrapper_GetHistoriesOfSymbol_Call) Return(_a0 []*dtos.DTOGetHistoryPriceResponse, _a1 error) *ServiceWrapper_GetHistoriesOfSymbol_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceWrapper_GetHistoriesOfSymbol_Call) RunAndReturn(run func(context.Context, *DTOS.GetHistories, string) ([]*dtos.DTOGetHistoryPriceResponse, error)) *ServiceWrapper_GetHistoriesOfSymbol_Call {
	_c.Call.Return(run)
	return _c
}

// NewServiceWrapper creates a new instance of ServiceWrapper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceWrapper(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceWrapper {
	mock := &ServiceWrapper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}