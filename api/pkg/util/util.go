package util

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/* TODO: make this function available for all handlers */

// getNextId generates a new id for a product being inserted into db
func getNextID() /* int */ {
	//lastItemInDB := courseList[len(courseList)-1]
	return //lastItemInDB.ID + 1
}

// GetIDfromRequest returns the ID from the URL
func GetIDfromRequest(r *http.Request) int {
	// parse the id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err) /* TODO must be no panic! */
	}
	return id
}
