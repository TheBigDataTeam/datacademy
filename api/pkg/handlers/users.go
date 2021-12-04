package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Serj1c/datalearn/api/pkg/middleware"
	"github.com/Serj1c/datalearn/api/pkg/session"
	"github.com/Serj1c/datalearn/api/pkg/users"
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
func (u *Users) Signup(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Cannot read body", http.StatusInternalServerError)
	}
	r.Body.Close()
	newUser := &signUpInfo{}
	err = json.Unmarshal(body, newUser)
	if err != nil {
		http.Error(rw, "Error unmarshaling request body", http.StatusInternalServerError)
	}

	if err = u.Validate(*newUser); err != nil {
		http.Error(rw, fmt.Sprintf("Error: %s", err), http.StatusBadRequest)
	}

	userID, err := u.r.Create(newUser.Email, newUser.Name, newUser.Surname, newUser.Password)
	switch err {
	case nil:
	case users.ErrBadRequest:
		http.Error(rw, "Wrong data provided", http.StatusBadRequest)
	case users.ErrAlreadyExists:
		http.Error(rw, "Such user already exists", http.StatusForbidden)
	default:
		http.Error(rw, "Internal error", http.StatusInternalServerError)
	}
	if err != nil {
		return
	}
	u.s.Create(rw, userID)
	rw.WriteHeader(http.StatusCreated)
}

// Login handles requests for user's authorization
func (u *Users) Login(rw http.ResponseWriter, r *http.Request) {
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

	if err = u.Validate(); err != nil {
		http.Error(rw, fmt.Sprintf("Error: %s", err), http.StatusBadRequest)
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

// Get returns a user based on a userID coming from query params.
// Currently is not used
func (u *Users) Get(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	userID := queryParams["user"][0]
	user, err := u.r.Get(userID)
	switch err {
	case nil:
	case users.ErrNoRecord:
		http.Error(rw, "User does not exist", http.StatusBadRequest)
	}
	response, err := json.Marshal(user)
	if err != nil {
		http.Error(rw, "Cannot marshal response", http.StatusInternalServerError)
	}
	rw.Write(response)
}

// GetBySessionID returns a user based on a session id
func (u *Users) GetBySessionID(rw http.ResponseWriter, r *http.Request) {
	sess, err := u.s.Check(r)
	switch {
	case err == nil:
	case err == session.ErrorNoAuth:
		http.Error(rw, "Not authenticated", http.StatusUnauthorized)
	case err != nil:
		http.Error(rw, "Internal error", http.StatusInternalServerError)
	}
	userID := sess.UserID
	user, err := u.r.Get(userID)
	switch err {
	case nil:
	case users.ErrNoRecord:
		http.Error(rw, "User does not exist", http.StatusBadRequest)
	}
	response, err := json.Marshal(user)
	if err != nil {
		http.Error(rw, "Cannot marshal response", http.StatusInternalServerError)
	}
	rw.Write(response)
}

// Logout handles requests for a user log out
func (u *Users) Logout(rw http.ResponseWriter, r *http.Request) {
	u.s.DestroyCurrent(rw, r)
}

// Validate validates the input provided
func (u *Users) Validate(user users.User) error {
	err := validate.Struct(user)
	return err
}
