package handlers

import (
	"net/http"

	"github.com/Serj1c/datalearn/api/data"
	"github.com/Serj1c/datalearn/api/util"
)

// UpdAuthor handles PUT requests to update authors
func (a *Authors) UpdAuthor(rw http.ResponseWriter, r *http.Request) {
	// fetch the author from the context
	author := r.Context().Value(KeyAuthor{}).(data.Author)
	a.l.Println("[SUCCESS] updating record id", author.ID)
	err := data.UpdateAuthor(author)
	if err == data.ErrorAuthorNotFound {
		a.l.Println("[ERROR] author not found", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&util.GenericError{Message: "Product not found in database"}, rw)
	}
	// write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}
