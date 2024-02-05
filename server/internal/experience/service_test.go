package experience_test

import (
	"context"
	"errors"
	"log/slog"
	"os"
	domain "reezanvisramportfolio/domain/experience"
	"reezanvisramportfolio/internal/experience"
	mock_database "reezanvisramportfolio/internal/mocks/database"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type experienceServiceMock struct {
	logger               *slog.Logger
	experienceRepository *mock_database.MockExperienceRepository
}

func newServiceMock(t *testing.T) experienceServiceMock {
	ctrl := gomock.NewController(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	experienceRepo := mock_database.NewMockExperienceRepository(ctrl)

	return experienceServiceMock{
		logger:               logger,
		experienceRepository: experienceRepo,
	}
}

func TestGetExperience(t *testing.T) {
	experiences := []domain.Experience{
		{
			Name:            "sample company 1",
			StartAndEndDate: "time 1 - time 2",
			JobTitle:        "job title 1",
		},
		{
			Name:            "sample company 2",
			StartAndEndDate: "time 3 - time 4",
			JobTitle:        "job title 2",
		},
	}

	randomErr := errors.New("random error")

	tests := map[string]struct {
		mocks       func() experienceServiceMock
		expectedErr error
	}{
		"successfullly gets experience": {
			mocks: func() experienceServiceMock {
				esm := newServiceMock(t)
				esm.experienceRepository.EXPECT().GetExperience(gomock.Any()).Return(experiences, nil)
				return esm
			},
			expectedErr: nil,
		},
		"fails when unable to get experience from the database": {
			mocks: func() experienceServiceMock {
				esm := newServiceMock(t)
				esm.experienceRepository.EXPECT().GetExperience(gomock.Any()).Return(nil, randomErr)
				return esm
			},
			expectedErr: randomErr,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			esm := tt.mocks()
			es := experience.NewExperienceService(esm.logger, esm.experienceRepository)
			ctx := context.Background()

			_, err := es.GetExperience(ctx)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
