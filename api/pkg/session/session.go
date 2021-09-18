package session

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/Serj1c/datalearn/api/pkg/users"
	"github.com/Serj1c/datalearn/api/pkg/util"
)

// Session represents a session of a user
type Session struct {
	ID     string
	UserID string
}

// DBSession ...
type DBSession struct {
	DB *sql.DB
}

// NewDBSession returns an instance of SessionDB
func NewDBSession(db *sql.DB) *DBSession {
	return &DBSession{
		DB: db,
	}
}

// Manager ...
type Manager interface {
	Create(http.ResponseWriter, string) error
	Check(*http.Request) (*Session, error)
	DestroyCurrent(http.ResponseWriter, *http.Request) error
	DestroyAll(http.ResponseWriter, *users.User) error
}

var (
	// ErrorNoAuth is an error which is raised when there is no such a session in the database
	ErrorNoAuth = errors.New("No session found")
)

var noAuthUrls = map[string]struct{}{
	"api/user/login":  {},
	"api/user/signup": {},
	"/courses":        {},
	"/authors":        {},
	"/courses/{id}":   {},
	"/authors/{id}":   {},
	"/":               {},
}

// Create creates a session and stores it in the databse
func (sdb *DBSession) Create(w http.ResponseWriter, UserID string) error {
	sessionID := util.RandString()
	_, err := sdb.DB.Exec("INSERT into sessions(id, user_id) VALUES($1, $2)", sessionID, UserID)
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sessionID,
		Expires: time.Now().Add(30 * 24 * time.Hour),
		Path:    "/",
	}
	http.SetCookie(w, cookie)
	return nil
}

// Check checks that a session exists in the database
func (sdb *DBSession) Check(r *http.Request) (*Session, error) {
	return &Session{}, nil
}

// DestroyCurrent removes current user's session from the database
func (sdb *DBSession) DestroyCurrent(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// DestroyAll removes all sessions of a current user from the databse
func (sdb *DBSession) DestroyAll(w http.ResponseWriter, user *users.User) error {
	return nil
}

type ctxKey int

const sessionKey ctxKey = 1

// FromContext returns session if exists from context
func FromContext(ctx context.Context) (*Session, error) {
	sess, ok := ctx.Value(sessionKey).(*Session)
	if !ok {
		return nil, ErrorNoAuth
	}
	return sess, nil
}

// AuthMiddleware is responsible for checking whether or not a user has rights to access a resource
func AuthMiddleware(sm Manager, next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if _, ok := noAuthUrls[r.URL.Path]; ok {
			next.ServeHTTP(rw, r)
			return
		}
		sess, err := sm.Check(r)
		if err != nil {
			http.Error(rw, "User is not authenticated", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), sessionKey, sess)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
