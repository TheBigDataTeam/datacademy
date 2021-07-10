package data

import "fmt"

// ErrorCourseNotFound is an error raised when a course can not be found
var ErrorCourseNotFound = fmt.Errorf("Course not found")

// ErrorAuthorNotFound is an error raised when a author can not be found
var ErrorAuthorNotFound = fmt.Errorf("Author not found")
