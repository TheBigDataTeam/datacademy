package handlers

import (
	"context"
	"net/http"

	"github.com/Serj1c/datalearn/api/data"
)

// MiddlewareValidateCourse validates the course in the request and calls next if ok
func (c *Courses) MiddlewareValidateCourse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		course := &data.Course{}

		err := data.FromJSON(course, r.Body)
		if err != nil {
			c.l.Println("[ERROR] deserializing course", err)
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		// validate the course
		errs := c.v.Validate(course)
		if len(errs) != 0 {
			c.l.Println("[ERROR] validating course", err)

			// return the validation messages
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}
		// add the course to the context
		ctx := context.WithValue(r.Context(), KeyCourse{}, course)
		r = r.WithContext(ctx)

		// call the next handler, which can be another middleware in the chain, or the final one
		next.ServeHTTP(rw, r)
	})
}
