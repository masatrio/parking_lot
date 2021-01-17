// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	slot "github.com/masatrio/parking_lot/parkinglot/slot"
	vehicle "github.com/masatrio/parking_lot/vehicle"
	reflect "reflect"
)

// MockAllocator is a mock of Allocator interface
type MockAllocator struct {
	ctrl     *gomock.Controller
	recorder *MockAllocatorMockRecorder
}

// MockAllocatorMockRecorder is the mock recorder for MockAllocator
type MockAllocatorMockRecorder struct {
	mock *MockAllocator
}

// NewMockAllocator creates a new mock instance
func NewMockAllocator(ctrl *gomock.Controller) *MockAllocator {
	mock := &MockAllocator{ctrl: ctrl}
	mock.recorder = &MockAllocatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAllocator) EXPECT() *MockAllocatorMockRecorder {
	return m.recorder
}

// Allocate mocks base method
func (m *MockAllocator) Allocate(vhc vehicle.Vehicle) (slot.Slot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Allocate", vhc)
	ret0, _ := ret[0].(slot.Slot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Allocate indicates an expected call of Allocate
func (mr *MockAllocatorMockRecorder) Allocate(vhc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Allocate", reflect.TypeOf((*MockAllocator)(nil).Allocate), vhc)
}