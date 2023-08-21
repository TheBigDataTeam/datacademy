package middleware

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

// ValidationError wraps the validator FieldError so we do not expose
// this to out code
type ValidationError struct {
	validator.FieldError
}

// Error ...
func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

// Validation ...
type Validation struct {
	validate *validator.Validate
}

// NewValidation creates a new Validation type
func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("user", validateUser)
	return &Validation{validate}
}

// Validate ...
func (v *Validation) Validate(i interface{}) ValidationErrors {
	errs := v.validate.Struct(i).(validator.ValidationErrors)

	if len(errs) == 0 {
		return nil
	}

	var returnErrs []ValidationError
	for _, err := range errs {
		// case the FieldError into our ValidationError and append to the slice
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}
	return returnErrs
}

func validateUser(fl validator.FieldLevel) bool {
	// author is expected to be in the following format: Abcccc Defffg
	regxp := regexp.MustCompile(`[A-Z][a-z]+ [A-Z][a-z]+`) // regexp is not valid TODO!
	allMatches := regxp.FindAllString(fl.Field().String(), -1)

	if len(allMatches) > 1 || len(allMatches) == 0 {
		return false
	}
	return true
}
