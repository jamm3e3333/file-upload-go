package error

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	Code    int    `json:"code"`
	Err     string `json:"error"`
	Message string `json:"message"`
}

func NewBadRequestError(err error) *HttpError {
	return &HttpError{
		Code:    http.StatusBadRequest,
		Err:     "BAD_REQUEST",
		Message: fmt.Sprintf("Bad request: %s", err.Error()),
	}
}

func NewUnsupportedImgContentTypeError(contentType string) *HttpError {
	return &HttpError{
		Code:    http.StatusUnsupportedMediaType,
		Err:     "UNSUPPORTED_MEDIA_TYPE",
		Message: fmt.Sprintf("Unsupported media type: content-type %q", contentType),
	}
}

func NewInternalError(err error) *HttpError {
	return &HttpError{
		Code:    http.StatusInternalServerError,
		Err:     "INTERNAL_SERVER_ERROR",
		Message: fmt.Sprintf("Internal server error: %s", err.Error()),
	}
}

func (e *HttpError) Error() string {
	return e.Message
}
