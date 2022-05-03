// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	types "github.com/axieinfinity/bridge-v2/internal/types"
)

// IReceipt is an autogenerated mock type for the IReceipt type
type IReceipt struct {
	mock.Mock
}

// GetLogs provides a mock function with given fields:
func (_m *IReceipt) GetLogs() []types.ILog {
	ret := _m.Called()

	var r0 []types.ILog
	if rf, ok := ret.Get(0).(func() []types.ILog); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.ILog)
		}
	}

	return r0
}

// GetStatus provides a mock function with given fields:
func (_m *IReceipt) GetStatus() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetTransaction provides a mock function with given fields:
func (_m *IReceipt) GetTransaction() types.ITransaction {
	ret := _m.Called()

	var r0 types.ITransaction
	if rf, ok := ret.Get(0).(func() types.ITransaction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ITransaction)
		}
	}

	return r0
}

// NewIReceipt creates a new instance of IReceipt. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewIReceipt(t testing.TB) *IReceipt {
	mock := &IReceipt{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
