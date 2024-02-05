package project_test

import (
	"context"
	"errors"
	"log/slog"
	"os"
	domain "reezanvisramportfolio/domain/project"
	mock_database "reezanvisramportfolio/internal/mocks/database"
	"reezanvisramportfolio/internal/project"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type projectServiceMock struct {
	logger            *slog.Logger
	projectRepository *mock_database.MockProjectRepository
}

func newServiceMock(t *testing.T) projectServiceMock {
	ctrl := gomock.NewController(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	projectRepo := mock_database.NewMockProjectRepository(ctrl)

	return projectServiceMock{
		logger:            logger,
		projectRepository: projectRepo,
	}
}

func TestGetAllProjects(t *testing.T) {
	projects := []domain.Project{
		{
			Name: "Sample Project 1",
		},
		{
			Name: "Sample Project 2",
		},
	}

	randomErr := errors.New("random error")

	tests := map[string]struct {
		mocks       func() projectServiceMock
		expectedErr error
	}{
		"successfully gets all projects": {
			mocks: func() projectServiceMock {
				psm := newServiceMock(t)
				psm.projectRepository.EXPECT().GetAllProjects(gomock.Any()).Return(projects, nil)
				return psm
			},
			expectedErr: nil,
		},
		"fails when unable to get all projects from the database": {
			mocks: func() projectServiceMock {
				psm := newServiceMock(t)
				psm.projectRepository.EXPECT().GetAllProjects(gomock.Any()).Return(nil, randomErr)
				return psm
			},
			expectedErr: randomErr,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			psm := tt.mocks()
			ps := project.NewProjectService(psm.logger, psm.projectRepository)
			ctx := context.Background()

			_, err := ps.GetAllProjects(ctx)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
