package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Serj1c/datalearn/api/pkg/errs"
	"github.com/Serj1c/datalearn/api/pkg/middleware"
	"github.com/Serj1c/datalearn/api/pkg/operation"
	"github.com/Serj1c/datalearn/api/pkg/session"
)

type UserHandler struct {
	l *log.Logger
	v *middleware.Validation
	r operation.UserRepository
	s session.Manager
}

func NewUserHandler(l *log.Logger, v *middleware.Validation, r operation.UserRepository, s session.Manager) *UserHandler {
	return &UserHandler{l, v, r, s}
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

func (u *UserHandler) Signup(rw http.ResponseWriter, r *http.Request) {
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

	userID, err := u.r.Create(newUser.Email, newUser.Name, newUser.Surname, newUser.Password)
	switch err {
	case nil:
	case errs.BadRequest:
		http.Error(rw, "Wrong data provided", http.StatusBadRequest)
	case errs.AlreadyExists:
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

func (u *UserHandler) Login(rw http.ResponseWriter, r *http.Request) {
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
	case errs.NotFound:
		http.Error(rw, "User does not exist", http.StatusBadRequest)
	case errs.WrongPassword:
		http.Error(rw, "Password is not correct", http.StatusBadRequest)
	}
	if err != nil {
		return
	}
	u.s.Create(rw, userID)
}

// Get returns a user based on a userID coming from query params.
// Currently is not used
func (u *UserHandler) Get(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	userID := queryParams["user"][0]
	user, err := u.r.Get(userID)
	switch err {
	case nil:
	case errs.NotFound:
		http.Error(rw, "User does not exist", http.StatusBadRequest)
	}
	response, err := json.Marshal(user)
	if err != nil {
		http.Error(rw, "Cannot marshal response", http.StatusInternalServerError)
	}
	rw.Write(response)
}

func (u *UserHandler) GetBySessionID(rw http.ResponseWriter, r *http.Request) {
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
	case errs.NotFound:
		http.Error(rw, "User does not exist", http.StatusBadRequest)
	}
	response, err := json.Marshal(user)
	if err != nil {
		http.Error(rw, "Cannot marshal response", http.StatusInternalServerError)
	}
	rw.Write(response)
}

func (u *UserHandler) Logout(rw http.ResponseWriter, r *http.Request) {
	u.s.DestroyCurrent(rw, r)
}
