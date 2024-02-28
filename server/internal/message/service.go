package message

import (
	"context"
	"log/slog"
	domain "reezanvisramportfolio/domain/message"
	"reezanvisramportfolio/internal/adapters"
	"reezanvisramportfolio/internal/custom_logging"
	"reezanvisramportfolio/internal/database"
)

type MessageService interface {
	HandleMessageCreated(
		ctx context.Context,
		name string,
		email string,
		subject string,
		message string,
		token string) error
}

type messageService struct {
	logger            *slog.Logger
	messageRepository database.MessageRepository
	recaptchaClient   adapters.RecaptchaClient
}

func NewMessageService(logger *slog.Logger, messageRepository database.MessageRepository, recaptchaClient adapters.RecaptchaClient) MessageService {
	return &messageService{
		logger:            logger,
		messageRepository: messageRepository,
		recaptchaClient:   recaptchaClient,
	}
}

func (ms *messageService) HandleMessageCreated(ctx context.Context, name string, email string, subject string, message string, token string) error {
	ms.logger.Info("messageService.HandleMessageCreated", "name", name, "email", email, "subject", subject, "message", message, "correlation_id", ctx.Value(custom_logging.KeyCorrelationId))

	isHuman, err := ms.recaptchaClient.VerifyToken(ctx, token)
	if err != nil {
		ms.logger.Error("messageService.HandleMessageCreated", "err", err.Error(), "correlation_id", ctx.Value(custom_logging.KeyCorrelationId))
		return ErrUnableToVerify
	}

	if !isHuman {
		ms.logger.Error("messageService.HandleMessageCreated", "err", ErrIsBot.Error(), "correlation_id", ctx.Value(custom_logging.KeyCorrelationId))
		return ErrIsBot
	}

	newMessage := domain.Message{
		Name:       name,
		Email:      email,
		Subject:    subject,
		Message:    message,
		HasAlerted: false,
	}

	err = ms.messageRepository.InsertMessage(ctx, newMessage)
	if err != nil {
		ms.logger.Error("messageService.HandleMessageCreated", "err", err.Error(), "correlation_id", ctx.Value(custom_logging.KeyCorrelationId))
		return ErrUnableToInsertMessage
	}

	// TODO: Figure out a system to send emails for messages (maybe batch them and send them once a day)

	return nil
}
