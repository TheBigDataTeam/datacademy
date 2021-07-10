package handlers

import (
	"net/http"

	"github.com/Serj1c/datalearn/api/data"
)

// CreateCourse handles POST requests to add new courses
func (c *Courses) CreateCourse(rw http.ResponseWriter, r *http.Request) {
	// fetch the course from the context
	course := r.Context().Value(KeyCourse{}).(data.Course)

	c.l.Printf("[SUCCESS] Inserting course: %#v\n", course)
	data.AddCourse(course)
}
