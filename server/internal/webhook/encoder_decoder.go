package webhook

import (
	"encoding/json"
	"net/http"
)

func encodeError(w http.ResponseWriter, err error) {
	errToWrite, _ := json.Marshal(Error{
		Error: err.Error(),
	})

	switch err {
	case ErrInvalidSignature, ErrMissingSignature:
		w.WriteHeader(http.StatusUnprocessableEntity)
	case ErrCouldNotReadBody, ErrInvalidStarWebhookRequestBody:
		w.WriteHeader(http.StatusBadRequest)
	case ErrInvalidSender, ErrInvalidOwner, ErrRepoPrivate, ErrIsFork, ErrProjectExists, ErrProjectDoesNotExist:
		w.WriteHeader(http.StatusPreconditionFailed)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(errToWrite)
}

func decodeStarWebhookRequest(payload []byte) (*StarWebhookRequest, error) {
	starWebhookRequest := &StarWebhookRequest{}

	err := json.Unmarshal(payload, starWebhookRequest)
	if err != nil {
		return nil, ErrInvalidStarWebhookRequestBody
	}

	return starWebhookRequest, nil
}

type Error struct {
	Error string `json:"error"`
}
