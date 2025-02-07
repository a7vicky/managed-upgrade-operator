// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/metrics (interfaces: Metrics)
//
// Generated by this command:
//
//	mockgen -destination=mocks/metrics.go -package=mocks github.com/openshift/managed-upgrade-operator/pkg/metrics Metrics
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	metrics "github.com/openshift/managed-upgrade-operator/pkg/metrics"
	gomock "go.uber.org/mock/gomock"
)

// MockMetrics is a mock of Metrics interface.
type MockMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsMockRecorder
}

// MockMetricsMockRecorder is the mock recorder for MockMetrics.
type MockMetricsMockRecorder struct {
	mock *MockMetrics
}

// NewMockMetrics creates a new mock instance.
func NewMockMetrics(ctrl *gomock.Controller) *MockMetrics {
	mock := &MockMetrics{ctrl: ctrl}
	mock.recorder = &MockMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetrics) EXPECT() *MockMetricsMockRecorder {
	return m.recorder
}

// AlertsFromUpgrade mocks base method.
func (m *MockMetrics) AlertsFromUpgrade(arg0, arg1 time.Time) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertsFromUpgrade", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertsFromUpgrade indicates an expected call of AlertsFromUpgrade.
func (mr *MockMetricsMockRecorder) AlertsFromUpgrade(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertsFromUpgrade", reflect.TypeOf((*MockMetrics)(nil).AlertsFromUpgrade), arg0, arg1)
}

// IsAlertFiring mocks base method.
func (m *MockMetrics) IsAlertFiring(arg0 string, arg1, arg2 []string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAlertFiring", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAlertFiring indicates an expected call of IsAlertFiring.
func (mr *MockMetricsMockRecorder) IsAlertFiring(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAlertFiring", reflect.TypeOf((*MockMetrics)(nil).IsAlertFiring), arg0, arg1, arg2)
}

// IsClusterVersionAtVersion mocks base method.
func (m *MockMetrics) IsClusterVersionAtVersion(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClusterVersionAtVersion", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsClusterVersionAtVersion indicates an expected call of IsClusterVersionAtVersion.
func (mr *MockMetricsMockRecorder) IsClusterVersionAtVersion(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClusterVersionAtVersion", reflect.TypeOf((*MockMetrics)(nil).IsClusterVersionAtVersion), arg0)
}

// IsMetricNotificationEventSentSet mocks base method.
func (m *MockMetrics) IsMetricNotificationEventSentSet(arg0, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMetricNotificationEventSentSet", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsMetricNotificationEventSentSet indicates an expected call of IsMetricNotificationEventSentSet.
func (mr *MockMetricsMockRecorder) IsMetricNotificationEventSentSet(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMetricNotificationEventSentSet", reflect.TypeOf((*MockMetrics)(nil).IsMetricNotificationEventSentSet), arg0, arg1, arg2)
}

// Query mocks base method.
func (m *MockMetrics) Query(arg0 string) (*metrics.AlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0)
	ret0, _ := ret[0].(*metrics.AlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockMetricsMockRecorder) Query(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockMetrics)(nil).Query), arg0)
}

// ResetAllMetricNodeDrainFailed mocks base method.
func (m *MockMetrics) ResetAllMetricNodeDrainFailed() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetAllMetricNodeDrainFailed")
}

// ResetAllMetricNodeDrainFailed indicates an expected call of ResetAllMetricNodeDrainFailed.
func (mr *MockMetricsMockRecorder) ResetAllMetricNodeDrainFailed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetAllMetricNodeDrainFailed", reflect.TypeOf((*MockMetrics)(nil).ResetAllMetricNodeDrainFailed))
}

// ResetEphemeralMetrics mocks base method.
func (m *MockMetrics) ResetEphemeralMetrics() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetEphemeralMetrics")
}

