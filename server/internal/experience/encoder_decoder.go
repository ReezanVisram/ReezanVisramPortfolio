package experience

import (
	"encoding/json"
	"net/http"
	domain "reezanvisramportfolio/domain/experience"
)

func encodeError(w http.ResponseWriter, err error) {
	errToWrite, _ := json.Marshal(Error{
		Error: err.Error(),
	})

	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(errToWrite)
}

func encodeResponse(w http.ResponseWriter, experience []domain.Experience) {
	body, err := json.Marshal(experience)
	if err != nil {
		encodeError(w, err)
		return
	}

	w.Write(body)
}

type Error struct {
	Error string `json:"error"`
}
