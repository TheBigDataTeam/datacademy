package util

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

/* TODO: make this function available for all handlers */

// getNextId generates a new id for a product being inserted into db
func getNextID() /* int */ {
	//lastItemInDB := courseList[len(courseList)-1]
	return //lastItemInDB.ID + 1
}

// GetIDfromRequest returns the ID from the URL
func GetIDfromRequest(r *http.Request) string {
	// parse the id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id := vars["id"]
	return id
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal
// as it does not have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
func ToJSON(i interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(i)
}

// FromJSON deserializes the object from JSON string
// in an io.Reader to the given interface
func FromJSON(i interface{}, r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(i)
}
