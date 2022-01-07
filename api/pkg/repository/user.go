package repository

import (
	"bytes"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/argon2"

	"github.com/Serj1c/datalearn/api/pkg/entity"
	"github.com/Serj1c/datalearn/api/pkg/errs"
	"github.com/Serj1c/datalearn/api/pkg/util"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(email string, name string, surname string, password string) (string, error) {
	id := util.RandString()
	salt := util.RandString()
	hashedPassword := hashPassword(password, salt)

	row, err := r.db.Exec(`INSERT into users(id, email, name, surname, password) 
		VALUES($1, $2, $3, $4, $5)`, id, email, name, surname, hashedPassword)
	if err != nil {
		fmt.Println(err)
		return "", errs.BadRequest
	}
	affected, err := row.RowsAffected()
	if affected == 0 {
		return "", errs.AlreadyExists
	}
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *UserRepository) Authenticate(email string, password string) (string, error) {
	row := r.db.QueryRow("SELECT id, email, name, surname, password, version FROM users WHERE email=$1", email)
	user, err := checkPassIsValid(row, password)
	switch err {
	case nil:
	case errs.NotFound:
		return "", errs.NotFound
	case errs.WrongPassword:
		return "", errs.WrongPassword
	}
	return user.ID, nil
}

func checkPassIsValid(row *sql.Row, password string) (*entity.User, error) {
	user := &entity.User{}
	var passStoredInDB string
	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Surname, &passStoredInDB, &user.Version)
	if err == sql.ErrNoRows {
		return nil, errs.NotFound
	}
	salt := passStoredInDB[0:24]
	if !bytes.Equal(hashPassword(password, salt), []byte(passStoredInDB)) {
		return nil, errs.WrongPassword
	}
	return user, nil
}

func hashPassword(password string, salt string) []byte {
	hashedPassword := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)
	temp := []byte(salt)
	return append(temp, hashedPassword...)
}

func (r *UserRepository) Get(id string) (*entity.User, error) {
	user := &entity.User{}
	row := r.db.QueryRow("SELECT id, email, name, surname, version FROM users WHERE id=$1", id)
	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Surname, &user.Version)
	if err == sql.ErrNoRows {
		return nil, errs.NotFound
	}
	return user, nil
}