// ResetEphemeralMetrics indicates an expected call of ResetEphemeralMetrics.
func (mr *MockMetricsMockRecorder) ResetEphemeralMetrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetEphemeralMetrics", reflect.TypeOf((*MockMetrics)(nil).ResetEphemeralMetrics))
}

// ResetFailureMetrics mocks base method.
func (m *MockMetrics) ResetFailureMetrics() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetFailureMetrics")
}

// ResetFailureMetrics indicates an expected call of ResetFailureMetrics.
func (mr *MockMetricsMockRecorder) ResetFailureMetrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetFailureMetrics", reflect.TypeOf((*MockMetrics)(nil).ResetFailureMetrics))
}

// ResetMetricNodeDrainFailed mocks base method.
func (m *MockMetrics) ResetMetricNodeDrainFailed(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetMetricNodeDrainFailed", arg0)
}

// ResetMetricNodeDrainFailed indicates an expected call of ResetMetricNodeDrainFailed.
func (mr *MockMetricsMockRecorder) ResetMetricNodeDrainFailed(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetMetricNodeDrainFailed", reflect.TypeOf((*MockMetrics)(nil).ResetMetricNodeDrainFailed), arg0)
}

// ResetMetricUpgradeControlPlaneTimeout mocks base method.
func (m *MockMetrics) ResetMetricUpgradeControlPlaneTimeout(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetMetricUpgradeControlPlaneTimeout", arg0, arg1)
}

// ResetMetricUpgradeControlPlaneTimeout indicates an expected call of ResetMetricUpgradeControlPlaneTimeout.
func (mr *MockMetricsMockRecorder) ResetMetricUpgradeControlPlaneTimeout(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetMetricUpgradeControlPlaneTimeout", reflect.TypeOf((*MockMetrics)(nil).ResetMetricUpgradeControlPlaneTimeout), arg0, arg1)
}

// ResetMetricUpgradeWorkerTimeout mocks base method.
func (m *MockMetrics) ResetMetricUpgradeWorkerTimeout(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetMetricUpgradeWorkerTimeout", arg0, arg1)
}

// ResetMetricUpgradeWorkerTimeout indicates an expected call of ResetMetricUpgradeWorkerTimeout.
func (mr *MockMetricsMockRecorder) ResetMetricUpgradeWorkerTimeout(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetMetricUpgradeWorkerTimeout", reflect.TypeOf((*MockMetrics)(nil).ResetMetricUpgradeWorkerTimeout), arg0, arg1)
}

// UpdateMetricHealthcheckFailed mocks base method.
func (m *MockMetrics) UpdateMetricHealthcheckFailed(arg0, arg1, arg2, arg3 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateMetricHealthcheckFailed", arg0, arg1, arg2, arg3)
}

// UpdateMetricHealthcheckFailed indicates an expected call of UpdateMetricHealthcheckFailed.
func (mr *MockMetricsMockRecorder) UpdateMetricHealthcheckFailed(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetricHealthcheckFailed", reflect.TypeOf((*MockMetrics)(nil).UpdateMetricHealthcheckFailed), arg0, arg1, arg2, arg3)
}

// UpdateMetricHealthcheckSucceeded mocks base method.
func (m *MockMetrics) UpdateMetricHealthcheckSucceeded(arg0, arg1, arg2, arg3 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateMetricHealthcheckSucceeded", arg0, arg1, arg2, arg3)
}

// UpdateMetricHealthcheckSucceeded indicates an expected call of UpdateMetricHealthcheckSucceeded.
func (mr *MockMetricsMockRecorder) UpdateMetricHealthcheckSucceeded(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetricHealthcheckSucceeded", reflect.TypeOf((*MockMetrics)(nil).UpdateMetricHealthcheckSucceeded), arg0, arg1, arg2, arg3)
}

// UpdateMetricNodeDrainFailed mocks base method.
func (m *MockMetrics) UpdateMetricNodeDrainFailed(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateMetricNodeDrainFailed", arg0)
}

