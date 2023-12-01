// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/maintenance (interfaces: Maintenance)
//
// Generated by this command:
//
//	mockgen -destination=mocks/maintenance.go -package=mocks github.com/openshift/managed-upgrade-operator/pkg/maintenance Maintenance
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockMaintenance is a mock of Maintenance interface.
type MockMaintenance struct {
	ctrl     *gomock.Controller
	recorder *MockMaintenanceMockRecorder
}

// MockMaintenanceMockRecorder is the mock recorder for MockMaintenance.
type MockMaintenanceMockRecorder struct {
	mock *MockMaintenance
}

// NewMockMaintenance creates a new mock instance.
func NewMockMaintenance(ctrl *gomock.Controller) *MockMaintenance {
	mock := &MockMaintenance{ctrl: ctrl}
	mock.recorder = &MockMaintenanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMaintenance) EXPECT() *MockMaintenanceMockRecorder {
	return m.recorder
}

// EndControlPlane mocks base method.
func (m *MockMaintenance) EndControlPlane() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndControlPlane")
	ret0, _ := ret[0].(error)
	return ret0
}

// EndControlPlane indicates an expected call of EndControlPlane.
func (mr *MockMaintenanceMockRecorder) EndControlPlane() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndControlPlane", reflect.TypeOf((*MockMaintenance)(nil).EndControlPlane))
}

// EndSilences mocks base method.
func (m *MockMaintenance) EndSilences(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndSilences", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EndSilences indicates an expected call of EndSilences.
func (mr *MockMaintenanceMockRecorder) EndSilences(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndSilences", reflect.TypeOf((*MockMaintenance)(nil).EndSilences), arg0)
}

// EndWorker mocks base method.
func (m *MockMaintenance) EndWorker() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndWorker")
	ret0, _ := ret[0].(error)
	return ret0
}

// EndWorker indicates an expected call of EndWorker.
func (mr *MockMaintenanceMockRecorder) EndWorker() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndWorker", reflect.TypeOf((*MockMaintenance)(nil).EndWorker))
}

// IsActive mocks base method.
func (m *MockMaintenance) IsActive() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsActive")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsActive indicates an expected call of IsActive.
func (mr *MockMaintenanceMockRecorder) IsActive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsActive", reflect.TypeOf((*MockMaintenance)(nil).IsActive))
}

// SetWorker mocks base method.
func (m *MockMaintenance) SetWorker(arg0 time.Time, arg1 string, arg2 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWorker", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWorker indicates an expected call of SetWorker.
func (mr *MockMaintenanceMockRecorder) SetWorker(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWorker", reflect.TypeOf((*MockMaintenance)(nil).SetWorker), arg0, arg1, arg2)
}

// StartControlPlane mocks base method.
func (m *MockMaintenance) StartControlPlane(arg0 time.Time, arg1 string, arg2 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartControlPlane", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartControlPlane indicates an expected call of StartControlPlane.
func (mr *MockMaintenanceMockRecorder) StartControlPlane(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartControlPlane", reflect.TypeOf((*MockMaintenance)(nil).StartControlPlane), arg0, arg1, arg2)
}
