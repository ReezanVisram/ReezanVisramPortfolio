package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"reezanvisramportfolio/internal/custom_logging"
	"strconv"
)

type WebhookRouter interface {
	PostWebhookHandler(w http.ResponseWriter, r *http.Request)
	validateWebhookSignature(payload []byte, receivedSignature string) error
	validateStarWebhookRequest(senderUsername string, ownerUsername string, isPrivate bool, isFork bool) error
}

type webhookRouter struct {
	logger         *slog.Logger
	webhookSecret  string
	webhookService WebhookService
}

func NewWebhookRouter(logger *slog.Logger, webhookSecret string, webhookService WebhookService) WebhookRouter {
	return &webhookRouter{
		logger:         logger,
		webhookSecret:  webhookSecret,
		webhookService: webhookService,
	}
}

func (wr *webhookRouter) PostWebhookHandler(w http.ResponseWriter, r *http.Request) {
	wr.logger.Info("webhookRouter.PostWebhookHandler", "path", "/webhooks/", "method", "POST", "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		encodeError(w, ErrCouldNotReadBody)
		return
	}

	err = wr.validateWebhookSignature(payload, r.Header.Get("X-Hub-Signature-256"))
	if err != nil {
		wr.logger.Error("webhookRouter.PostWebhookHandler", "signature", r.Header.Get("X-Hub-Signature-256"), "err", err.Error(), "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))
		encodeError(w, err)
		return
	}

	if r.Header.Get("X-Github-Event") == "star" {
		wr.logger.Info("webhookRouter.PostWebhookHandler", "event", "star", "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))
		starWebhookRequest, err := decodeStarWebhookRequest(payload)
		if err != nil {
			wr.logger.Error("webhookRouter.PostWebhookHandler", "err", err.Error(), "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))
			encodeError(w, err)
			return
		}

		err = wr.validateStarWebhookRequest(starWebhookRequest.Sender.Username,
			starWebhookRequest.Repository.Owner.Username,
			starWebhookRequest.Repository.IsPrivate,
			starWebhookRequest.Repository.IsFork)

		if err != nil {
			wr.logger.Error("webhookRouter.PostWebhookHandler", "err", err.Error(), "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))
			encodeError(w, err)
			return
		}

		if starWebhookRequest.Action == "created" {
			wr.logger.Info("webhookRouter.PostWebhookHandler", "action", "created", "repository_name", starWebhookRequest.Repository.Name, "repository_id", strconv.FormatInt(starWebhookRequest.Repository.Id, 10), "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))
			err := wr.webhookService.HandleStarWebhookCreated(
				r.Context(),
				starWebhookRequest.Repository.Name,
				starWebhookRequest.Repository.Id,
				starWebhookRequest.Repository.Description,
				starWebhookRequest.Repository.RepoLink,
				starWebhookRequest.Repository.ReleaseLink,
				starWebhookRequest.Repository.DefaultBranch,
				starWebhookRequest.Repository.Tags)
			if err != nil {
				wr.logger.Error("webhookRouter.PostWebhookHandler", "err", err.Error(), "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))
				encodeError(w, err)
				return
			}
		} else if starWebhookRequest.Action == "deleted" {
			wr.logger.Info("webhookRouter.PostWebhookHandler", "action", "deleted", "repository_name", starWebhookRequest.Repository.Name, "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))
			err := wr.webhookService.HandleStarWebhookDeleted(r.Context(), starWebhookRequest.Repository.Id)
			if err != nil {
				wr.logger.Error("webhookRouter.PostWebhookHandler", "err", err.Error(), "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))
				encodeError(w, err)
				return
			}
		}
	}

	wr.logger.Info("webhookRouter.PostWebhookHandler", "outcome", "success", "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))
	w.WriteHeader(200)
}

func (wr *webhookRouter) validateWebhookSignature(payload []byte, receivedSignature string) error {
	if receivedSignature == "" {
		return ErrMissingSignature
	}

	h := hmac.New(sha256.New, []byte(wr.webhookSecret))

	h.Write(payload)

	expectedSignature := fmt.Sprintf("sha256=%s", hex.EncodeToString(h.Sum(nil)))

	if expectedSignature != receivedSignature {
		return ErrInvalidSignature
	}

	return nil
}

func (wr *webhookRouter) validateStarWebhookRequest(senderUsername string, ownerUsername string, isPrivate bool, isFork bool) error {
	if senderUsername != "ReezanVisram" {
		return ErrInvalidSender
	}

	if ownerUsername != "ReezanVisram" {
		return ErrInvalidOwner
	}

	if isPrivate {
		return ErrRepoPrivate
	}

	if isFork {
		return ErrIsFork
	}

	return nil
}
