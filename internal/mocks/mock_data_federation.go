// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: DataFederationLister,DataFederationStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

// MockDataFederationLister is a mock of DataFederationLister interface.
type MockDataFederationLister struct {
	ctrl     *gomock.Controller
	recorder *MockDataFederationListerMockRecorder
}

// MockDataFederationListerMockRecorder is the mock recorder for MockDataFederationLister.
type MockDataFederationListerMockRecorder struct {
	mock *MockDataFederationLister
}

// NewMockDataFederationLister creates a new mock instance.
func NewMockDataFederationLister(ctrl *gomock.Controller) *MockDataFederationLister {
	mock := &MockDataFederationLister{ctrl: ctrl}
	mock.recorder = &MockDataFederationListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataFederationLister) EXPECT() *MockDataFederationListerMockRecorder {
	return m.recorder
}

// DataFederationList mocks base method.
func (m *MockDataFederationLister) DataFederationList(arg0 string) ([]admin.DataLakeTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataFederationList", arg0)
	ret0, _ := ret[0].([]admin.DataLakeTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataFederationList indicates an expected call of DataFederationList.
func (mr *MockDataFederationListerMockRecorder) DataFederationList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataFederationList", reflect.TypeOf((*MockDataFederationLister)(nil).DataFederationList), arg0)
}

// MockDataFederationStore is a mock of DataFederationStore interface.
type MockDataFederationStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataFederationStoreMockRecorder
}

// MockDataFederationStoreMockRecorder is the mock recorder for MockDataFederationStore.
type MockDataFederationStoreMockRecorder struct {
	mock *MockDataFederationStore
}

// NewMockDataFederationStore creates a new mock instance.
func NewMockDataFederationStore(ctrl *gomock.Controller) *MockDataFederationStore {
	mock := &MockDataFederationStore{ctrl: ctrl}
	mock.recorder = &MockDataFederationStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataFederationStore) EXPECT() *MockDataFederationStoreMockRecorder {
	return m.recorder
}

// DataFederation mocks base method.
func (m *MockDataFederationStore) DataFederation(arg0, arg1 string) (*admin.DataLakeTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataFederation", arg0, arg1)
	ret0, _ := ret[0].(*admin.DataLakeTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataFederation indicates an expected call of DataFederation.
func (mr *MockDataFederationStoreMockRecorder) DataFederation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataFederation", reflect.TypeOf((*MockDataFederationStore)(nil).DataFederation), arg0, arg1)
}

// DataFederationList mocks base method.
func (m *MockDataFederationStore) DataFederationList(arg0 string) ([]admin.DataLakeTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataFederationList", arg0)
	ret0, _ := ret[0].([]admin.DataLakeTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataFederationList indicates an expected call of DataFederationList.
func (mr *MockDataFederationStoreMockRecorder) DataFederationList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataFederationList", reflect.TypeOf((*MockDataFederationStore)(nil).DataFederationList), arg0)
}
