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
	// ErrNoRecord  is returned when no record in database is found
	ErrNoRecord = errors.New("No user record found")
	// ErrWrongPassword is returned when passwords do not match
	ErrWrongPassword = errors.New("Wrong password")
	// ErrorUserAlreadyExists is returned when one tries to add user which already exists
	ErrorUserAlreadyExists = errors.New("User already exists")
	// ErrorBadRequest is returned when one tries to create a user with a wrong data
	ErrorBadRequest = errors.New("Error inserting info into db")
)

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

// Authenticate checks email and password and gets back userID from the database or an error
func (r *Repo) Authenticate(email string, password string) (string, error) {
	row := r.db.QueryRow("SELECT id, email, name, surname, password, version FROM users WHERE email=$1", email)
	user, err := checkPasswordIsValid(row, password)
	switch err {
	case nil:
	case ErrNoRecord:
		return "", ErrNoRecord
	case ErrWrongPassword:
		return "", ErrWrongPassword
	}
	return user.ID, nil
}

func checkPasswordIsValid(row *sql.Row, password string) (*User, error) {
	user := &User{}
	var passwordStoredInDB string
	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Surname, &passwordStoredInDB, &user.Version)
	if err == sql.ErrNoRows {
		return nil, ErrNoRecord
	} /* TODO: when password is hashed this function would have to be changed */
	if password != passwordStoredInDB {
		return nil, ErrWrongPassword
	}
	return user, nil
}

// Get retrives a user from the database
func (r *Repo) Get(userID string) (*User, error) {
	user := &User{}
	row := r.db.QueryRow("SELECT id, email, name, surname, version FROM users WHERE id=$1", userID)
	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Surname, &user.Version)
	if err == sql.ErrNoRows {
		return nil, ErrNoRecord
	}
	return user, nil
}
