package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

type WebhookRouter interface {
	PostWebhookHandler(w http.ResponseWriter, r *http.Request)
	validateWebhookSignature(payload []byte, receivedSignature string) error
	validateStarWebhookRequest(senderUsername string, ownerUsername string, isPrivate bool, isFork bool) error
}

type webhookRouter struct {
	webhookSecret  string
	webhookService WebhookService
}

func NewWebhookRouter(webhookSecret string, webhookService WebhookService) WebhookRouter {
	return &webhookRouter{
		webhookSecret:  webhookSecret,
		webhookService: webhookService,
	}
}

func (wr *webhookRouter) PostWebhookHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		encodeError(w, ErrCouldNotReadBody)
		return
	}

	err = wr.validateWebhookSignature(payload, r.Header.Get("X-Hub-Signature-256"))
	if err != nil {
		encodeError(w, err)
		return
	}

	if r.Header.Get("X-Github-Event") == "star" {
		starWebhookRequest, err := decodeStarWebhookRequest(payload)
		if err != nil {
			encodeError(w, err)
			return
		}

		err = wr.validateStarWebhookRequest(starWebhookRequest.Sender.Username,
			starWebhookRequest.Repository.Owner.Username,
			starWebhookRequest.Repository.IsPrivate,
			starWebhookRequest.Repository.IsFork)

		if err != nil {
			encodeError(w, err)
			return
		}

		if starWebhookRequest.Action == "created" {
			err := wr.webhookService.HandleStarWebhookCreated(starWebhookRequest.Repository.Name,
				starWebhookRequest.Repository.Description,
				starWebhookRequest.Repository.RepoLink,
				starWebhookRequest.Repository.ReleaseLink,
				starWebhookRequest.Repository.DefaultBranch,
				starWebhookRequest.Repository.Tags)
			if err != nil {
				encodeError(w, err)
				return
			}
		} else if starWebhookRequest.Action == "deleted" {
			err := wr.webhookService.HandleStarWebhookDeleted(starWebhookRequest.Repository.Name)
			if err != nil {
				encodeError(w, err)
				return
			}
		}
	}

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
