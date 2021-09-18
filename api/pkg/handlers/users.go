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

// Users is a handler for getting and updating users
type Users struct {
	l *log.Logger
	v *middleware.Validation
	r *users.Repo
	s session.Manager
}

// NewUsers creates User's handler
func NewUsers(l *log.Logger, v *middleware.Validation, r *users.Repo, s session.Manager) *Users {
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

type loginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Signup handles request for creating new users
func (u *Users) Signup(w http.ResponseWriter, r *http.Request) {
	u.l.Println("[ATTENTION] request to create a new user has come")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read body", http.StatusInternalServerError)
	}
	r.Body.Close()
	newUser := &signUpInfo{}
	err = json.Unmarshal(body, newUser)
	if err != nil {
		http.Error(w, "Error unmarshaling request body", http.StatusInternalServerError)
	}

	salt := util.RandString()
	hashedPassword := string(u.hashPassword(newUser.Password, salt)) /* TODO: now passwords are not hashed!! */
	u.l.Println(hashedPassword)

	userID, err := u.r.Create(newUser.Email, newUser.Name, newUser.Surname, newUser.Password)
	switch err {
	case nil:
	case users.ErrorBadRequest:
		http.Error(w, "Wrong data provided", http.StatusBadRequest)
	case users.ErrorUserAlreadyExists:
		http.Error(w, "Such user already exists", http.StatusForbidden)
	default:
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}
	if err != nil {
		return
	}
	u.s.Create(w, userID)
	w.WriteHeader(http.StatusCreated)
}

// Login handles requests for user's authorization
func (u *Users) Login(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[ATTENTION] request for user login has come")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Cannot read body", http.StatusInternalServerError)
	}
	r.Body.Close()
	userForAuth := &loginInfo{}
	err = json.Unmarshal(body, userForAuth)
	if err != nil {
		http.Error(rw, "Error unmarshaling request body", http.StatusInternalServerError)
	}
	userID, err := u.r.Authenticate(userForAuth.Email, userForAuth.Password)
	switch err {
	case nil:
	case users.ErrNoRecord:
		http.Error(rw, "User does not exist", http.StatusBadRequest)
	case users.ErrWrongPassword:
		http.Error(rw, "Password is not correct", http.StatusBadRequest)
	}
	if err != nil {
		return
	}
	u.s.Create(rw, userID)
}
