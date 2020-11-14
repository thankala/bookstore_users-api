package errors

import "net/http"

type RestError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"code"`
	Error      string `json:"error"`
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      "Bad_Request",
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Error:      "Not_Found",
	}
}

func NewInternalError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Error:      "Internal_Server_Error",
	}
}
