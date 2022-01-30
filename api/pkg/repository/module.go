package repository

import (
	"database/sql"

	"github.com/Serj1c/datalearn/api/pkg/entity"
	"github.com/Serj1c/datalearn/api/pkg/errs"
	"github.com/Serj1c/datalearn/api/pkg/util"
)

type ModuleRepository struct {
	db *sql.DB
}

func NewModuleRepository(db *sql.DB) *ModuleRepository {
	return &ModuleRepository{
		db: db,
	}
}

func (mr *ModuleRepository) List(courseId string) ([]*entity.Module, error) {
	modules := make([]*entity.Module, 0, 10)
	rows, err := mr.db.Query("SELECT id, course_id, title, body, badge FROM modules WHERE course_id=$1", courseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		module := &entity.Module{}
		err := rows.Scan(&module.Id, &module.CourseId, &module.Title, &module.Body, &module.Badge)
		if err != nil {
			return nil, err
		}
		modules = append(modules, module)
	}
	return modules, nil
}

func (mr *ModuleRepository) Get(id string) (*entity.Module, error) {
	module := &entity.Module{}
	row := mr.db.QueryRow("SELECT id, course_id, title, body, badge FROM modules WHERE id=$1", id)
	err := row.Scan(&module.Id, &module.CourseId, &module.Title, &module.Body, &module.Badge)
	if err == sql.ErrNoRows {
		return nil, errs.NotFound
	}
	return module, nil
}

/* Administration part of repository functions */

func (mr *ModuleRepository) Create(m entity.Module) error {
	id := util.RandString()
	row, err := mr.db.Exec("INSERT into modules(id, course_id, title, body, badge) VALUES($1, $2, $3, $4, $5)",
		id, m.CourseId, m.Title, m.Body, m.Badge)
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

func (mr *ModuleRepository) Update(m entity.Module) error {
	return nil
}

func (mr *ModuleRepository) Delete(id string) error {
	return nil
}
