package session

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

// DBSession is an abstraction on top of a connection to database
type DBSession struct {
	DB *sql.DB
}

// NewDBSession returns an instance of DBSession
func NewDBSession(db *sql.DB) *DBSession {
	return &DBSession{
		DB: db,
	}
}

// Manager is a placeholder for methods that each session's representation needs to implement
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

// Create creates a session and stores it in the databse
func (sdb *DBSession) Create(rw http.ResponseWriter, UserID string) error {
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
		Domain:  "localhost", /* TODO: possibly is not necessary */
	}
	http.SetCookie(rw, cookie)
	return nil
}

// Check checks that a session exists in the database
func (sdb *DBSession) Check(r *http.Request) (*Session, error) {
	sessID, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		fmt.Println("from Check:", err)
		return nil, ErrorNoAuth
	}
	sess := &Session{}
	row := sdb.DB.QueryRow("SELECT user_id from sessions WHERE id=$1", sessID.Value)
	err = row.Scan(&sess.UserID)
	if err == sql.ErrNoRows {
		return nil, ErrorNoAuth
	} else if err != nil {
		return nil, err
	}
	sess.ID = sessID.Value
	return sess, nil
}

// DestroyCurrent removes current user's session from the database
func (sdb *DBSession) DestroyCurrent(rw http.ResponseWriter, r *http.Request) error {
	sess, err := FromContext(r.Context())
	switch {
	case err == ErrorNoAuth:
		return ErrorNoAuth
	case err != nil:
		return errors.New("Internal server error")
	}
	_, err = sdb.DB.Exec("DELETE FROM sessions WHERE id=$1", sess.ID)
	if err != nil {

	}
	cookie := http.Cookie{
		Name:    "session_id",
		Expires: time.Now().AddDate(0, 0, -1),
		Path:    "/",
	}
	http.SetCookie(rw, &cookie)
	return nil
}

// DestroyAll removes all sessions of a current user from the databse
func (sdb *DBSession) DestroyAll(rw http.ResponseWriter, user *users.User) error {
	return nil
}

type ctxKey int

const sessionKey ctxKey = 1

// FromContext returns session from context if it exists
func FromContext(ctx context.Context) (*Session, error) {
	sess, ok := ctx.Value(sessionKey).(*Session)
	if !ok {
		return nil, ErrorNoAuth
	}
	return sess, nil
}

var noAuthUrls = map[string]struct{}{
	"/api/auth/login":       {},
	"/api/auth/signup":      {},
	"/courses":              {},
	"/api/authors":          {},
	"/courses/{id}":         {},
	"/":                     {},
	"/api/admin/add/author": {}, /* TODO: remove after testing with curl */
}

// AuthMiddleware is responsible for checking whether or not a user has rights to access a resource
func AuthMiddleware(sm Manager, next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if _, ok := noAuthUrls[r.URL.Path]; ok {
			next.ServeHTTP(rw, r)
			return
		}
		sess, err := sm.Check(r)
		fmt.Println(r)
		if err != nil {
			http.Error(rw, "User is not authenticated", http.StatusUnauthorized)
			fmt.Println(err)
			return
		}

		ctx := context.WithValue(r.Context(), sessionKey, sess)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
