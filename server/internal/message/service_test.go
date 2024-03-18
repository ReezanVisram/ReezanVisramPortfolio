package message_test

import (
	"context"
	"errors"
	"log/slog"
	"os"
	domain "reezanvisramportfolio/domain/message"
	"reezanvisramportfolio/internal/message"
	mock_adapters "reezanvisramportfolio/internal/mocks/adapters/adapters"
	mock_database "reezanvisramportfolio/internal/mocks/database"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type messageServiceMock struct {
	logger            *slog.Logger
	messageRepository *mock_database.MockMessageRepository
	recaptchaClient   *mock_adapters.MockRecaptchaClient
}

func newServiceMock(t *testing.T) messageServiceMock {
	ctrl := gomock.NewController(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	messageRepo := mock_database.NewMockMessageRepository(ctrl)
	recaptchaClient := mock_adapters.NewMockRecaptchaClient(ctrl)

	return messageServiceMock{
		logger:            logger,
		messageRepository: messageRepo,
		recaptchaClient:   recaptchaClient,
	}
}

func TestHandleMessageCreateed(t *testing.T) {
	mockName := "Test Name"
	mockEmail := "test@email.com"
	mockSubject := "Test Subject"
	mockMessageBody := "Test Message"

	mockMessage := domain.Message{
		Name:       mockName,
		Email:      mockEmail,
		Subject:    mockSubject,
		Message:    mockMessageBody,
		HasAlerted: false,
	}

	mockToken := "mocktoken"

	randomErr := errors.New("random error")

	tests := map[string]struct {
		mocks       func() messageServiceMock
		expectedErr error
	}{
		"successfully creates message when user is human": {
			mocks: func() messageServiceMock {
				msm := newServiceMock(t)
				msm.recaptchaClient.EXPECT().VerifyToken(gomock.Any(), mockToken).Return(true, nil)
				msm.messageRepository.EXPECT().InsertMessage(gomock.Any(), mockMessage).Return(nil)
				return msm
			},
		},
		"does not insert message when unable to verify is user is human": {
			mocks: func() messageServiceMock {
				msm := newServiceMock(t)
				msm.recaptchaClient.EXPECT().VerifyToken(gomock.Any(), mockToken).Return(true, randomErr)
				return msm
			},
			expectedErr: message.ErrUnableToVerify,
		},
		"does not insert message when sender is not human": {
			mocks: func() messageServiceMock {
				msm := newServiceMock(t)
				msm.recaptchaClient.EXPECT().VerifyToken(gomock.Any(), mockToken).Return(false, nil)
				return msm
			},
			expectedErr: message.ErrIsBot,
		},
		"returns an error when unable to insert message": {
			mocks: func() messageServiceMock {
				msm := newServiceMock(t)
				msm.recaptchaClient.EXPECT().VerifyToken(gomock.Any(), mockToken).Return(true, nil)
				msm.messageRepository.EXPECT().InsertMessage(gomock.Any(), mockMessage).Return(randomErr)
				return msm
			},
			expectedErr: message.ErrUnableToInsertMessage,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			msm := tt.mocks()
			ms := message.NewMessageService(msm.logger, msm.messageRepository, msm.recaptchaClient)
			ctx := context.Background()

			err := ms.HandleMessageCreated(ctx, mockName, mockEmail, mockSubject, mockMessageBody, mockToken)
			assert.Equal(t, tt.expectedErr, err)
		})
	}

}
