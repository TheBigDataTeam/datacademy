package handlers

import (
	"log"
	"net/http"

	"github.com/Serj1c/datalearn/api/pkg/courses"
	"github.com/Serj1c/datalearn/api/pkg/middleware"
	"github.com/Serj1c/datalearn/api/pkg/util"
)

// KeyCourse is a key used for the Course object in the context
type KeyCourse struct{}

// Courses is a handler for getting and updating courses
type Courses struct {
	l *log.Logger
	v *middleware.Validation
	r *courses.Repo
}

// NewCourses creates a courses handler with the given logger
func NewCourses(l *log.Logger, v *middleware.Validation, r *courses.Repo) *Courses {
	return &Courses{l, v, r}
}

// ListAll handles GET requests and returns all current courses
func (c *Courses) ListAll(rw http.ResponseWriter, r *http.Request) {
	c.l.Println("[SUCCESS] get all records")
	courses, err := c.r.GetCourses()
	if err != nil {
		c.l.Println("[ERROR] receiving courses from db: ", err)
	}

	err = util.ToJSON(courses, rw)
	if err != nil {
		c.l.Println("[ERROR] serializing courses ", err)
	}
}

// ListOne handles GET requests for a single course
func (c *Courses) ListOne(rw http.ResponseWriter, r *http.Request) {
	id := util.GetIDfromRequest(r)
	c.l.Println("[SUCCESS] get record id", id)

	course, err := c.r.GetCourseByID(id)
	c.l.Println(course)
	if err == courses.ErrorCourseNotFound {
		c.l.Println("[ERROR]", err)
		rw.WriteHeader(http.StatusNotFound)
		util.ToJSON(&util.GenericError{Message: err.Error()}, rw)
		return
	}
	if err != nil {
		c.l.Println("[ERROR] fetching course", err)
		rw.WriteHeader(http.StatusInternalServerError)
		util.ToJSON(&util.GenericError{Message: err.Error()}, rw)
		return
	}
	err = util.ToJSON(course, rw)
	if err != nil {
		c.l.Println("[ERROR] serializing course", err)
	}
}

// Delete handles DELETE request and removes items from the DB
func (c *Courses) Delete(rw http.ResponseWriter, r *http.Request) {
	id := util.GetIDfromRequest(r)

	c.l.Println("[DEBUG] deleting record id", id)

	err := c.r.DeleteCourse(id)
	if err == courses.ErrorCourseNotFound {
		c.l.Println("[ERROR] deleting record id does not exist")
		rw.WriteHeader(http.StatusNotFound)
		util.ToJSON(&util.GenericError{Message: err.Error()}, rw)
		return
	}
	if err != nil {
		c.l.Println("[ERROR] deleting record", err)
		rw.WriteHeader(http.StatusInternalServerError)
		util.ToJSON(&util.GenericError{Message: err.Error()}, rw)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}

// Create handles POST requests to add new courses
func (c *Courses) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the course from the context
	course := r.Context().Value(KeyCourse{}).(courses.Course)

	c.l.Printf("[SUCCESS] Inserting course: %#v\n", course)
	c.r.AddCourse(course)
}

// Update handles PUT requests to update courses
func (c *Courses) Update(rw http.ResponseWriter, r *http.Request) {
	// fetch the course from the context
	course := r.Context().Value(KeyCourse{}).(courses.Course)
	c.l.Println("[SUCCESS] updating record id", course.ID)

	err := c.r.UpdateCourse(course)
	if err == courses.ErrorCourseNotFound {
		c.l.Println("[ERROR] course not found", err)
		rw.WriteHeader(http.StatusNotFound)
		util.ToJSON(&util.GenericError{Message: "Product not found in database"}, rw)
		return
	}
	// write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}