// UpdateMetricNodeDrainFailed indicates an expected call of UpdateMetricNodeDrainFailed.
func (mr *MockMetricsMockRecorder) UpdateMetricNodeDrainFailed(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetricNodeDrainFailed", reflect.TypeOf((*MockMetrics)(nil).UpdateMetricNodeDrainFailed), arg0)
}

// UpdateMetricNotificationEventSent mocks base method.
func (m *MockMetrics) UpdateMetricNotificationEventSent(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateMetricNotificationEventSent", arg0, arg1, arg2)
}

// UpdateMetricNotificationEventSent indicates an expected call of UpdateMetricNotificationEventSent.
func (mr *MockMetricsMockRecorder) UpdateMetricNotificationEventSent(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetricNotificationEventSent", reflect.TypeOf((*MockMetrics)(nil).UpdateMetricNotificationEventSent), arg0, arg1, arg2)
}

// UpdateMetricScalingFailed mocks base method.
func (m *MockMetrics) UpdateMetricScalingFailed(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateMetricScalingFailed", arg0)
}

// UpdateMetricScalingFailed indicates an expected call of UpdateMetricScalingFailed.
func (mr *MockMetricsMockRecorder) UpdateMetricScalingFailed(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetricScalingFailed", reflect.TypeOf((*MockMetrics)(nil).UpdateMetricScalingFailed), arg0)
}

// UpdateMetricScalingSucceeded mocks base method.
func (m *MockMetrics) UpdateMetricScalingSucceeded(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateMetricScalingSucceeded", arg0)
}

// UpdateMetricScalingSucceeded indicates an expected call of UpdateMetricScalingSucceeded.
func (mr *MockMetricsMockRecorder) UpdateMetricScalingSucceeded(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetricScalingSucceeded", reflect.TypeOf((*MockMetrics)(nil).UpdateMetricScalingSucceeded), arg0)
}

// UpdateMetricUpgradeConfigSyncTimestamp mocks base method.
func (m *MockMetrics) UpdateMetricUpgradeConfigSyncTimestamp(arg0 string, arg1 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateMetricUpgradeConfigSyncTimestamp", arg0, arg1)
}

// UpdateMetricUpgradeConfigSyncTimestamp indicates an expected call of UpdateMetricUpgradeConfigSyncTimestamp.
func (mr *MockMetricsMockRecorder) UpdateMetricUpgradeConfigSyncTimestamp(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetricUpgradeConfigSyncTimestamp", reflect.TypeOf((*MockMetrics)(nil).UpdateMetricUpgradeConfigSyncTimestamp), arg0, arg1)
}

// UpdateMetricUpgradeControlPlaneTimeout mocks base method.
func (m *MockMetrics) UpdateMetricUpgradeControlPlaneTimeout(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateMetricUpgradeControlPlaneTimeout", arg0, arg1)
}

// UpdateMetricUpgradeControlPlaneTimeout indicates an expected call of UpdateMetricUpgradeControlPlaneTimeout.
func (mr *MockMetricsMockRecorder) UpdateMetricUpgradeControlPlaneTimeout(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetricUpgradeControlPlaneTimeout", reflect.TypeOf((*MockMetrics)(nil).UpdateMetricUpgradeControlPlaneTimeout), arg0, arg1)
}

// UpdateMetricUpgradeResult mocks base method.
func (m *MockMetrics) UpdateMetricUpgradeResult(arg0, arg1, arg2, arg3 string, arg4 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateMetricUpgradeResult", arg0, arg1, arg2, arg3, arg4)
}

// UpdateMetricUpgradeResult indicates an expected call of UpdateMetricUpgradeResult.
func (mr *MockMetricsMockRecorder) UpdateMetricUpgradeResult(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetricUpgradeResult", reflect.TypeOf((*MockMetrics)(nil).UpdateMetricUpgradeResult), arg0, arg1, arg2, arg3, arg4)
}

// UpdateMetricUpgradeWindowBreached mocks base method.
func (m *MockMetrics) UpdateMetricUpgradeWindowBreached(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateMetricUpgradeWindowBreached", arg0)
}

