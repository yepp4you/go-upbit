package common

import (
	"fmt"
)

// APIError define API error when response status is 4xx or 5xx
type APIError struct {
	Err ErrorDetail `json:"error"`
}
type ErrorDetail struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

// Error return error code and message
func (e APIError) Error() string {
	return fmt.Sprintf("<APIError> Message=%s, Name=%s", e.Err.Message, e.Err.Name)
}

// IsAPIError check if e is an API error
func IsAPIError(e error) bool {
	_, ok := e.(*APIError)
	return ok
}
