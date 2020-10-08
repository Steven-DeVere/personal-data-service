package errors

import "net/http"

// RestErr defines the rest error object
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// NewBadRequestError returns a bad request error
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// NewNotFoundError returns a bad "not found in the db" error
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

// NewNotFoundError returns a bad "not found in the db" error
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
