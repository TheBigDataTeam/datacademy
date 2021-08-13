package users

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Serj1c/datalearn/api/pkg/util"
)

// Repo represents a data base
type Repo struct {
	db *sql.DB
}

// NewRepo returns an instance of a Repo
func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

// ErrorUserAlreadyExists is returned when one tries to add user which already exists
var ErrorUserAlreadyExists = errors.New("User already exists")

// Create creates a new user and adds her/him to the a data base
func (r *Repo) Create(email string, name string, surname string, password string) (string, error) {
	id := util.RandString()
	row, err := r.db.Exec("INSERT into users(id, email, name, surname, password) VALUES($1, $2, $3, $4, $5)", id, email, name, surname, password)
	if err != nil {
		return "", fmt.Errorf("Error while inserting of user in db: %s", err)
	}
	affected, err := row.RowsAffected()
	if affected == 0 {
		return "", ErrorUserAlreadyExists
	}
	if err != nil {
		return "", err
	}
	return id, nil
}
