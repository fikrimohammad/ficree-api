package apierror

import "errors"

// Error .....
type Error struct {
	Code    int
	Message string
	GoErr   error
}

// Error ......
func (e *Error) Error() string {
	return e.Message
}

// New ......
func New(code int, message string) error {
	return &Error{
		Code:    code,
		Message: message,
		GoErr:   errors.New(message),
	}
}

// FromError ...
func FromError(code int, err error) error {
	return &Error{
		Code:    code,
		Message: err.Error(),
		GoErr:   err,
	}
}
