package repository

import (
	"database/sql"

	"github.com/Serj1c/datalearn/api/pkg/entity"
	"github.com/Serj1c/datalearn/api/pkg/errs"
	"github.com/Serj1c/datalearn/api/pkg/util"
)

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

func (r *CourseRepository) List() ([]*entity.Course, error) {
	courses := make([]*entity.Course, 0, 10)
	rows, err := r.db.Query("SELECT id, title, author, description, techstack, moduleQuantity, workshopQuantity FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		course := &entity.Course{}
		err := rows.Scan(&course.ID, &course.Title, &course.Author, &course.Description, &course.TechStack, &course.ModuleQuantity, &course.WorkshopQuantity)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (r *CourseRepository) Get(id string) (*entity.Course, error) {
	course := &entity.Course{}
	row := r.db.QueryRow("SELECT id, title, author, description, techstack, moduleQuantity, workshopQuantity FROM courses WHERE id=$1", id)
	err := row.Scan(&course.ID, &course.Title, &course.Author, &course.Description, &course.TechStack, &course.ModuleQuantity, &course.WorkshopQuantity)
	if err == sql.ErrNoRows {
		return nil, errs.NotFound
	}
	return course, nil
}

/* Administration part of repository functions */

func (r *CourseRepository) Create(c entity.Course) error {
	id := util.RandString()
	row, err := r.db.Exec("INSERT into courses(id, title, author, description, techstack, moduleQuantity, workshopQuantity) VALUES($1, $2, $3, $4, $5, $6, $7)",
		id, c.Title, c.Author, c.Description, c.TechStack, c.ModuleQuantity, c.WorkshopQuantity)
	if err != nil {
		return errs.BadRequest
	}
	affected, err := row.RowsAffected()
	if affected == 0 {
		return errs.AlreadyExists
	}
	if err != nil {
		return err
	}
	return nil
}

func (r *CourseRepository) Update(c entity.Course) error {
	return nil
}

func (r *CourseRepository) Delete(id string) error {
	return nil
}
