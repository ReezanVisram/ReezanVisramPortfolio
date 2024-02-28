package message

import (
	"encoding/json"
	"net/http"
)

func encodeError(w http.ResponseWriter, err error) {
	errToWrite, _ := json.Marshal(Error{
		Error: err.Error(),
	})

	switch err {
	case ErrCouldNotReadBody, ErrInvalidMessageRequestBody:
		w.WriteHeader(http.StatusBadRequest)
	case ErrUnableToVerify, ErrIsBot:
		w.WriteHeader(http.StatusForbidden)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(errToWrite)
}

func decodeMessageRequest(payload []byte) (*MessageRequest, error) {
	messageRequest := &MessageRequest{}

	err := json.Unmarshal(payload, messageRequest)

	if err != nil {
		return nil, ErrInvalidMessageRequestBody
	}

	return messageRequest, nil
}

type Error struct {
	Error string `json:"error"`
}
