package experience_test

import (
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	domain "reezanvisramportfolio/domain/experience"
	"reezanvisramportfolio/internal/experience"
	mock_experience "reezanvisramportfolio/internal/mocks/experience"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type experienceRouterMock struct {
	logger            *slog.Logger
	experienceService *mock_experience.MockExperienceService
}

func newRouterMock(t *testing.T) experienceRouterMock {
	ctrl := gomock.NewController(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	experienceService := mock_experience.NewMockExperienceService(ctrl)

	return experienceRouterMock{
		logger:            logger,
		experienceService: experienceService,
	}
}

func Test_GetExperience(t *testing.T) {
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

	getExperienceRequest := httptest.NewRequest("GET", "/experience/", nil)

	tests := map[string]struct {
		mocks          func() experienceRouterMock
		request        *http.Request
		expectedStatus int
	}{
		"returns 200 when successfully able to get experience from service": {
			mocks: func() experienceRouterMock {
				erm := newRouterMock(t)
				erm.experienceService.EXPECT().GetExperience(gomock.Any()).Return(experiences, nil)
				return erm
			},
			request:        getExperienceRequest,
			expectedStatus: http.StatusOK,
		},
		"returns 500 when unable to get experience from service": {
			mocks: func() experienceRouterMock {
				erm := newRouterMock(t)
				erm.experienceService.EXPECT().GetExperience(gomock.Any()).Return(nil, randomErr)
				return erm
			},
			request:        getExperienceRequest,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			erm := tt.mocks()
			er := experience.NewExperienceRouter(erm.logger, erm.experienceService)
			res := httptest.NewRecorder()
			er.GetExperience(res, tt.request)
			assert.Equal(t, tt.expectedStatus, res.Code)
		})
	}

}
