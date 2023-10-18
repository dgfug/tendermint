// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	state "github.com/tendermint/tendermint/internal/state"
)

// ConsSyncReactor is an autogenerated mock type for the ConsSyncReactor type
type ConsSyncReactor struct {
	mock.Mock
}

// SetBlockSyncingMetrics provides a mock function with given fields: _a0
func (_m *ConsSyncReactor) SetBlockSyncingMetrics(_a0 float64) {
	_m.Called(_a0)
}

// SetStateSyncingMetrics provides a mock function with given fields: _a0
func (_m *ConsSyncReactor) SetStateSyncingMetrics(_a0 float64) {
	_m.Called(_a0)
}

// SwitchToConsensus provides a mock function with given fields: _a0, _a1
func (_m *ConsSyncReactor) SwitchToConsensus(_a0 state.State, _a1 bool) {
	_m.Called(_a0, _a1)
}

type NewConsSyncReactorT interface {
	mock.TestingT
	Cleanup(func())
}

// NewConsSyncReactor creates a new instance of ConsSyncReactor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConsSyncReactor(t NewConsSyncReactorT) *ConsSyncReactor {
	mock := &ConsSyncReactor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
