package errs

import "errors"

var (
	NotFound      = errors.New("Not found")
	AlreadyExists = errors.New("Already exists")
	BadRequest    = errors.New("Wrong data provided")
)
