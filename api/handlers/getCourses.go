package handlers

import (
	"net/http"

	"github.com/Serj1c/datalearn/api/data"
	"github.com/Serj1c/datalearn/api/util"
)

// ListAllCourses handles GET requests and returns all current courses
func (c *Courses) ListAllCourses(rw http.ResponseWriter, r *http.Request) {
	c.l.Println("[SUCCESS] get all records")

	courses := data.GetCourses()

	err := data.ToJSON(courses, rw)
	if err != nil {
		c.l.Println("[ERROR] serializing product", err)
	}
}

// ListSingleCourse handles GET requests for a single course
func (c *Courses) ListSingleCourse(rw http.ResponseWriter, r *http.Request) {
	id := util.GetCourseIDfromRequest(r)

	c.l.Println("[SUCCESS] get record id", id)

	course, err := data.GetCourseByID(id)

	switch err {
	case nil:

	case data.ErrorCourseNotFound:
		c.l.Println("[ERROR] fetching course", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return

	default:
		c.l.Println("[ERROR] fetching course", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	err = data.ToJSON(course, rw)
	if err != nil {
		c.l.Println("[ERROR] serializing course", err)
	}
}
