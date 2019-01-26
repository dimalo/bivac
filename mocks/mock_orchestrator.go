// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/orchestrators/orchestrators.go

// Package mocks is a generated GoMock package.
package mocks

import (
	volume "github.com/camptocamp/bivac/pkg/volume"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOrchestrator is a mock of Orchestrator interface
type MockOrchestrator struct {
	ctrl     *gomock.Controller
	recorder *MockOrchestratorMockRecorder
}

// MockOrchestratorMockRecorder is the mock recorder for MockOrchestrator
type MockOrchestratorMockRecorder struct {
	mock *MockOrchestrator
}

// NewMockOrchestrator creates a new mock instance
func NewMockOrchestrator(ctrl *gomock.Controller) *MockOrchestrator {
	mock := &MockOrchestrator{ctrl: ctrl}
	mock.recorder = &MockOrchestratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrchestrator) EXPECT() *MockOrchestratorMockRecorder {
	return m.recorder
}

// GetName mocks base method
func (m *MockOrchestrator) GetName() string {
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockOrchestratorMockRecorder) GetName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockOrchestrator)(nil).GetName))
}

// GetVolumes mocks base method
func (m *MockOrchestrator) GetVolumes(volumeFilters volume.Filters) ([]*volume.Volume, error) {
	ret := m.ctrl.Call(m, "GetVolumes", volumeFilters)
	ret0, _ := ret[0].([]*volume.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumes indicates an expected call of GetVolumes
func (mr *MockOrchestratorMockRecorder) GetVolumes(volumeFilters interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumes", reflect.TypeOf((*MockOrchestrator)(nil).GetVolumes), volumeFilters)
}

// DeployAgent mocks base method
func (m *MockOrchestrator) DeployAgent(image string, cmd, envs []string, volume *volume.Volume) (bool, string, error) {
	ret := m.ctrl.Call(m, "DeployAgent", image, cmd, envs, volume)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeployAgent indicates an expected call of DeployAgent
func (mr *MockOrchestratorMockRecorder) DeployAgent(image, cmd, envs, volume interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployAgent", reflect.TypeOf((*MockOrchestrator)(nil).DeployAgent), image, cmd, envs, volume)
}