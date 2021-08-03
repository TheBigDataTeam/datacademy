package handlers

import (
	"net/http"

	"github.com/Serj1c/datalearn/api/pkg/data"
)

// CreateAuthor handles POST requests to add new authors
func (a *Authors) CreateAuthor(rw http.ResponseWriter, r *http.Request) {
	// fetch the author from the context
	author := r.Context().Value(KeyAuthor{}).(data.Author)

	a.l.Printf("[SUCCESS] Inserting course: %#v\n", author)

	data.AddAuthor(author)
}
