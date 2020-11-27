package errors

import "net/http"

// RestErr defines how the REST api should send errors.
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// NewBadRequestError creates a bad request error with given message
func NewBadRequestError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusBadRequest,
		Error:   "bad request",
	}
}

// NewNotFountError creates a not found error with given message
func NewNotFountError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusNotFound,
		Error:   "not found",
	}
}

// NewServerError creates a server error with given message
func NewServerError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusInternalServerError,
		Error:   "internal server error",
	}
}
