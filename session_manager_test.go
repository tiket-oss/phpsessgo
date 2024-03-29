// Code generated by MockGen. DO NOT EDIT.
// Source: session_manager.go

// Package mock_phpsessgo is a generated GoMock package.
package phpsessgo_test

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	phpsessgo "github.com/tiket-oss/phpsessgo"
)

// MockSessionManager is a mock of SessionManager interface.
type MockSessionManager struct {
	ctrl     *gomock.Controller
	recorder *MockSessionManagerMockRecorder
}

// MockSessionManagerMockRecorder is the mock recorder for MockSessionManager.
type MockSessionManagerMockRecorder struct {
	mock *MockSessionManager
}

// NewMockSessionManager creates a new mock instance.
func NewMockSessionManager(ctrl *gomock.Controller) *MockSessionManager {
	mock := &MockSessionManager{ctrl: ctrl}
	mock.recorder = &MockSessionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionManager) EXPECT() *MockSessionManagerMockRecorder {
	return m.recorder
}

// Encoder mocks base method.
func (m *MockSessionManager) Encoder() phpsessgo.SessionEncoder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encoder")
	ret0, _ := ret[0].(phpsessgo.SessionEncoder)
	return ret0
}

// Encoder indicates an expected call of Encoder.
func (mr *MockSessionManagerMockRecorder) Encoder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encoder", reflect.TypeOf((*MockSessionManager)(nil).Encoder))
}

// SIDCreator mocks base method.
func (m *MockSessionManager) SIDCreator() phpsessgo.SessionIDCreator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SIDCreator")
	ret0, _ := ret[0].(phpsessgo.SessionIDCreator)
	return ret0
}

// SIDCreator indicates an expected call of SIDCreator.
func (mr *MockSessionManagerMockRecorder) SIDCreator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SIDCreator", reflect.TypeOf((*MockSessionManager)(nil).SIDCreator))
}

// Save mocks base method.
func (m *MockSessionManager) Save(session *phpsessgo.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", session)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockSessionManagerMockRecorder) Save(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockSessionManager)(nil).Save), session)
}

// SessionName mocks base method.
func (m *MockSessionManager) SessionName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SessionName")
	ret0, _ := ret[0].(string)
	return ret0
}

// SessionName indicates an expected call of SessionName.
func (mr *MockSessionManagerMockRecorder) SessionName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SessionName", reflect.TypeOf((*MockSessionManager)(nil).SessionName))
}

// SetCookieString mocks base method.
func (m *MockSessionManager) SetCookieString(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCookieString", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// SetCookieString indicates an expected call of SetCookieString.
func (mr *MockSessionManagerMockRecorder) SetCookieString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCookieString", reflect.TypeOf((*MockSessionManager)(nil).SetCookieString), arg0)
}

// Start mocks base method.
func (m *MockSessionManager) Start(w http.ResponseWriter, r *http.Request) (*phpsessgo.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", w, r)
	ret0, _ := ret[0].(*phpsessgo.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Start indicates an expected call of Start.
func (mr *MockSessionManagerMockRecorder) Start(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockSessionManager)(nil).Start), w, r)
}
