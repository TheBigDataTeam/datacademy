package handlers

import (
	"log"

	"github.com/Serj1c/datalearn/api/data"
)

// KeyAuthor is a key used for the Author object in the context
type KeyAuthor struct{}

// Authors is a handler for getting and updating authors
type Authors struct {
	l *log.Logger
	v *data.Validation
}

// NewAuthors creates an authors handler with the given logger
func NewAuthors(l *log.Logger, v *data.Validation) *Authors {
	return &Authors{l, v}
}
