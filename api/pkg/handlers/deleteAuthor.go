package handlers

import (
	"net/http"

	"github.com/Serj1c/datalearn/api/pkg/data"
	"github.com/Serj1c/datalearn/api/pkg/util"
)

// DelAuthor ...
func (a Authors) DelAuthor(rw http.ResponseWriter, r *http.Request) {
	id := util.GetIDfromRequest(r)

	a.l.Println("[DEBUG] deleting record id ", id)

	err := data.DeleteAuthor(id)
	if err == data.ErrorAuthorNotFound {
		a.l.Println("[ERROR] deleting record id does not exist")
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&util.GenericError{Message: err.Error()}, rw)
		return
	}
	if err != nil {
		a.l.Println("[ERROR] deleting record id")
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&util.GenericError{Message: err.Error()}, rw)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
