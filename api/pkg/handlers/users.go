package handlers

import (
	"log"
	"net/http"

	"golang.org/x/crypto/argon2"

	"github.com/Serj1c/datalearn/api/pkg/middleware"
	"github.com/Serj1c/datalearn/api/pkg/session"
	"github.com/Serj1c/datalearn/api/pkg/users"
	"github.com/Serj1c/datalearn/api/pkg/util"
)

// Users is a handler for getting and updating authors
type Users struct {
	l *log.Logger
	v *middleware.Validation
	r *users.Repo
	s session.SessionManager
}

// NewUsers creates an authors handler with the given logger
func NewUsers(l *log.Logger, v *middleware.Validation, r *users.Repo, s session.SessionManager) *Users {
	return &Users{l, v, r, s}
}

func (u *Users) hashPassword(password string, salt string) []byte {
	hashedPassword := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)
	temp := []byte(salt)
	return append(temp, hashedPassword...)
}

// Signup handles request for creating new users
func (u *Users) Signup(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")       /* TODO check that login is valid */
	password := r.FormValue("password") /* TODO check that password is valid */
	email := r.FormValue("email")       /* TODO check that email is valid */
	salt := util.RandString()
	hashedPassword := u.hashPassword(password, salt)

	userID, err := u.r.Create(login, email, hashedPassword)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}
	u.s.Create(w, userID)
}
