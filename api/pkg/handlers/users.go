package handlers

import (
	"encoding/json"
	"io/ioutil"
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

type signUpInfo struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
}

// Signup handles request for creating new users
func (u *Users) Signup(w http.ResponseWriter, r *http.Request) {
	u.l.Println("[SUCCESS] post request has come")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read body", http.StatusInternalServerError)
	}
	newUser := &signUpInfo{}
	err = json.Unmarshal(body, newUser)
	if err != nil {
		http.Error(w, "Error unmarshaling request body", http.StatusInternalServerError)
	}

	salt := util.RandString()
	hashedPassword := string(u.hashPassword(newUser.Password, salt))
	u.l.Println(hashedPassword)

	userID, err := u.r.Create(newUser.Email, newUser.Name, newUser.Surname, newUser.Password)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}
	u.s.Create(w, userID)
}
