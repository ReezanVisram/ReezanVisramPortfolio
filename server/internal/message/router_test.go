package message_test

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"reezanvisramportfolio/internal/message"
	mock_message "reezanvisramportfolio/internal/mocks/message"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type messageRouterMock struct {
	logger         *slog.Logger
	messageService *mock_message.MockMessageService
}

func newRouterMock(t *testing.T) messageRouterMock {
	ctrl := gomock.NewController(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	messageService := mock_message.NewMockMessageService(ctrl)

	return messageRouterMock{
		logger:         logger,
		messageService: messageService,
	}
}

func Test_PostMessageHandler(t *testing.T) {
	validMessageRequestBody := message.MessageRequest{
		Name:    "Test Name",
		Email:   "test@email.com",
		Subject: "Test Subject",
		Message: "Test Message",
		Token:   "mocktoken",
	}

	invalidMessageRequestBody := message.MessageRequest{}

	validBody, _ := json.Marshal(validMessageRequestBody)
	validMessageRequest := httptest.NewRequest("POST", "/message/", bytes.NewBuffer(validBody))

	validBody2, _ := json.Marshal(validMessageRequestBody)
	validMessageRequest2 := httptest.NewRequest("POST", "/message/", bytes.NewBuffer(validBody2))

	validBody3, _ := json.Marshal(validMessageRequestBody)
	validMessageRequest3 := httptest.NewRequest("POST", "/message/", bytes.NewBuffer(validBody3))

	invalidBody, _ := json.Marshal(invalidMessageRequestBody)
	invalidMessageRequest := httptest.NewRequest("POST", "/message/", bytes.NewBuffer(invalidBody))

	tests := map[string]struct {
		mocks          func() messageRouterMock
		request        *http.Request
		expectedStatus int
	}{
		"returns 200 when receiving a valid message request and is able to successfully insert message": {
			mocks: func() messageRouterMock {
				mrm := newRouterMock(t)
				mrm.messageService.EXPECT().HandleMessageCreated(
					gomock.Any(),
					validMessageRequestBody.Name,
					validMessageRequestBody.Email,
					validMessageRequestBody.Subject,
					validMessageRequestBody.Message,
					validMessageRequestBody.Token,
				).Return(nil)
				return mrm
			},
			request:        validMessageRequest,
			expectedStatus: 200,
		},
		"returns a 403 when unable to verify if message came from a bot": {
			mocks: func() messageRouterMock {
				mrm := newRouterMock(t)
				mrm.messageService.EXPECT().HandleMessageCreated(
					gomock.Any(),
					validMessageRequestBody.Name,
					validMessageRequestBody.Email,
					validMessageRequestBody.Subject,
					validMessageRequestBody.Message,
					validMessageRequestBody.Token,
				).Return(message.ErrUnableToVerify)
				return mrm
			},
			request:        validMessageRequest2,
			expectedStatus: 403,
		},
		"returns a 403 when message came from a bot": {
			mocks: func() messageRouterMock {
				mrm := newRouterMock(t)
				mrm.messageService.EXPECT().HandleMessageCreated(
					gomock.Any(),
					validMessageRequestBody.Name,
					validMessageRequestBody.Email,
					validMessageRequestBody.Subject,
					validMessageRequestBody.Message,
					validMessageRequestBody.Token,
				).Return(message.ErrIsBot)
				return mrm
			},
			request:        validMessageRequest3,
			expectedStatus: 403,
		},
		"returns a 400 when mesasge is invalid": {
			mocks: func() messageRouterMock {
				mrm := newRouterMock(t)
				return mrm
			},
			request:        invalidMessageRequest,
			expectedStatus: 400,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mrm := tt.mocks()
			mr := message.NewMessageRouter(mrm.logger, mrm.messageService)
			res := httptest.NewRecorder()
			mr.PostMessageHandler(res, tt.request)
			assert.Equal(t, tt.expectedStatus, res.Code)
		})
	}
}
