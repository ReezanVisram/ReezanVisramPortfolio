package webhook_test

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	domain "reezanvisramportfolio/domain/project"
	mock_database "reezanvisramportfolio/internal/mocks/database"
	"reezanvisramportfolio/internal/webhook"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/mock/gomock"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type webhookServiceMock struct {
	logger            *slog.Logger
	projectRepository *mock_database.MockProjectRepository
	caser             cases.Caser
}

func newServiceMock(t *testing.T) webhookServiceMock {
	ctrl := gomock.NewController(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	projectRepo := mock_database.NewMockProjectRepository(ctrl)

	return webhookServiceMock{
		logger:            logger,
		projectRepository: projectRepo,
		caser:             cases.Title(language.BritishEnglish),
	}
}

func TestHandleStarWebhookCreated(t *testing.T) {
	repoName := "Sample-Project"
	repoId := int64(111111)
	repoDescription := "A sample project for unit testing"
	repoLink := "https://github.com/Sample-Project/"
	repoReleaseLink := "https://sampleproject.reezanvisram.com/"
	repoDefaultBranch := "main"
	repoTags := []string{"software", "cpp", "opengl"}

	sampleProject := domain.Project{
		Name:         "Sample Project",
		Id:           repoId,
		Description:  repoDescription,
		RepoLink:     repoLink,
		ReleaseLink:  repoReleaseLink,
		ImageLink:    fmt.Sprintf("%s/blob/%s/featured_screenshot.png", repoLink, repoDefaultBranch),
		IsHardware:   false,
		Technologies: []string{"C++", "OpenGL"},
	}

	randomErr := errors.New("random error")

	tests := map[string]struct {
		mocks       func() webhookServiceMock
		expectedErr error
	}{
		"successfully inserts a new project into the database": {
			mocks: func() webhookServiceMock {
				wsm := newServiceMock(t)
				wsm.projectRepository.EXPECT().GetProjectById(gomock.Any(), repoId).Return(nil, mongo.ErrNoDocuments)
				wsm.projectRepository.EXPECT().InsertProject(gomock.Any(), sampleProject).Return(nil)
				return wsm
			},
			expectedErr: nil,
		},
		"fails when trying to insert a project with an ID that already exists": {
			mocks: func() webhookServiceMock {
				wsm := newServiceMock(t)
				wsm.projectRepository.EXPECT().GetProjectById(gomock.Any(), repoId).Return(&sampleProject, nil)
				return wsm
			},
			expectedErr: webhook.ErrProjectExists,
		},
		"fails when unable to insert a project into the database": {
			mocks: func() webhookServiceMock {
				wsm := newServiceMock(t)
				wsm.projectRepository.EXPECT().GetProjectById(gomock.Any(), repoId).Return(nil, mongo.ErrNoDocuments)
				wsm.projectRepository.EXPECT().InsertProject(gomock.Any(), sampleProject).Return(randomErr)
				return wsm
			},
			expectedErr: randomErr,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			wsm := tt.mocks()
			ws := webhook.NewWebhookService(wsm.logger, wsm.projectRepository)
			ctx := context.Background()

			err := ws.HandleStarWebhookCreated(ctx,
				repoName,
				repoId,
				repoDescription,
				repoLink,
				repoReleaseLink,
				repoDefaultBranch,
				repoTags)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestHandleStarWebhookDeleted(t *testing.T) {
	repoId := int64(111111)

	sampleProject := domain.Project{
		Id: repoId,
	}

	randomErr := errors.New("random error")

	tests := map[string]struct {
		mocks       func() webhookServiceMock
		expectedErr error
	}{
		"successfully removes a project with the given id": {
			mocks: func() webhookServiceMock {
				wsm := newServiceMock(t)
				wsm.projectRepository.EXPECT().GetProjectById(gomock.Any(), repoId).Return(&sampleProject, nil)
				wsm.projectRepository.EXPECT().RemoveProjectById(gomock.Any(), repoId).Return(nil)
				return wsm
			},
			expectedErr: nil,
		},
		"fails when trying to delete a project that does not exist": {
			mocks: func() webhookServiceMock {
				wsm := newServiceMock(t)
				wsm.projectRepository.EXPECT().GetProjectById(gomock.Any(), repoId).Return(nil, mongo.ErrNoDocuments)
				return wsm
			},
			expectedErr: webhook.ErrProjectDoesNotExist,
		},
		"fails when unable to delete a project": {
			mocks: func() webhookServiceMock {
				wsm := newServiceMock(t)
				wsm.projectRepository.EXPECT().GetProjectById(gomock.Any(), repoId).Return(&sampleProject, nil)
				wsm.projectRepository.EXPECT().RemoveProjectById(gomock.Any(), repoId).Return(randomErr)
				return wsm
			},
			expectedErr: randomErr,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			wsm := tt.mocks()
			ws := webhook.NewWebhookService(wsm.logger, wsm.projectRepository)
			ctx := context.Background()

			err := ws.HandleStarWebhookDeleted(ctx, repoId)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
