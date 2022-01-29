package course

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Serj1c/datalearn/api/pkg/entity"
	"github.com/Serj1c/datalearn/api/pkg/errs"
	"github.com/Serj1c/datalearn/api/pkg/middleware"
	"github.com/Serj1c/datalearn/api/pkg/service"
)

type CourseHandler struct {
	l *log.Logger
	v *middleware.Validation
	p *service.CourseProcessor
}

func NewCourseHandler(l *log.Logger, v *middleware.Validation, p *service.CourseProcessor) *CourseHandler {
	return &CourseHandler{
		l: l,
		v: v,
		p: p,
	}
}

type courseData struct {
	Course entity.Course `json:"course"`
}

func (c *CourseHandler) List(rw http.ResponseWriter, r *http.Request) {
	courses, err := c.p.List()
	if err != nil {
		http.Error(rw, "Internal error", http.StatusInternalServerError)
	}
	response, err := json.Marshal(courses)
	if err != nil {
		http.Error(rw, "Error marshaling response", http.StatusInternalServerError)
	}
	rw.Write(response)
}

func (c *CourseHandler) Get(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courseID := vars["id"] /* TODO: check that ID of a correct type */
	course, err := c.p.Get(courseID)
	switch err {
	case nil:
	case errs.BadRequest:
		http.Error(rw, "Wrong data provided", http.StatusBadRequest)
	case errs.NotFound:
		http.Error(rw, "There is no such course", http.StatusBadRequest)
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

	err = c.p.Create(data.Course)
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
