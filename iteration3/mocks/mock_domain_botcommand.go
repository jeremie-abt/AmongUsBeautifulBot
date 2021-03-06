// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain (interfaces: IBotCommand)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIBotCommand is a mock of IBotCommand interface
type MockIBotCommand struct {
	ctrl     *gomock.Controller
	recorder *MockIBotCommandMockRecorder
}

// MockIBotCommandMockRecorder is the mock recorder for MockIBotCommand
type MockIBotCommandMockRecorder struct {
	mock *MockIBotCommand
}

// NewMockIBotCommand creates a new mock instance
func NewMockIBotCommand(ctrl *gomock.Controller) *MockIBotCommand {
	mock := &MockIBotCommand{ctrl: ctrl}
	mock.recorder = &MockIBotCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIBotCommand) EXPECT() *MockIBotCommandMockRecorder {
	return m.recorder
}

// AddPlayer mocks base method
func (m *MockIBotCommand) AddPlayer(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPlayer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPlayer indicates an expected call of AddPlayer
func (mr *MockIBotCommandMockRecorder) AddPlayer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPlayer", reflect.TypeOf((*MockIBotCommand)(nil).AddPlayer), arg0, arg1)
}

// DeletePlayer mocks base method
func (m *MockIBotCommand) DeletePlayer(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlayer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePlayer indicates an expected call of DeletePlayer
func (mr *MockIBotCommandMockRecorder) DeletePlayer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlayer", reflect.TypeOf((*MockIBotCommand)(nil).DeletePlayer), arg0, arg1)
}

// IsGameIdExisting mocks base method
func (m *MockIBotCommand) IsGameIdExisting(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsGameIdExisting", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsGameIdExisting indicates an expected call of IsGameIdExisting
func (mr *MockIBotCommandMockRecorder) IsGameIdExisting(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsGameIdExisting", reflect.TypeOf((*MockIBotCommand)(nil).IsGameIdExisting), arg0)
}

// StartGame mocks base method
func (m *MockIBotCommand) StartGame(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartGame", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartGame indicates an expected call of StartGame
func (mr *MockIBotCommandMockRecorder) StartGame(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartGame", reflect.TypeOf((*MockIBotCommand)(nil).StartGame), arg0)
}

// StopGame mocks base method
func (m *MockIBotCommand) StopGame(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopGame", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopGame indicates an expected call of StopGame
func (mr *MockIBotCommandMockRecorder) StopGame(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopGame", reflect.TypeOf((*MockIBotCommand)(nil).StopGame), arg0)
}
