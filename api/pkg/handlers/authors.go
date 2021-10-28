package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Serj1c/datalearn/api/pkg/authors"
	"github.com/Serj1c/datalearn/api/pkg/middleware"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
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

type authorData struct {
	Author authors.Author `json:"author"`
}

// Create handles POST requests to add new authors
func (a *Authors) Create(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Cannot read body", http.StatusInternalServerError)
	}
	r.Body.Close()

	data := &authorData{}
	err = json.Unmarshal(body, data)
	if err != nil {
		http.Error(rw, "Error unmarshaling request body", http.StatusInternalServerError)
	}

	err = a.r.AddAuthor(data.Author)
	switch {
	case err == nil:
		rw.WriteHeader(http.StatusCreated)
	case err == authors.ErrorBadRequest:
		http.Error(rw, "Wrong data provided", http.StatusBadRequest)
	case err == authors.ErrorAuthorAlreadyExists:
		http.Error(rw, "Such author already exists", http.StatusConflict)
	default:
		http.Error(rw, "Internal error", http.StatusInternalServerError)
	}
}

// List handles GET requests for all authors
func (a *Authors) List(rw http.ResponseWriter, r *http.Request) {
	listOfAuthors, err := a.r.GetAuthors()
	switch {
	case err == nil:
	case err == authors.ErrorBadRequest:
		http.Error(rw, "Wrong data provided", http.StatusBadRequest)
	}
	response, err := json.Marshal(listOfAuthors)
	if err != nil {
		http.Error(rw, "Error marshaling response", http.StatusInternalServerError)
	}
	rw.Write(response)
}

// Get handles GET requests for a single author
func (a *Authors) Get(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["id"]) {
		http.Error(rw, "ID is of wrong format", http.StatusBadRequest)
		return
	}
	authorID := bson.ObjectIdHex(vars["id"])
	fmt.Println(authorID)
	author, err := a.r.GetAuthorByID(authorID)
	switch err {
	case nil:
	case authors.ErrorBadRequest:
		http.Error(rw, "Wrong data provided", http.StatusBadRequest)
	case authors.ErrorNoRecord:
		http.Error(rw, "", http.StatusBadRequest)
	}
	response, err := json.Marshal(author)
	if err != nil {
		http.Error(rw, "Error marshaling response", http.StatusInternalServerError)
	}
	rw.Write(response)
}

// Update handles PUT requests to update authors
func (a *Authors) Update(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Update resourse is not allowed", http.StatusForbidden)
}

// Delete handles requests for removing authors from db
func (a *Authors) Delete(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Delete resourse is not allowed", http.StatusForbidden)
}
