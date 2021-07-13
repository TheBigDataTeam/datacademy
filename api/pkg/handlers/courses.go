package handlers

import (
	"log"

	"github.com/Serj1c/datalearn/api/pkg/data"
)

// KeyCourse is a key used for the Course object in the context
type KeyCourse struct{}

// Courses is a handler for getting and updating courses
type Courses struct {
	l *log.Logger
	v *data.Validation
}

// NewCourses creates a courses handler with the given logger
func NewCourses(l *log.Logger, v *data.Validation) *Courses {
	return &Courses{l, v}
}
