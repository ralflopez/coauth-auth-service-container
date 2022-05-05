package utils

import (
	"encoding/json"
	"net/http"
)

type RequestError struct {
	Status  int			`json:"status"`
	Message string		`json:"message"`
}

func NewRequestError(status int, message string) *RequestError {
	return &RequestError{status, message}
}

func ThrowNotFoundException(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNotFound)
	exception := NewRequestError(http.StatusNotFound, message)
	json.NewEncoder(w).Encode(exception)
}

func ThrowForbiddenException(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusForbidden)
	exception := NewRequestError(http.StatusForbidden, message)
	json.NewEncoder(w).Encode(exception)
}

func ThrowBadRequestException(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	exception := NewRequestError(http.StatusBadRequest, message)
	json.NewEncoder(w).Encode(exception)
}
