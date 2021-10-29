package courses

import (
	"database/sql"
	"errors"

	"github.com/Serj1c/datalearn/api/pkg/util"
)

// Repo represents DB
type Repo struct {
	db *sql.DB
}

// New returns an instance of a Repo
func New(db *sql.DB) *Repo {
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

// GetCourses returns a list of courses.
func (r *Repo) GetCourses() ([]*Course, error) {
	coursesFromDB := make([]*Course, 0, 10)
	rows, err := r.db.Query("SELECT title, author, description, techstack, moduleQuantity, workshopQuantity FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		course := &Course{}
		err := rows.Scan(&course.Title, &course.Author, &course.Description, &course.TechStack, &course.ModuleQuantity, &course.WorkshopQuantity)
		if err != nil {
			return nil, err
		}
		coursesFromDB = append(coursesFromDB, course)
	}
	return coursesFromDB, nil
}

// GetCourseByID returns a single course which matches the id from the DB.
func (r *Repo) GetCourseByID(id string) (*Course, error) {
	return nil, nil
}

/* Administration part of repository functions */

// AddCourse adds a new course to the DB.
func (r *Repo) AddCourse(c Course) error {
	id := util.RandString()
	row, err := r.db.Exec("INSERT into courses(id, title, author, description, techstack, moduleQuantity, workshopQuantity) VALUES($1, $2, $3, $4, $5, $6, $7)",
		id, c.Title, c.Author, c.Description, c.TechStack, c.ModuleQuantity, c.WorkshopQuantity)
	if err != nil {
		return ErrBadRequest
	}
	affected, err := row.RowsAffected()
	if affected == 0 {
		return ErrAlreadyExists
	}
	if err != nil {
		return err
	}
	return nil
}

// UpdateCourse replaces a course in the DB with the given item.
func (r *Repo) UpdateCourse(c Course) error {
	return nil
}

// DeleteCourse deletes a course from the DB
func (r *Repo) DeleteCourse(id string) error {
	return nil
}
