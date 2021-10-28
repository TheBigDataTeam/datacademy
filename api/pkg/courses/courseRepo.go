package courses

import (
	"database/sql"
	"errors"
)

// Repo represents DB
type Repo struct {
	db *sql.DB
}

// NewRepo returns an instance of a Repo
func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

var (
	// ErrNoRecord is returned when no record in database is found
	ErrNoRecord = errors.New("No course record found")
	// ErrAlreadyExists is returned when one tries to add author which already exists
	ErrAlreadyExists = errors.New("Course already exists")
	// ErrBadRequest is returned when wrong data provided
	ErrBadRequest = errors.New("Wrong data provided")
)

// GetCourses returns a list of courses
func (r *Repo) GetCourses() ([]*Course, error) {
	return nil, nil
}

// GetCourseByID returns a single course which matches the id from the DB
// if a course is not found this function returns a CourseNotFound error
func (r *Repo) GetCourseByID(id string) (*Course, error) {
	return nil, nil
}

// UpdateCourse replaces a course in the DB with the given item.
func (r *Repo) UpdateCourse(c Course) error {
	return nil
}

// AddCourse adds new course to the DB
func (r *Repo) AddCourse(c Course) error {
	return nil
}

// DeleteCourse deletes a course from the DB
func (r *Repo) DeleteCourse(id string) error {
	return nil
}
