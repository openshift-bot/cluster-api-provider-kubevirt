// Code generated by MockGen. DO NOT EDIT.
// Source: ./kubevirt.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	machinescope "github.com/openshift/cluster-api-provider-kubevirt/pkg/machinescope"
	reflect "reflect"
)

// MockKubevirtVM is a mock of KubevirtVM interface
type MockKubevirtVM struct {
	ctrl     *gomock.Controller
	recorder *MockKubevirtVMMockRecorder
}

// MockKubevirtVMMockRecorder is the mock recorder for MockKubevirtVM
type MockKubevirtVMMockRecorder struct {
	mock *MockKubevirtVM
}

// NewMockKubevirtVM creates a new mock instance
func NewMockKubevirtVM(ctrl *gomock.Controller) *MockKubevirtVM {
	mock := &MockKubevirtVM{ctrl: ctrl}
	mock.recorder = &MockKubevirtVMMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubevirtVM) EXPECT() *MockKubevirtVMMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockKubevirtVM) Create(machineScope machinescope.MachineScope, userData []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", machineScope, userData)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockKubevirtVMMockRecorder) Create(machineScope, userData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockKubevirtVM)(nil).Create), machineScope, userData)
}

// Delete mocks base method
func (m *MockKubevirtVM) Delete(machineScope machinescope.MachineScope) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", machineScope)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockKubevirtVMMockRecorder) Delete(machineScope interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockKubevirtVM)(nil).Delete), machineScope)
}

// Update mocks base method
func (m *MockKubevirtVM) Update(machineScope machinescope.MachineScope) (bool, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", machineScope)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update
func (mr *MockKubevirtVMMockRecorder) Update(machineScope interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockKubevirtVM)(nil).Update), machineScope)
}

// Exists mocks base method
func (m *MockKubevirtVM) Exists(machineName, infraNamespace string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", machineName, infraNamespace)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists
func (mr *MockKubevirtVMMockRecorder) Exists(machineName, infraNamespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockKubevirtVM)(nil).Exists), machineName, infraNamespace)
}
