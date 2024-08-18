package errorutil

import (
	"errors"
	"fmt"
	"net/http"
)

type Error struct {
	Message string
	Cause   error
}

func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func New(message string) error {
	return &Error{Message: message}
}

func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return &Error{Message: message, Cause: err}
}

var (
	ErrorNotFound    = New("data is not found")
	ErrUnauthorized  = New("unauthorized")
	ErrForbidden     = New("forbidden")
	ErrBadRequest    = New("bad request")
	ErrInternalError = New("internal server error")
	ErrNonFailure    = New("non failure")
)

func HandleError(err error) (int, string) {
	var status int
	var errorMessage string

	switch {
	case errors.Is(err, ErrorNotFound):
		status = http.StatusNotFound
		errorMessage = err.Error()
	case errors.Is(err, ErrUnauthorized):
		status = http.StatusUnauthorized
		errorMessage = err.Error()
	case errors.Is(err, ErrForbidden):
		status = http.StatusForbidden
		errorMessage = err.Error()
	case errors.Is(err, ErrBadRequest):
		status = http.StatusBadRequest
		errorMessage = err.Error()
	case errors.Is(err, ErrInternalError):
		status = http.StatusInternalServerError
		errorMessage = err.Error()
	default:
		status = http.StatusInternalServerError
		errorMessage = "Unknown error"
	}

	return status, errorMessage
}
