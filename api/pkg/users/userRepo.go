package users

import (
	"database/sql"
	"errors"

	"github.com/Serj1c/datalearn/api/pkg/util"
)

// Repo represents a database
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
	errNoRecord      = errors.New("No user record found")
	errWrongPassword = errors.New("Wrong password")
	errUserExists    = errors.New("User already exists")
)

// ErrorUserAlreadyExists is returned when one tries to add user which already exists
var ErrorUserAlreadyExists = errors.New("User already exists")

// ErrorBadRequest is returned when one tries to create a user with a wrong data
var ErrorBadRequest = errors.New("Error inserting info into db")

func checkPasswordIsValid(row *sql.Row, password string) (*User, error) {
	user := &User{}
	//var passwordInDB string
	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Surname, &user.Version)
	if err == sql.ErrNoRows {
		return nil, errNoRecord
	}
	return user, nil
}

// Create creates a new user and adds her/him to the database
func (r *Repo) Create(email string, name string, surname string, password string) (string, error) {
	id := util.RandString()
	row, err := r.db.Exec("INSERT into users(id, email, name, surname, password) VALUES($1, $2, $3, $4, $5)", id, email, name, surname, password)
	if err != nil {
		return "", ErrorBadRequest
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

// Authorize checks email and password and gets back userID from the database or an error
func (r *Repo) Authorize(email string, password string) (string, error) {
	row := r.db.QueryRow("SELECT id, email, name, surname, password, version FROM users WHERE email=$1", email)
	user, err := checkPasswordIsValid(row, password)
	if err == errNoRecord {
		return "", err
	}
	return user.ID, nil
}
