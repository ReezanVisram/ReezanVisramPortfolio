package project

import (
	"encoding/json"
	"net/http"
	domain "reezanvisramportfolio/domain/project"
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

func encodeResponse(w http.ResponseWriter, projects []domain.Project) {
	body, err := json.Marshal(projects)
	if err != nil {
		encodeError(w, err)
		return
	}

	w.Write(body)
}

type Error struct {
	Error string `json:"error"`
}