// UpdateMetricUpgradeWindowBreached indicates an expected call of UpdateMetricUpgradeWindowBreached.
func (mr *MockMetricsMockRecorder) UpdateMetricUpgradeWindowBreached(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetricUpgradeWindowBreached", reflect.TypeOf((*MockMetrics)(nil).UpdateMetricUpgradeWindowBreached), arg0)
}

// UpdateMetricUpgradeWindowNotBreached mocks base method.
func (m *MockMetrics) UpdateMetricUpgradeWindowNotBreached(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateMetricUpgradeWindowNotBreached", arg0)
}

// UpdateMetricUpgradeWindowNotBreached indicates an expected call of UpdateMetricUpgradeWindowNotBreached.
func (mr *MockMetricsMockRecorder) UpdateMetricUpgradeWindowNotBreached(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetricUpgradeWindowNotBreached", reflect.TypeOf((*MockMetrics)(nil).UpdateMetricUpgradeWindowNotBreached), arg0)
}

// UpdateMetricUpgradeWorkerTimeout mocks base method.
func (m *MockMetrics) UpdateMetricUpgradeWorkerTimeout(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateMetricUpgradeWorkerTimeout", arg0, arg1)
}

// UpdateMetricUpgradeWorkerTimeout indicates an expected call of UpdateMetricUpgradeWorkerTimeout.
func (mr *MockMetricsMockRecorder) UpdateMetricUpgradeWorkerTimeout(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetricUpgradeWorkerTimeout", reflect.TypeOf((*MockMetrics)(nil).UpdateMetricUpgradeWorkerTimeout), arg0, arg1)
}

// UpdateMetricValidationFailed mocks base method.
func (m *MockMetrics) UpdateMetricValidationFailed(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateMetricValidationFailed", arg0)
}

// UpdateMetricValidationFailed indicates an expected call of UpdateMetricValidationFailed.
func (mr *MockMetricsMockRecorder) UpdateMetricValidationFailed(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetricValidationFailed", reflect.TypeOf((*MockMetrics)(nil).UpdateMetricValidationFailed), arg0)
}

// UpdateMetricValidationSucceeded mocks base method.
func (m *MockMetrics) UpdateMetricValidationSucceeded(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateMetricValidationSucceeded", arg0)
}

// UpdateMetricValidationSucceeded indicates an expected call of UpdateMetricValidationSucceeded.
func (mr *MockMetricsMockRecorder) UpdateMetricValidationSucceeded(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetricValidationSucceeded", reflect.TypeOf((*MockMetrics)(nil).UpdateMetricValidationSucceeded), arg0)
}

// UpdatemetricUpgradeNotificationFailed mocks base method.
func (m *MockMetrics) UpdatemetricUpgradeNotificationFailed(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatemetricUpgradeNotificationFailed", arg0, arg1)
}

// UpdatemetricUpgradeNotificationFailed indicates an expected call of UpdatemetricUpgradeNotificationFailed.
func (mr *MockMetricsMockRecorder) UpdatemetricUpgradeNotificationFailed(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatemetricUpgradeNotificationFailed", reflect.TypeOf((*MockMetrics)(nil).UpdatemetricUpgradeNotificationFailed), arg0, arg1)
}

// UpdatemetricUpgradeNotificationSucceeded mocks base method.
func (m *MockMetrics) UpdatemetricUpgradeNotificationSucceeded(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatemetricUpgradeNotificationSucceeded", arg0, arg1)
}

// UpdatemetricUpgradeNotificationSucceeded indicates an expected call of UpdatemetricUpgradeNotificationSucceeded.
func (mr *MockMetricsMockRecorder) UpdatemetricUpgradeNotificationSucceeded(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatemetricUpgradeNotificationSucceeded", reflect.TypeOf((*MockMetrics)(nil).UpdatemetricUpgradeNotificationSucceeded), arg0, arg1)
}
