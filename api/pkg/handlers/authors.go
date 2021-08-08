package handlers

import (
	"log"
	"net/http"

	"github.com/Serj1c/datalearn/api/pkg/authors"
	"github.com/Serj1c/datalearn/api/pkg/middleware"
	"github.com/Serj1c/datalearn/api/pkg/util"
)

// KeyAuthor is a key used for the Author object in the context
type KeyAuthor struct{}

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

// ListAll handles GET requests and returns all current authors
func (a *Authors) ListAll(rw http.ResponseWriter, r *http.Request) {
	a.l.Println("[SUCCESS] get all authors")
	authors, err := a.r.GetAuthors()
	if err != nil {
		a.l.Println("[ERROR] receiving authors from db: ", err)
	}
	err = util.ToJSON(authors, rw)
	if err != nil {
		a.l.Println("[ERROR] serializing author", err)
	}
}

// ListOne handles GET requests for a single author
func (a *Authors) ListOne(rw http.ResponseWriter, r *http.Request) {
	id := util.GetIDfromRequest(r)

	a.l.Println("[SUCCESS] get record id", id)

	author, err := a.r.GetAuthorByID(id)

	switch err {
	case nil:

	case authors.ErrorAuthorNotFound:
		a.l.Println("[ERROR] fetching author - author not found", err)
		rw.WriteHeader(http.StatusNotFound)
		util.ToJSON(&util.GenericError{Message: err.Error()}, rw)
		return

	default:
		a.l.Println("[ERROR] fetching author", err)
		rw.WriteHeader(http.StatusInternalServerError)
		util.ToJSON(&util.GenericError{Message: err.Error()}, rw)
		return
	}
	err = util.ToJSON(author, rw)
	if err != nil {
		a.l.Println("[ERROR] serializing author", err)
	}
}

// Create handles POST requests to add new authors
func (a *Authors) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the author from the context
	author := r.Context().Value(KeyAuthor{}).(authors.Author)

	a.l.Printf("[SUCCESS] Inserting course: %#v\n", author)

	a.r.AddAuthor(author)
}

// Update handles PUT requests to update authors
func (a *Authors) Update(rw http.ResponseWriter, r *http.Request) {
	// fetch the author from the context
	author := r.Context().Value(KeyAuthor{}).(authors.Author)
	a.l.Println("[SUCCESS] updating record id", author.ID)
	err := a.r.UpdateAuthor(author)
	if err == authors.ErrorAuthorNotFound {
		a.l.Println("[ERROR] author not found", err)
		rw.WriteHeader(http.StatusNotFound)
		util.ToJSON(&util.GenericError{Message: "Product not found in database"}, rw)
	}
	// write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}

// Delete handles requests for removing authors from db
func (a *Authors) Delete(rw http.ResponseWriter, r *http.Request) {
	id := util.GetIDfromRequest(r)

	a.l.Println("[DEBUG] deleting record id ", id)

	err := a.r.DeleteAuthor(id)
	if err == authors.ErrorAuthorNotFound {
		a.l.Println("[ERROR] deleting record id does not exist")
		rw.WriteHeader(http.StatusNotFound)
		util.ToJSON(&util.GenericError{Message: err.Error()}, rw)
		return
	}
	if err != nil {
		a.l.Println("[ERROR] deleting record id")
		rw.WriteHeader(http.StatusInternalServerError)
		util.ToJSON(&util.GenericError{Message: err.Error()}, rw)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
