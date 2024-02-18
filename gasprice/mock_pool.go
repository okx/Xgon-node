// Code generated by mockery v2.40.1. DO NOT EDIT.

package gasprice

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	pool "github.com/0xPolygonHermez/zkevm-node/pool"

	time "time"
)

// poolMock is an autogenerated mock type for the poolInterface type
type poolMock struct {
	mock.Mock
}

// CountPendingTransactions provides a mock function with given fields: ctx
func (_m *poolMock) CountPendingTransactions(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CountPendingTransactions")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteGasPricesHistoryOlderThan provides a mock function with given fields: ctx, date
func (_m *poolMock) DeleteGasPricesHistoryOlderThan(ctx context.Context, date time.Time) error {
	ret := _m.Called(ctx, date)

	if len(ret) == 0 {
		panic("no return value specified for DeleteGasPricesHistoryOlderThan")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time) error); ok {
		r0 = rf(ctx, date)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetGasPrices provides a mock function with given fields: ctx
func (_m *poolMock) GetGasPrices(ctx context.Context) (pool.GasPrices, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetGasPrices")
	}

	var r0 pool.GasPrices
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pool.GasPrices, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pool.GasPrices); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(pool.GasPrices)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetGasPrices provides a mock function with given fields: ctx, l2GasPrice, l1GasPrice
func (_m *poolMock) SetGasPrices(ctx context.Context, l2GasPrice uint64, l1GasPrice uint64) error {
	ret := _m.Called(ctx, l2GasPrice, l1GasPrice)

	if len(ret) == 0 {
		panic("no return value specified for SetGasPrices")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64) error); ok {
		r0 = rf(ctx, l2GasPrice, l1GasPrice)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newPoolMock creates a new instance of poolMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newPoolMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *poolMock {
	mock := &poolMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
