// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package mock_manager is a generated GoMock package.
package mock_manager

import (
	context "context"
	reflect "reflect"

	q "g.hz.netease.com/horizon/lib/q"
	models "g.hz.netease.com/horizon/pkg/application/models"
	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockManager) Create(ctx context.Context, application *models.Application, extraMembers map[string]string) (*models.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, application, extraMembers)
	ret0, _ := ret[0].(*models.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockManagerMockRecorder) Create(ctx, application, extraMembers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockManager)(nil).Create), ctx, application, extraMembers)
}

// DeleteByID mocks base method.
func (m *MockManager) DeleteByID(ctx context.Context, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockManagerMockRecorder) DeleteByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockManager)(nil).DeleteByID), ctx, id)
}

// GetByGroupIDs mocks base method.
func (m *MockManager) GetByGroupIDs(ctx context.Context, groupIDs []uint) ([]*models.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByGroupIDs", ctx, groupIDs)
	ret0, _ := ret[0].([]*models.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByGroupIDs indicates an expected call of GetByGroupIDs.
func (mr *MockManagerMockRecorder) GetByGroupIDs(ctx, groupIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByGroupIDs", reflect.TypeOf((*MockManager)(nil).GetByGroupIDs), ctx, groupIDs)
}

// GetByID mocks base method.
func (m *MockManager) GetByID(ctx context.Context, id uint) (*models.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*models.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockManagerMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockManager)(nil).GetByID), ctx, id)
}

// GetByIDs mocks base method.
func (m *MockManager) GetByIDs(ctx context.Context, ids []uint) ([]*models.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", ctx, ids)
	ret0, _ := ret[0].([]*models.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs.
func (mr *MockManagerMockRecorder) GetByIDs(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockManager)(nil).GetByIDs), ctx, ids)
}

// GetByName mocks base method.
func (m *MockManager) GetByName(ctx context.Context, name string) (*models.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", ctx, name)
	ret0, _ := ret[0].(*models.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockManagerMockRecorder) GetByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockManager)(nil).GetByName), ctx, name)
}

// GetByNameFuzzily mocks base method.
func (m *MockManager) GetByNameFuzzily(ctx context.Context, name string) ([]*models.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByNameFuzzily", ctx, name)
	ret0, _ := ret[0].([]*models.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByNameFuzzily indicates an expected call of GetByNameFuzzily.
func (mr *MockManagerMockRecorder) GetByNameFuzzily(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByNameFuzzily", reflect.TypeOf((*MockManager)(nil).GetByNameFuzzily), ctx, name)
}

// List mocks base method.
func (m *MockManager) List(ctx context.Context, groupIDs []uint, query *q.Query) (int, []*models.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, groupIDs, query)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]*models.Application)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockManagerMockRecorder) List(ctx, groupIDs, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockManager)(nil).List), ctx, groupIDs, query)
}

// Transfer mocks base method.
func (m *MockManager) Transfer(ctx context.Context, id, groupID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer", ctx, id, groupID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transfer indicates an expected call of Transfer.
func (mr *MockManagerMockRecorder) Transfer(ctx, id, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockManager)(nil).Transfer), ctx, id, groupID)
}

// UpdateByID mocks base method.
func (m *MockManager) UpdateByID(ctx context.Context, id uint, application *models.Application) (*models.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", ctx, id, application)
	ret0, _ := ret[0].(*models.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByID indicates an expected call of UpdateByID.
func (mr *MockManagerMockRecorder) UpdateByID(ctx, id, application interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockManager)(nil).UpdateByID), ctx, id, application)
}
