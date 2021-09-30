// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package mock_templaterelease is a generated GoMock package.
package mock_templaterelease

import (
	context "context"
	models "g.hz.netease.com/horizon/pkg/templaterelease/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockManager) Create(ctx context.Context, user *models.TemplateRelease) (*models.TemplateRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(*models.TemplateRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockManagerMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockManager)(nil).Create), ctx, user)
}

// ListByTemplateName mocks base method
func (m *MockManager) ListByTemplateName(ctx context.Context, templateName string) ([]models.TemplateRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByTemplateName", ctx, templateName)
	ret0, _ := ret[0].([]models.TemplateRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByTemplateName indicates an expected call of ListByTemplateName
func (mr *MockManagerMockRecorder) ListByTemplateName(ctx, templateName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByTemplateName", reflect.TypeOf((*MockManager)(nil).ListByTemplateName), ctx, templateName)
}

// GetByTemplateNameAndRelease mocks base method
func (m *MockManager) GetByTemplateNameAndRelease(ctx context.Context, templateName, release string) (*models.TemplateRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTemplateNameAndRelease", ctx, templateName, release)
	ret0, _ := ret[0].(*models.TemplateRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTemplateNameAndRelease indicates an expected call of GetByTemplateNameAndRelease
func (mr *MockManagerMockRecorder) GetByTemplateNameAndRelease(ctx, templateName, release interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTemplateNameAndRelease", reflect.TypeOf((*MockManager)(nil).GetByTemplateNameAndRelease), ctx, templateName, release)
}