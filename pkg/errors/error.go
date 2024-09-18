package errors

import (
	"net/http"
)

type status struct {
	code        int
	description string
}

// NewError returns a new error with status code and description
func NewError(code int, description string) error {
	return status{code: code, description: description}
}

// Error implements the error interface
func (s status) Error() string {
	return s.description
}

// ExtractStatusCode extracts the status code from the error
func ExtractStatusCode(err error) (int, bool) {
	switch err := err.(type) {
	case status:
		// status error
		return err.code, true
	default:
		// non-status error
		return http.StatusInternalServerError, false
	}
}
