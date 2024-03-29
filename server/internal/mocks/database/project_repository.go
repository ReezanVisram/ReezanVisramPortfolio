// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/database/project_repository.go
//
// Generated by this command:
//
//	mockgen -source=./internal/database/project_repository.go -destination=./internal/mocks/database/project_repository.go
//

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	domain "reezanvisramportfolio/domain/project"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockProjectRepository is a mock of ProjectRepository interface.
type MockProjectRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProjectRepositoryMockRecorder
}

// MockProjectRepositoryMockRecorder is the mock recorder for MockProjectRepository.
type MockProjectRepositoryMockRecorder struct {
	mock *MockProjectRepository
}

// NewMockProjectRepository creates a new mock instance.
func NewMockProjectRepository(ctrl *gomock.Controller) *MockProjectRepository {
	mock := &MockProjectRepository{ctrl: ctrl}
	mock.recorder = &MockProjectRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectRepository) EXPECT() *MockProjectRepositoryMockRecorder {
	return m.recorder
}

// GetAllProjects mocks base method.
func (m *MockProjectRepository) GetAllProjects(ctx context.Context) ([]domain.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProjects", ctx)
	ret0, _ := ret[0].([]domain.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProjects indicates an expected call of GetAllProjects.
func (mr *MockProjectRepositoryMockRecorder) GetAllProjects(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProjects", reflect.TypeOf((*MockProjectRepository)(nil).GetAllProjects), ctx)
}

// GetProjectById mocks base method.
func (m *MockProjectRepository) GetProjectById(ctx context.Context, id int64) (*domain.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectById", ctx, id)
	ret0, _ := ret[0].(*domain.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectById indicates an expected call of GetProjectById.
func (mr *MockProjectRepositoryMockRecorder) GetProjectById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectById", reflect.TypeOf((*MockProjectRepository)(nil).GetProjectById), ctx, id)
}

// InsertProject mocks base method.
func (m *MockProjectRepository) InsertProject(ctx context.Context, project domain.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertProject", ctx, project)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertProject indicates an expected call of InsertProject.
func (mr *MockProjectRepositoryMockRecorder) InsertProject(ctx, project any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertProject", reflect.TypeOf((*MockProjectRepository)(nil).InsertProject), ctx, project)
}

// RemoveProjectById mocks base method.
func (m *MockProjectRepository) RemoveProjectById(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveProjectById", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProjectById indicates an expected call of RemoveProjectById.
func (mr *MockProjectRepositoryMockRecorder) RemoveProjectById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProjectById", reflect.TypeOf((*MockProjectRepository)(nil).RemoveProjectById), ctx, id)
}
