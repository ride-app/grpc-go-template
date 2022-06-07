// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ride-app/entity-service/repositories/entity (interfaces: EntityRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	entityv1alpha1 "github.com/ride-app/entity-service/api/gen/ride/entity/v1alpha1"
)

// MockEntityRepository is a mock of EntityRepository interface.
type MockEntityRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEntityRepositoryMockRecorder
}

// MockEntityRepositoryMockRecorder is the mock recorder for MockEntityRepository.
type MockEntityRepositoryMockRecorder struct {
	mock *MockEntityRepository
}

// NewMockEntityRepository creates a new mock instance.
func NewMockEntityRepository(ctrl *gomock.Controller) *MockEntityRepository {
	mock := &MockEntityRepository{ctrl: ctrl}
	mock.recorder = &MockEntityRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEntityRepository) EXPECT() *MockEntityRepositoryMockRecorder {
	return m.recorder
}

// CreateEntity mocks base method.
func (m *MockEntityRepository) CreateEntity(arg0 context.Context, arg1 *entityv1alpha1.Entity) (*time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntity", arg0, arg1)
	ret0, _ := ret[0].(*time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntity indicates an expected call of CreateEntity.
func (mr *MockEntityRepositoryMockRecorder) CreateEntity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntity", reflect.TypeOf((*MockEntityRepository)(nil).CreateEntity), arg0, arg1)
}

// DeleteEntity mocks base method.
func (m *MockEntityRepository) DeleteEntity(arg0 context.Context, arg1 string) (*time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEntity", arg0, arg1)
	ret0, _ := ret[0].(*time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEntity indicates an expected call of DeleteEntity.
func (mr *MockEntityRepositoryMockRecorder) DeleteEntity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEntity", reflect.TypeOf((*MockEntityRepository)(nil).DeleteEntity), arg0, arg1)
}

// GetEntity mocks base method.
func (m *MockEntityRepository) GetEntity(arg0 context.Context, arg1 string) (*entityv1alpha1.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntity", arg0, arg1)
	ret0, _ := ret[0].(*entityv1alpha1.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntity indicates an expected call of GetEntity.
func (mr *MockEntityRepositoryMockRecorder) GetEntity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntity", reflect.TypeOf((*MockEntityRepository)(nil).GetEntity), arg0, arg1)
}

// GetEntitys mocks base method.
func (m *MockEntityRepository) GetEntitys(arg0 context.Context, arg1 string) ([]*entityv1alpha1.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntitys", arg0, arg1)
	ret0, _ := ret[0].([]*entityv1alpha1.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntitys indicates an expected call of GetEntitys.
func (mr *MockEntityRepositoryMockRecorder) GetEntitys(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntitys", reflect.TypeOf((*MockEntityRepository)(nil).GetEntitys), arg0, arg1)
}

// UpdateEntity mocks base method.
func (m *MockEntityRepository) UpdateEntity(arg0 context.Context, arg1 *entityv1alpha1.Entity) (*time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEntity", arg0, arg1)
	ret0, _ := ret[0].(*time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEntity indicates an expected call of UpdateEntity.
func (mr *MockEntityRepositoryMockRecorder) UpdateEntity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntity", reflect.TypeOf((*MockEntityRepository)(nil).UpdateEntity), arg0, arg1)
}
