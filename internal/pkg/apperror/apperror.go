package apperror

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

type Error struct {
	statusCode int
	stack      []byte
	err        error
	message    string
}

func (e *Error) Error() string {
	return fmt.Sprintf("error %d: %s (%q)", e.statusCode, e.message, e.err)
}

func (e *Error) GetError() error {
	return e.err
}

func (e *Error) GetStatusCode() int {
	return e.statusCode
}

func (e *Error) GetErrorMessage() string {
	return e.message
}

func (e *Error) GetStackTrace() string {
	return e.message
}

func StatusBadRequest(err error, message string) *Error {
	return &Error{
		statusCode: http.StatusBadRequest,
		stack:      debug.Stack(),
		err:        err,
		message:    message,
	}
}

func StatusInternalServerError(err error, message string) *Error {
	return &Error{
		statusCode: http.StatusInternalServerError,
		stack:      debug.Stack(),
		err:        err,
		message:    message,
	}
}

func StatusUnauthorized(err error, message string) *Error {
	return &Error{
		statusCode: http.StatusUnauthorized,
		stack:      debug.Stack(),
		err:        err,
		message:    message,
	}
}
