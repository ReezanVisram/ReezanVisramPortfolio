package message

import (
	"io"
	"log/slog"
	"net/http"
	"reezanvisramportfolio/internal/custom_logging"
)

type MessageRouter interface {
	PostMessageHandler(w http.ResponseWriter, r *http.Request)
}

type messageRouter struct {
	logger         *slog.Logger
	messageService MessageService
}

func NewMessageRouter(logger *slog.Logger, messageService MessageService) MessageRouter {
	return &messageRouter{
		logger:         logger,
		messageService: messageService,
	}
}

func (mr *messageRouter) PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	mr.logger.Info("messageRouter.PostMessageHandler", "path", "/message", "method", "POST", "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		mr.logger.Error("messageRouter.PostMessageHandler", "err", "could not read request body", "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))
		encodeError(w, ErrCouldNotReadBody)
		return
	}

	messageRequest, err := decodeMessageRequest(payload)
	if err != nil {
		mr.logger.Error("messageRouter.PostMessageHandler", "err", err.Error(), "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))
		encodeError(w, err)
		return
	}

	err = mr.messageService.HandleMessageCreated(r.Context(), messageRequest.Name, messageRequest.Email, messageRequest.Subject, messageRequest.Message, messageRequest.Token)
	if err != nil {
		mr.logger.Error("messageRouter.PostMessageHandler", "err", err.Error(), "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))
		encodeError(w, err)
		return
	}

	mr.logger.Info("messageRouter.PostMessageHandler", "outcome", "success", "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))
	w.WriteHeader(200)
}
