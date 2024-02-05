package project_test

import (
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	domain "reezanvisramportfolio/domain/project"
	mock_project "reezanvisramportfolio/internal/mocks/project"
	"reezanvisramportfolio/internal/project"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type projectRouterMock struct {
	logger         *slog.Logger
	projectService *mock_project.MockProjectService
}

func newRouterMock(t *testing.T) projectRouterMock {
	ctrl := gomock.NewController(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	projectService := mock_project.NewMockProjectService(ctrl)

	return projectRouterMock{
		logger:         logger,
		projectService: projectService,
	}
}

func Test_GetAllProjects(t *testing.T) {
	projects := []domain.Project{
		{
			Name: "Project 1",
		},
		{
			Name: "Project 2",
		},
	}

	randomErr := errors.New("random error")

	getAllProjectsRequest := httptest.NewRequest("GET", "/projects/", nil)

	tests := map[string]struct {
		mocks          func() projectRouterMock
		request        *http.Request
		expectedStatus int
	}{
		"returns 200 when successfully able to get all projects from service": {
			mocks: func() projectRouterMock {
				prm := newRouterMock(t)
				prm.projectService.EXPECT().GetAllProjects(gomock.Any()).Return(projects, nil)
				return prm
			},
			request:        getAllProjectsRequest,
			expectedStatus: http.StatusOK,
		},
		"returns 500 when unable to get all projects from service": {
			mocks: func() projectRouterMock {
				prm := newRouterMock(t)
				prm.projectService.EXPECT().GetAllProjects(gomock.Any()).Return(nil, randomErr)
				return prm
			},
			request:        getAllProjectsRequest,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			prm := tt.mocks()
			pr := project.NewProjectRouter(prm.logger, prm.projectService)
			res := httptest.NewRecorder()
			pr.GetProjects(res, tt.request)
			assert.Equal(t, tt.expectedStatus, res.Code)
		})
	}
}
