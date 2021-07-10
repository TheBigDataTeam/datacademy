package handlers

import (
	"net/http"

	"github.com/Serj1c/datalearn/api/data"
	"github.com/Serj1c/datalearn/api/util"
)

// ListAllAuthors handles GET requests and returns all current authors
func (a *Authors) ListAllAuthors(rw http.ResponseWriter, r *http.Request) {
	a.l.Println("[SUCCESS] get all records")

	authors := data.GetAuthors()

	err := data.ToJSON(authors, rw)
	if err != nil {
		a.l.Println("[ERROR] serializing author", err)
	}
}

// ListSingleAuthor handles GET requests for a single author
func (a *Authors) ListSingleAuthor(rw http.ResponseWriter, r *http.Request) {
	id := util.GetIDfromRequest(r)

	a.l.Println("[SUCCESS] get record id", id)

	author, err := data.GetAuthorByID(id)

	switch err {
	case nil:

	case data.ErrorAuthorNotFound:
		a.l.Println("[ERROR] fetching author - author not found", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&util.GenericError{Message: err.Error()}, rw)
		return

	default:
		a.l.Println("[ERROR] fetching author", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&util.GenericError{Message: err.Error()}, rw)
		return
	}
	err = data.ToJSON(author, rw)
	if err != nil {
		a.l.Println("[ERROR] serializing author", err)
	}
}
