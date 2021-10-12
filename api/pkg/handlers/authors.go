package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Serj1c/datalearn/api/pkg/authors"
	"github.com/Serj1c/datalearn/api/pkg/middleware"
)

// Authors is a handler for getting and updating authors
type Authors struct {
	l *log.Logger
	v *middleware.Validation
	r *authors.Repo
}

// NewAuthors creates an authors handler with the given logger
func NewAuthors(l *log.Logger, v *middleware.Validation, r *authors.Repo) *Authors {
	return &Authors{l, v, r}
}

// Create handles POST requests to add new authors
func (a *Authors) Create(rw http.ResponseWriter, r *http.Request) {
	a.l.Println("[ATTENTION] Create a new author")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Cannot read body", http.StatusInternalServerError)
	}
	r.Body.Close()

	author := &authors.Author{}
	err = json.Unmarshal(body, author)
	if err != nil {
		http.Error(rw, "Error unmarshaling request body", http.StatusInternalServerError)
	}
	err = a.r.AddAuthor(*author)
	switch {
	case err == nil:
		rw.WriteHeader(http.StatusCreated)
	case err == authors.ErrorBadRequest:
		http.Error(rw, "Wrong data provided", http.StatusBadRequest)
	case err == authors.ErrorAuthorAlreadyExists:
		http.Error(rw, "Such author already exists", http.StatusForbidden)
	default:
		http.Error(rw, "Internal error", http.StatusInternalServerError)
	}
}

// ListAll handles GET requests and returns all current authors
func (a *Authors) ListAll(rw http.ResponseWriter, r *http.Request) {
	listOfAuthors, err := a.r.GetAuthors()
	switch {
	case err == nil:
	case err == authors.ErrorBadRequest:
		http.Error(rw, "", http.StatusBadRequest)
	}
	response, err := json.Marshal(listOfAuthors)
	if err != nil {
		http.Error(rw, "Internal error", http.StatusInternalServerError)
	}
	rw.Write(response)
}

// ListOne handles GET requests for a single author
func (a *Authors) ListOne(rw http.ResponseWriter, r *http.Request) {

}

// Update handles PUT requests to update authors
func (a *Authors) Update(rw http.ResponseWriter, r *http.Request) {

}

// Delete handles requests for removing authors from db
func (a *Authors) Delete(rw http.ResponseWriter, r *http.Request) {

}
