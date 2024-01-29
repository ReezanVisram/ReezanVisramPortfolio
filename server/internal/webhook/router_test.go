package webhook_test

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	mock_webhook "reezanvisramportfolio/internal/mocks/webhook"
	"reezanvisramportfolio/internal/webhook"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type webhookRouterMock struct {
	logger         *slog.Logger
	webhookSecret  string
	webhookService *mock_webhook.MockWebhookService
}

func newRouterMock(t *testing.T) webhookRouterMock {
	ctrl := gomock.NewController(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	webhookService := mock_webhook.NewMockWebhookService(ctrl)

	return webhookRouterMock{
		logger:         logger,
		webhookSecret:  "secret",
		webhookService: webhookService,
	}
}

const webhookSecret = "secret"

func TestPostWebhookHandler(t *testing.T) {
	validStarWebhookCreatedRequestBody := webhook.StarWebhookRequest{
		Action: "created",
		Repository: webhook.StarWebhookRepositoryRequest{
			Name:      "Sample-Project",
			Id:        int64(111111),
			IsPrivate: false,
			Owner: webhook.StarWebhookOwnerRequest{
				Username: "ReezanVisram",
			},
			Description:   "Sample Project Description",
			RepoLink:      "https://github.com/reezanvisram/sample-project",
			ReleaseLink:   "https://sampleproject.reezanvisram.com",
			Tags:          []string{"software", "cpp", "opengl"},
			NumStars:      2,
			IsFork:        false,
			DefaultBranch: "main",
		},
		Sender: webhook.StarWebhookSenderRequest{
			Username: "ReezanVisram",
		},
	}

	validStarWebhookDeletedRequestBody := webhook.StarWebhookRequest{
		Action: "deleted",
		Repository: webhook.StarWebhookRepositoryRequest{
			Name:      "Sample-Project",
			Id:        int64(111111),
			IsPrivate: false,
			Owner: webhook.StarWebhookOwnerRequest{
				Username: "ReezanVisram",
			},
			Description:   "Sample Project Description",
			RepoLink:      "https://github.com/reezanvisram/sample-project",
			ReleaseLink:   "https://sampleproject.reezanvisram.com",
			Tags:          []string{"software", "cpp", "opengl"},
			NumStars:      2,
			IsFork:        false,
			DefaultBranch: "main",
		},
		Sender: webhook.StarWebhookSenderRequest{
			Username: "ReezanVisram",
		},
	}

	invalidSenderStarWebhookRequestBody := webhook.StarWebhookRequest{
		Action: "deleted",
		Repository: webhook.StarWebhookRepositoryRequest{
			Name:      "Sample-Project",
			Id:        int64(111111),
			IsPrivate: false,
			Owner: webhook.StarWebhookOwnerRequest{
				Username: "ReezanVisram",
			},
			Description:   "Sample Project Description",
			RepoLink:      "https://github.com/reezanvisram/sample-project",
			ReleaseLink:   "https://sampleproject.reezanvisram.com",
			Tags:          []string{"software", "cpp", "opengl"},
			NumStars:      2,
			IsFork:        false,
			DefaultBranch: "main",
		},
		Sender: webhook.StarWebhookSenderRequest{
			Username: "FarzanMirshekari",
		},
	}

	invalidOwnerStarWebhookRequestBody := webhook.StarWebhookRequest{
		Action: "created",
		Repository: webhook.StarWebhookRepositoryRequest{
			Name:      "Sample-Project",
			Id:        int64(111111),
			IsPrivate: false,
			Owner: webhook.StarWebhookOwnerRequest{
				Username: "TristanParry",
			},
			Description:   "Sample Project Description",
			RepoLink:      "https://github.com/reezanvisram/sample-project",
			ReleaseLink:   "https://sampleproject.reezanvisram.com",
			Tags:          []string{"software", "cpp", "opengl"},
			NumStars:      2,
			IsFork:        false,
			DefaultBranch: "main",
		},
		Sender: webhook.StarWebhookSenderRequest{
			Username: "ReezanVisram",
		},
	}

	privateRepoStarWebhookRequestBody := webhook.StarWebhookRequest{
		Action: "created",
		Repository: webhook.StarWebhookRepositoryRequest{
			Name:      "Sample-Project",
			Id:        int64(111111),
			IsPrivate: true,
			Owner: webhook.StarWebhookOwnerRequest{
				Username: "ReezanVisram",
			},
			Description:   "Sample Project Description",
			RepoLink:      "https://github.com/reezanvisram/sample-project",
			ReleaseLink:   "https://sampleproject.reezanvisram.com",
			Tags:          []string{"software", "cpp", "opengl"},
			NumStars:      2,
			IsFork:        false,
			DefaultBranch: "main",
		},
		Sender: webhook.StarWebhookSenderRequest{
			Username: "ReezanVisram",
		},
	}

	forkRepoStarWebhookRequestBody := webhook.StarWebhookRequest{
		Action: "created",
		Repository: webhook.StarWebhookRepositoryRequest{
			Name:      "Sample-Project",
			Id:        int64(111111),
			IsPrivate: false,
			Owner: webhook.StarWebhookOwnerRequest{
				Username: "ReezanVisram",
			},
			Description:   "Sample Project Description",
			RepoLink:      "https://github.com/reezanvisram/sample-project",
			ReleaseLink:   "https://sampleproject.reezanvisram.com",
			Tags:          []string{"software", "cpp", "opengl"},
			NumStars:      2,
			IsFork:        true,
			DefaultBranch: "main",
		},
		Sender: webhook.StarWebhookSenderRequest{
			Username: "ReezanVisram",
		},
	}

	validStarWebhookCreatedRequestBody2 := webhook.StarWebhookRequest{
		Action: "created",
		Repository: webhook.StarWebhookRepositoryRequest{
			Name:      "Sample-Project2",
			Id:        int64(111111),
			IsPrivate: false,
			Owner: webhook.StarWebhookOwnerRequest{
				Username: "ReezanVisram",
			},
			Description:   "Sample Project Description",
			RepoLink:      "https://github.com/reezanvisram/sample-project2",
			ReleaseLink:   "https://sampleproject2.reezanvisram.com",
			Tags:          []string{"software", "cpp", "opengl"},
			NumStars:      2,
			IsFork:        false,
			DefaultBranch: "main",
		},
		Sender: webhook.StarWebhookSenderRequest{
			Username: "ReezanVisram",
		},
	}

	validStarWebhookDeletedRequestBody2 := webhook.StarWebhookRequest{
		Action: "deleted",
		Repository: webhook.StarWebhookRepositoryRequest{
			Name:      "Sample-Project",
			Id:        int64(111111),
			IsPrivate: false,
			Owner: webhook.StarWebhookOwnerRequest{
				Username: "ReezanVisram",
			},
			Description:   "Sample Project Description",
			RepoLink:      "https://github.com/reezanvisram/sample-project",
			ReleaseLink:   "https://sampleproject.reezanvisram.com",
			Tags:          []string{"software", "cpp", "opengl"},
			NumStars:      2,
			IsFork:        false,
			DefaultBranch: "main",
		},
		Sender: webhook.StarWebhookSenderRequest{
			Username: "ReezanVisram",
		},
	}

	createdBody, _ := json.Marshal(validStarWebhookCreatedRequestBody)
	validStarWebhookCreatedRequest := httptest.NewRequest("POST", "/webhooks/", bytes.NewBuffer(createdBody))
	validStarWebhookCreatedRequest.Header.Add("X-Github-Event", "star")
	validStarWebhookCreatedRequest.Header.Add("X-Hub-Signature-256", calculateSignature(createdBody))

	missingSignatureWebhookRequest := httptest.NewRequest("POST", "/webhooks/", bytes.NewBuffer(createdBody))
	missingSignatureWebhookRequest.Header.Add("X-Github-Event", "star")

	invalidSignatureWebhookRequest := httptest.NewRequest("POST", "/webhooks/", bytes.NewBuffer(createdBody))
	invalidSignatureWebhookRequest.Header.Add("X-Github-Event", "star")
	invalidSignatureWebhookRequest.Header.Add("X-Hub-Signature-256", "invalid-signature")

	deletedBody, _ := json.Marshal(validStarWebhookDeletedRequestBody)
	validStarWebhookDeletedRequest := httptest.NewRequest("POST", "/webhooks/", bytes.NewBuffer(deletedBody))
	validStarWebhookDeletedRequest.Header.Add("X-Github-Event", "star")
	validStarWebhookDeletedRequest.Header.Add("X-Hub-Signature-256", calculateSignature(deletedBody))

	invalidSenderBody, _ := json.Marshal(invalidSenderStarWebhookRequestBody)
	invalidSenderStarWebhookRequest := httptest.NewRequest("POST", "/webhooks/", bytes.NewBuffer(invalidSenderBody))
	invalidSenderStarWebhookRequest.Header.Add("X-Github-Event", "star")
	invalidSenderStarWebhookRequest.Header.Add("X-Hub-Signature-256", calculateSignature(invalidSenderBody))

	invalidOwnerBody, _ := json.Marshal(invalidOwnerStarWebhookRequestBody)
	invalidOwnerStarWebhookRequest := httptest.NewRequest("POST", "/webhooks/", bytes.NewBuffer(invalidOwnerBody))
	invalidOwnerStarWebhookRequest.Header.Add("X-Github-Event", "star")
	invalidOwnerStarWebhookRequest.Header.Add("X-Hub-Signature-256", calculateSignature(invalidOwnerBody))

	invalidBodyStarWebhookRequest := httptest.NewRequest("POST", "/webhooks/", nil)
	invalidBodyStarWebhookRequest.Header.Add("X-Github-Event", "star")
	invalidBodyStarWebhookRequest.Header.Add("X-Hub-Signature-256", calculateSignature(nil))

	createdBody2, _ := json.Marshal(validStarWebhookCreatedRequestBody2)
	validStarWebhookCreatedRequest2 := httptest.NewRequest("POST", "/webhooks/", bytes.NewBuffer(createdBody2))
	validStarWebhookCreatedRequest2.Header.Add("X-Github-Event", "star")
	validStarWebhookCreatedRequest2.Header.Add("X-Hub-Signature-256", calculateSignature(createdBody2))

	deletedBody2, _ := json.Marshal(validStarWebhookDeletedRequestBody2)
	validStarWebhookDeletedRequest2 := httptest.NewRequest("POST", "/webhooks/", bytes.NewBuffer(deletedBody2))
	validStarWebhookDeletedRequest2.Header.Add("X-Github-Event", "star")
	validStarWebhookDeletedRequest2.Header.Add("X-Hub-Signature-256", calculateSignature(deletedBody2))

	privateRepoBody, _ := json.Marshal(privateRepoStarWebhookRequestBody)
	privateRepoStarWebhookRequest := httptest.NewRequest("POST", "/webhooks/", bytes.NewBuffer(privateRepoBody))
	privateRepoStarWebhookRequest.Header.Add("X-Github-Event", "star")
	privateRepoStarWebhookRequest.Header.Add("X-Hub-Signature-256", calculateSignature(privateRepoBody))

	forkRepoBody, _ := json.Marshal(forkRepoStarWebhookRequestBody)
	forkRepoStarWebhookRequest := httptest.NewRequest("POST", "/webhooks/", bytes.NewBuffer(forkRepoBody))
	forkRepoStarWebhookRequest.Header.Add("X-Github-Event", "star")
	forkRepoStarWebhookRequest.Header.Add("X-Hub-Signature-256", calculateSignature(forkRepoBody))

	randomErr := errors.New("random error")

	tests := map[string]struct {
		mocks          func() webhookRouterMock
		request        *http.Request
		expectedStatus int
	}{
		"returns 200 when receiving a valid star creation webhook request and is successfully able to insert project": {
			mocks: func() webhookRouterMock {
				wrm := newRouterMock(t)
				wrm.webhookService.EXPECT().HandleStarWebhookCreated(
					gomock.Any(),
					validStarWebhookCreatedRequestBody.Repository.Name,
					validStarWebhookCreatedRequestBody.Repository.Id,
					validStarWebhookCreatedRequestBody.Repository.Description,
					validStarWebhookCreatedRequestBody.Repository.RepoLink,
					validStarWebhookCreatedRequestBody.Repository.ReleaseLink,
					validStarWebhookCreatedRequestBody.Repository.DefaultBranch,
					validStarWebhookCreatedRequestBody.Repository.Tags,
				).Return(nil)
				return wrm
			},
			request:        validStarWebhookCreatedRequest,
			expectedStatus: http.StatusOK,
		},
		"returns 422 when a webhook request that does not have a signature": {
			mocks: func() webhookRouterMock {
				wrm := newRouterMock(t)
				return wrm
			},
			request:        missingSignatureWebhookRequest,
			expectedStatus: http.StatusUnprocessableEntity,
		},
		"returns 422 when a webhook request does not have a valid signature": {
			mocks: func() webhookRouterMock {
				wrm := newRouterMock(t)
				return wrm
			},
			request:        invalidSignatureWebhookRequest,
			expectedStatus: http.StatusUnprocessableEntity,
		},
		"returns 200 when receiving a valid star deleted webhook request, and is successfully able to delete project": {
			mocks: func() webhookRouterMock {
				wrm := newRouterMock(t)
				wrm.webhookService.EXPECT().HandleStarWebhookDeleted(
					gomock.Any(),
					validStarWebhookDeletedRequestBody.Repository.Id,
				).Return(nil)
				return wrm
			},
			request:        validStarWebhookDeletedRequest,
			expectedStatus: http.StatusOK,
		},
		"returns a 412 when receiving a star webhook request sent by someone other than ReezanVisram": {
			mocks: func() webhookRouterMock {
				wrm := newRouterMock(t)
				return wrm
			},
			request:        invalidSenderStarWebhookRequest,
			expectedStatus: http.StatusPreconditionFailed,
		},
		"returns a 412 when receiving a star webhook request on a repo owned by someone other than ReezanVisram": {
			mocks: func() webhookRouterMock {
				wrm := newRouterMock(t)
				return wrm
			},
			request:        invalidOwnerStarWebhookRequest,
			expectedStatus: http.StatusPreconditionFailed,
		},
		"returns a 400 when the star webhook request body is invalid": {
			mocks: func() webhookRouterMock {
				wrm := newRouterMock(t)
				return wrm
			},
			request:        invalidBodyStarWebhookRequest,
			expectedStatus: http.StatusBadRequest,
		},
		"returns a 500 when there is an issue inserting an otherwise valid star webhook creation request": {
			mocks: func() webhookRouterMock {
				wrm := newRouterMock(t)
				wrm.webhookService.EXPECT().HandleStarWebhookCreated(
					gomock.Any(),
					validStarWebhookCreatedRequestBody2.Repository.Name,
					validStarWebhookCreatedRequestBody2.Repository.Id,
					validStarWebhookCreatedRequestBody2.Repository.Description,
					validStarWebhookCreatedRequestBody2.Repository.RepoLink,
					validStarWebhookCreatedRequestBody2.Repository.ReleaseLink,
					validStarWebhookCreatedRequestBody2.Repository.DefaultBranch,
					validStarWebhookCreatedRequestBody2.Repository.Tags,
				).Return(randomErr)
				return wrm
			},
			request:        validStarWebhookCreatedRequest2,
			expectedStatus: http.StatusInternalServerError,
		},
		"returns a 500 when there is an issue deleting an otherwise valid star webhook deletion request": {
			mocks: func() webhookRouterMock {
				wrm := newRouterMock(t)
				wrm.webhookService.EXPECT().HandleStarWebhookDeleted(
					gomock.Any(),
					validStarWebhookDeletedRequestBody2.Repository.Id,
				).Return(randomErr)
				return wrm
			},
			request:        validStarWebhookDeletedRequest2,
			expectedStatus: http.StatusInternalServerError,
		},
		"returns a 412 when a star webhook request is received for a private repository": {
			mocks: func() webhookRouterMock {
				wrm := newRouterMock(t)
				return wrm
			},
			request:        privateRepoStarWebhookRequest,
			expectedStatus: http.StatusPreconditionFailed,
		},
		"returns a 412 when a star webhook request is received for a repository that is a fork": {
			mocks: func() webhookRouterMock {
				wrm := newRouterMock(t)
				return wrm
			},
			request:        forkRepoStarWebhookRequest,
			expectedStatus: http.StatusPreconditionFailed,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			wrm := tt.mocks()
			wr := webhook.NewWebhookRouter(wrm.logger, wrm.webhookSecret, wrm.webhookService)
			res := httptest.NewRecorder()
			wr.PostWebhookHandler(res, tt.request)
			assert.Equal(t, tt.expectedStatus, res.Code)
		})
	}
}

func calculateSignature(payload []byte) string {
	h := hmac.New(sha256.New, []byte(webhookSecret))

	h.Write(payload)

	return fmt.Sprintf("sha256=%s", hex.EncodeToString(h.Sum(nil)))
}
