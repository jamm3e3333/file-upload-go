package error

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func NewBadRequestError(err error) *HttpError {
	return &HttpError{
		Code:    http.StatusBadRequest,
		Error:   "BAD_REQUEST",
		Message: fmt.Sprintf("Bad request: %s", err.Error()),
	}
}

func NewInternalError(err error) *HttpError {
	return &HttpError{
		Code:    http.StatusInternalServerError,
		Error:   "INTERNAL_SERVER_ERROR",
		Message: fmt.Sprintf("Internal server error: %s", err.Error()),
	}
}
