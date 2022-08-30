//
//   Copyright (c) 2022 Intel Corporation.
//
//   SPDX-License-Identifier: Apache-2.0
//
//
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/intel/edge-conductor/pkg/eputils/docker (interfaces: DockerInterface)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	plugins "github.com/intel/edge-conductor/pkg/api/plugins"
)

// MockDockerInterface is a mock of DockerInterface interface.
type MockDockerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDockerInterfaceMockRecorder
}

// MockDockerInterfaceMockRecorder is the mock recorder for MockDockerInterface.
type MockDockerInterfaceMockRecorder struct {
	mock *MockDockerInterface
}

// NewMockDockerInterface creates a new mock instance.
func NewMockDockerInterface(ctrl *gomock.Controller) *MockDockerInterface {
	mock := &MockDockerInterface{ctrl: ctrl}
	mock.recorder = &MockDockerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDockerInterface) EXPECT() *MockDockerInterfaceMockRecorder {
	return m.recorder
}

// DockerCreate mocks base method.
func (m *MockDockerInterface) DockerCreate(arg0 *plugins.ContainersItems0) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DockerCreate", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DockerCreate indicates an expected call of DockerCreate.
func (mr *MockDockerInterfaceMockRecorder) DockerCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DockerCreate", reflect.TypeOf((*MockDockerInterface)(nil).DockerCreate), arg0)
}

// DockerRemove mocks base method.
func (m *MockDockerInterface) DockerRemove(arg0 *plugins.ContainersItems0) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DockerRemove", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DockerRemove indicates an expected call of DockerRemove.
func (mr *MockDockerInterfaceMockRecorder) DockerRemove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DockerRemove", reflect.TypeOf((*MockDockerInterface)(nil).DockerRemove), arg0)
}

// DockerRun mocks base method.
func (m *MockDockerInterface) DockerRun(arg0 *plugins.ContainersItems0) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DockerRun", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DockerRun indicates an expected call of DockerRun.
func (mr *MockDockerInterfaceMockRecorder) DockerRun(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DockerRun", reflect.TypeOf((*MockDockerInterface)(nil).DockerRun), arg0)
}

// DockerStart mocks base method.
func (m *MockDockerInterface) DockerStart(arg0 *plugins.ContainersItems0) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DockerStart", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DockerStart indicates an expected call of DockerStart.
func (mr *MockDockerInterfaceMockRecorder) DockerStart(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DockerStart", reflect.TypeOf((*MockDockerInterface)(nil).DockerStart), arg0)
}

// DockerStop mocks base method.
func (m *MockDockerInterface) DockerStop(arg0 *plugins.ContainersItems0) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DockerStop", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DockerStop indicates an expected call of DockerStop.
func (mr *MockDockerInterfaceMockRecorder) DockerStop(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DockerStop", reflect.TypeOf((*MockDockerInterface)(nil).DockerStop), arg0)
}
