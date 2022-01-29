package course

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Serj1c/datalearn/api/pkg/entity"
	"github.com/Serj1c/datalearn/api/pkg/errs"
	"github.com/Serj1c/datalearn/api/pkg/middleware"
	"github.com/Serj1c/datalearn/api/pkg/operation"
	"github.com/gorilla/mux"
)

type KeyCourse struct{}

type CourseHandler struct {
	l *log.Logger
	v *middleware.Validation
	r operation.CourseRepository
}

func NewCourseHandler(l *log.Logger, v *middleware.Validation, r operation.CourseRepository) *CourseHandler {
	return &CourseHandler{
		l: l,
		v: v,
		r: r,
	}
}

type courseData struct {
	Course entity.Course `json:"course"`
}

func (c *CourseHandler) List(rw http.ResponseWriter, r *http.Request) {
	listOfCourses, err := c.r.List()
	if err != nil {
		http.Error(rw, "Internal error", http.StatusInternalServerError)
	}
	response, err := json.Marshal(listOfCourses)
	if err != nil {
		http.Error(rw, "Error marshaling response", http.StatusInternalServerError)
	}
	rw.Write(response)
}

func (c *CourseHandler) Get(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courseID := vars["id"] /* TODO: check that ID of a correct type */
	course, err := c.r.Get(courseID)
	switch err {
	case nil:
	case errs.BadRequest:
		http.Error(rw, "Wrong data provided", http.StatusBadRequest)
	case errs.NotFound:
		http.Error(rw, "There is no such a course", http.StatusBadRequest)
	}
	response, err := json.Marshal(course)
	if err != nil {
		http.Error(rw, "Error marshaling response", http.StatusInternalServerError)
	}
	rw.Write(response)
}

/* Administration endpoints handlers */

func (c *CourseHandler) Create(rw http.ResponseWriter, r *http.Request) {
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

	err = c.r.Create(data.Course)
	switch {
	case err == nil:
		rw.WriteHeader(http.StatusCreated)
	case err == errs.BadRequest:
		http.Error(rw, "Wrong data provided", http.StatusBadRequest)
	case err == errs.AlreadyExists:
		http.Error(rw, "Such course already exists", http.StatusConflict)
	default:
		http.Error(rw, "Internal error", http.StatusInternalServerError)
	}
}

func (c *CourseHandler) Delete(rw http.ResponseWriter, r *http.Request) {}

func (c *CourseHandler) Update(rw http.ResponseWriter, r *http.Request) {}
