package handlers

import (
	"net/http"

	"github.com/Serj1c/datalearn/api/pkg/data"
	"github.com/Serj1c/datalearn/api/pkg/util"
)

// DelCourse handles DELETE request and removes items from the DB
func (c Courses) DelCourse(rw http.ResponseWriter, r *http.Request) {
	id := util.GetIDfromRequest(r)

	c.l.Println("[DEBUG] deleting record id", id)

	err := data.DeleteCourse(id)
	if err == data.ErrorCourseNotFound {
		c.l.Println("[ERROR] deleting record id does not exist")
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&util.GenericError{Message: err.Error()}, rw)
		return
	}
	if err != nil {
		c.l.Println("[ERROR] deleting record", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&util.GenericError{Message: err.Error()}, rw)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
