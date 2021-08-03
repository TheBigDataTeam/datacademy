package util

import "fmt"

// ErrInvalidCoursePath is an error message when the course path is not valid
var ErrInvalidCoursePath = fmt.Errorf("Invalid Path, path should be /courses/[id]")

// ErrInvalidAuthorPath is an error message when the author path is not valid
var ErrInvalidAuthorPath = fmt.Errorf("Invalid Path, path should be /authors/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}
