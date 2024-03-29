// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/project/service.go
//
// Generated by this command:
//
//	mockgen -source=./internal/project/service.go -destination=./internal/mocks/project/service.go
//

// Package mock_project is a generated GoMock package.
package mock_project

import (
	context "context"
	domain "reezanvisramportfolio/domain/project"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockProjectService is a mock of ProjectService interface.
type MockProjectService struct {
	ctrl     *gomock.Controller
	recorder *MockProjectServiceMockRecorder
}

// MockProjectServiceMockRecorder is the mock recorder for MockProjectService.
type MockProjectServiceMockRecorder struct {
	mock *MockProjectService
}

// NewMockProjectService creates a new mock instance.
func NewMockProjectService(ctrl *gomock.Controller) *MockProjectService {
	mock := &MockProjectService{ctrl: ctrl}
	mock.recorder = &MockProjectServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectService) EXPECT() *MockProjectServiceMockRecorder {
	return m.recorder
}

// GetAllProjects mocks base method.
func (m *MockProjectService) GetAllProjects(ctx context.Context) ([]domain.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProjects", ctx)
	ret0, _ := ret[0].([]domain.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProjects indicates an expected call of GetAllProjects.
func (mr *MockProjectServiceMockRecorder) GetAllProjects(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProjects", reflect.TypeOf((*MockProjectService)(nil).GetAllProjects), ctx)
}
