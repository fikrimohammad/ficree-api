package apierror

import (
	"errors"
	"net/http"
)

// APIError .....
type APIError struct {
	Code    int
	Message string
	GoErr   error
}

// Error ......
func (e *APIError) Error() string {
	return e.Message
}

// New ......
func New(code int, message string) error {
	return &APIError{
		Code:    code,
		Message: message,
		GoErr:   errors.New(message),
	}
}

// FromError ...
func FromError(code int, err error) error {
	return &APIError{
		Code:    code,
		Message: err.Error(),
		GoErr:   err,
	}
}

// GetHTTPStatus ...
func GetHTTPStatus(err error) int {
	e, ok := err.(*APIError)
	if ok {
		return e.Code
	}
	return http.StatusInternalServerError
}
