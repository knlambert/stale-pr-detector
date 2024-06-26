// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/formatter.go

// Package pkg_mock is a generated GoMock package.
package pkg_mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFormatter is a mock of Formatter interface
type MockFormatter struct {
	ctrl     *gomock.Controller
	recorder *MockFormatterMockRecorder
}

// MockFormatterMockRecorder is the mock recorder for MockFormatter
type MockFormatterMockRecorder struct {
	mock *MockFormatter
}

// NewMockFormatter creates a new mock instance
func NewMockFormatter(ctrl *gomock.Controller) *MockFormatter {
	mock := &MockFormatter{ctrl: ctrl}
	mock.recorder = &MockFormatterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFormatter) EXPECT() *MockFormatterMockRecorder {
	return m.recorder
}

// PrettyPrint mocks base method
func (m *MockFormatter) PrettyPrint(values interface{}) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrettyPrint", values)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrettyPrint indicates an expected call of PrettyPrint
func (mr *MockFormatterMockRecorder) PrettyPrint(values interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrettyPrint", reflect.TypeOf((*MockFormatter)(nil).PrettyPrint), values)
}
