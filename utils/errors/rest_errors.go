package errors

import "net/http"

type RestError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"code"`
	Error      string `json:"error"`
}

func NewBadRequest(message string, error string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      error,
	}
}
