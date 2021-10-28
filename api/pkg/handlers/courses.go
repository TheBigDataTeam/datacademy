package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Serj1c/datalearn/api/pkg/courses"
	"github.com/Serj1c/datalearn/api/pkg/middleware"
)

// KeyCourse is a key used for the Course object in the context
type KeyCourse struct{}

// Courses is a handler for getting and updating courses
type Courses struct {
	l *log.Logger
	v *middleware.Validation
	r *courses.Repo
}

// NewCourses creates a courses handler
func NewCourses(l *log.Logger, v *middleware.Validation, r *courses.Repo) *Courses {
	return &Courses{l, v, r}
}

type courseData struct {
	Course courses.Course `json:"course"`
}

// Create handles POST requests to add new courses
func (c *Courses) Create(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Cannot read body", http.StatusInternalServerError)
	}
	r.Body.Close()

	data := &courseData{}
	err = json.Unmarshal(body, data)
	if err != nil {
		http.Error(rw, "Error unmarshaling request body", http.StatusInternalServerError)
	}

	err = c.r.AddCourse(data.Course)
	switch {
	case err == nil:
		rw.WriteHeader(http.StatusCreated)
	case err == courses.ErrBadRequest:
		http.Error(rw, "Wrong data provided", http.StatusBadRequest)
	case err == courses.ErrAlreadyExists:
		http.Error(rw, "Such course already exists", http.StatusConflict)
	default:
		http.Error(rw, "Internal error", http.StatusInternalServerError)
	}
}

// List handles GET requests and returns all current courses
func (c *Courses) List(rw http.ResponseWriter, r *http.Request) {

}

// Get handles GET requests for a single course
func (c *Courses) Get(rw http.ResponseWriter, r *http.Request) {

}

// Delete handles DELETE request and removes items from the DB
func (c *Courses) Delete(rw http.ResponseWriter, r *http.Request) {

}

// Update handles PUT requests to update courses
func (c *Courses) Update(rw http.ResponseWriter, r *http.Request) {

}
