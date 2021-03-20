package handlers

import (
	"fmt"
	"log"

	"github.com/Serj1c/datalearn/api/data"
)

// KeyCourse is a key used for the Course obkect in the context
type KeyCourse struct{}

// Courses is a handler for getting and updating products
type Courses struct {
	l *log.Logger
	v *data.Validation
}

// NewCourses creates a courses handler with the given logger
func NewCourses(l *log.Logger, v *data.Validation) *Courses {
	return &Courses{l, v}
}

// ErrInvalidCoursePath is an error message when the course path is not valid
var ErrInvalidCoursePath = fmt.Errorf("Invalid Path, path should be /courses/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}
