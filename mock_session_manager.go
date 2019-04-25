// Code generated by MockGen. DO NOT EDIT.
// Source: session_manager.go

// Package phpsessgo is a generated GoMock package.
package phpsessgo

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockSessionManager is a mock of SessionManager interface
type MockSessionManager struct {
	ctrl     *gomock.Controller
	recorder *MockSessionManagerMockRecorder
}

// MockSessionManagerMockRecorder is the mock recorder for MockSessionManager
type MockSessionManagerMockRecorder struct {
	mock *MockSessionManager
}

// NewMockSessionManager creates a new mock instance
func NewMockSessionManager(ctrl *gomock.Controller) *MockSessionManager {
	mock := &MockSessionManager{ctrl: ctrl}
	mock.recorder = &MockSessionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessionManager) EXPECT() *MockSessionManagerMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockSessionManager) Start(w http.ResponseWriter, r *http.Request) (*Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", w, r)
	ret0, _ := ret[0].(*Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Start indicates an expected call of Start
func (mr *MockSessionManagerMockRecorder) Start(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockSessionManager)(nil).Start), w, r)
}

// Save mocks base method
func (m *MockSessionManager) Save(session *Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", session)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockSessionManagerMockRecorder) Save(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockSessionManager)(nil).Save), session)
}

// SessionName mocks base method
func (m *MockSessionManager) SessionName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SessionName")
	ret0, _ := ret[0].(string)
	return ret0
}

// SessionName indicates an expected call of SessionName
func (mr *MockSessionManagerMockRecorder) SessionName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SessionName", reflect.TypeOf((*MockSessionManager)(nil).SessionName))
}

// SIDCreator mocks base method
func (m *MockSessionManager) SIDCreator() SessionIDCreator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SIDCreator")
	ret0, _ := ret[0].(SessionIDCreator)
	return ret0
}

// SIDCreator indicates an expected call of SIDCreator
func (mr *MockSessionManagerMockRecorder) SIDCreator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SIDCreator", reflect.TypeOf((*MockSessionManager)(nil).SIDCreator))
}

// Handler mocks base method
func (m *MockSessionManager) Handler() SessionHandler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handler")
	ret0, _ := ret[0].(SessionHandler)
	return ret0
}

// Handler indicates an expected call of Handler
func (mr *MockSessionManagerMockRecorder) Handler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handler", reflect.TypeOf((*MockSessionManager)(nil).Handler))
}

// Encoder mocks base method
func (m *MockSessionManager) Encoder() SessionEncoder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encoder")
	ret0, _ := ret[0].(SessionEncoder)
	return ret0
}

// Encoder indicates an expected call of Encoder
func (mr *MockSessionManagerMockRecorder) Encoder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encoder", reflect.TypeOf((*MockSessionManager)(nil).Encoder))
}
