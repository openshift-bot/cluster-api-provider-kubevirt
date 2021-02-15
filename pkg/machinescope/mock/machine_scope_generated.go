// Code generated by MockGen. DO NOT EDIT.
// Source: ./machine_scope.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
	v1 "k8s.io/api/core/v1"
	v10 "kubevirt.io/client-go/api/v1"
	reflect "reflect"
	time "time"
)

// MockMachineScope is a mock of MachineScope interface
type MockMachineScope struct {
	ctrl     *gomock.Controller
	recorder *MockMachineScopeMockRecorder
}

// MockMachineScopeMockRecorder is the mock recorder for MockMachineScope
type MockMachineScopeMockRecorder struct {
	mock *MockMachineScope
}

// NewMockMachineScope creates a new mock instance
func NewMockMachineScope(ctrl *gomock.Controller) *MockMachineScope {
	mock := &MockMachineScope{ctrl: ctrl}
	mock.recorder = &MockMachineScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMachineScope) EXPECT() *MockMachineScopeMockRecorder {
	return m.recorder
}

// UpdateAllowed mocks base method
func (m *MockMachineScope) UpdateAllowed(requeueAfterSeconds time.Duration) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllowed", requeueAfterSeconds)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UpdateAllowed indicates an expected call of UpdateAllowed
func (mr *MockMachineScopeMockRecorder) UpdateAllowed(requeueAfterSeconds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllowed", reflect.TypeOf((*MockMachineScope)(nil).UpdateAllowed), requeueAfterSeconds)
}

// CreateIgnitionSecretFromMachine mocks base method
func (m *MockMachineScope) CreateIgnitionSecretFromMachine(userData []byte) *v1.Secret {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIgnitionSecretFromMachine", userData)
	ret0, _ := ret[0].(*v1.Secret)
	return ret0
}

// CreateIgnitionSecretFromMachine indicates an expected call of CreateIgnitionSecretFromMachine
func (mr *MockMachineScopeMockRecorder) CreateIgnitionSecretFromMachine(userData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIgnitionSecretFromMachine", reflect.TypeOf((*MockMachineScope)(nil).CreateIgnitionSecretFromMachine), userData)
}

// SyncMachine mocks base method
func (m *MockMachineScope) SyncMachine(vm v10.VirtualMachine, vmi *v10.VirtualMachineInstance, providerID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncMachine", vm, vmi, providerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncMachine indicates an expected call of SyncMachine
func (mr *MockMachineScopeMockRecorder) SyncMachine(vm, vmi, providerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncMachine", reflect.TypeOf((*MockMachineScope)(nil).SyncMachine), vm, vmi, providerID)
}

// CreateVirtualMachineFromMachine mocks base method
func (m *MockMachineScope) CreateVirtualMachineFromMachine() (*v10.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVirtualMachineFromMachine")
	ret0, _ := ret[0].(*v10.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVirtualMachineFromMachine indicates an expected call of CreateVirtualMachineFromMachine
func (mr *MockMachineScopeMockRecorder) CreateVirtualMachineFromMachine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVirtualMachineFromMachine", reflect.TypeOf((*MockMachineScope)(nil).CreateVirtualMachineFromMachine))
}

// GetMachine mocks base method
func (m *MockMachineScope) GetMachine() *v1beta1.Machine {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachine")
	ret0, _ := ret[0].(*v1beta1.Machine)
	return ret0
}

// GetMachine indicates an expected call of GetMachine
func (mr *MockMachineScopeMockRecorder) GetMachine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachine", reflect.TypeOf((*MockMachineScope)(nil).GetMachine))
}

// GetMachineName mocks base method
func (m *MockMachineScope) GetMachineName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetMachineName indicates an expected call of GetMachineName
func (mr *MockMachineScopeMockRecorder) GetMachineName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineName", reflect.TypeOf((*MockMachineScope)(nil).GetMachineName))
}

// GetMachineNamespace mocks base method
func (m *MockMachineScope) GetMachineNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetMachineNamespace indicates an expected call of GetMachineNamespace
func (mr *MockMachineScopeMockRecorder) GetMachineNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineNamespace", reflect.TypeOf((*MockMachineScope)(nil).GetMachineNamespace))
}

// GetInfraNamespace mocks base method
func (m *MockMachineScope) GetInfraNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfraNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetInfraNamespace indicates an expected call of GetInfraNamespace
func (mr *MockMachineScopeMockRecorder) GetInfraNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfraNamespace", reflect.TypeOf((*MockMachineScope)(nil).GetInfraNamespace))
}

// GetIgnitionSecretName mocks base method
func (m *MockMachineScope) GetIgnitionSecretName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIgnitionSecretName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetIgnitionSecretName indicates an expected call of GetIgnitionSecretName
func (mr *MockMachineScopeMockRecorder) GetIgnitionSecretName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIgnitionSecretName", reflect.TypeOf((*MockMachineScope)(nil).GetIgnitionSecretName))
}
