package handlers

import (
	"net/http"

	"github.com/Serj1c/datalearn/api/data"
	"github.com/Serj1c/datalearn/api/util"
)

// UpdCourse handles PUT requests to update courses
func (c *Courses) UpdCourse(rw http.ResponseWriter, r *http.Request) {
	// fetch the course from the context
	course := r.Context().Value(KeyCourse{}).(data.Course)
	c.l.Println("[SUCCESS] updating record id", course.ID)

	err := data.UpdateCourse(course)
	if err != data.ErrorCourseNotFound {
		c.l.Println("[ERROR] course not found", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&util.GenericError{Message: "Product not found in database"}, rw)
		return
	}
	// write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}
