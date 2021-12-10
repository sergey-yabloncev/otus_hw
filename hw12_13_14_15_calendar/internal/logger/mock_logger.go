// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package logger is a generated GoMock package.
package logger

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockContract is a mock of Contract interface.
type MockContract struct {
	ctrl     *gomock.Controller
	recorder *MockContractMockRecorder
}

// MockContractMockRecorder is the mock recorder for MockContract.
type MockContractMockRecorder struct {
	mock *MockContract
}

// NewMockContract creates a new mock instance.
func NewMockContract(ctrl *gomock.Controller) *MockContract {
	mock := &MockContract{ctrl: ctrl}
	mock.recorder = &MockContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContract) EXPECT() *MockContractMockRecorder {
	return m.recorder
}

// Debug mocks base method.
func (m *MockContract) Debug(msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Debug", msg)
}

// Debug indicates an expected call of Debug.
func (mr *MockContractMockRecorder) Debug(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockContract)(nil).Debug), msg)
}

// Error mocks base method.
func (m *MockContract) Error(msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Error", msg)
}

// Error indicates an expected call of Error.
func (mr *MockContractMockRecorder) Error(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockContract)(nil).Error), msg)
}

// Info mocks base method.
func (m *MockContract) Info(msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Info", msg)
}

// Info indicates an expected call of Info.
func (mr *MockContractMockRecorder) Info(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockContract)(nil).Info), msg)
}

// Warn mocks base method.
func (m *MockContract) Warn(msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Warn", msg)
}

// Warn indicates an expected call of Warn.
func (mr *MockContractMockRecorder) Warn(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockContract)(nil).Warn), msg)
}
