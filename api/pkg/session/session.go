package session

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/Serj1c/datalearn/api/pkg/entity"
	"github.com/Serj1c/datalearn/api/pkg/util"
)

type Session struct {
	ID     string
	UserID string
}

type DBSession struct {
	DB *sql.DB
}

func New(db *sql.DB) *DBSession {
	return &DBSession{
		DB: db,
	}
}

// Manager is a placeholder for methods that each session's representation needs to implement
type Manager interface {
	Create(http.ResponseWriter, string) error
	Check(*http.Request) (*Session, error)
	DestroyCurrent(http.ResponseWriter, *http.Request) error
	DestroyAll(http.ResponseWriter, *entity.User) error
}

var (
	// ErrorNoAuth is an error which is raised when there is no such a session in the database
	ErrorNoAuth = errors.New("No session found")
)

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

func (sdb *DBSession) Check(r *http.Request) (*Session, error) {
	sessID, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
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

func (sdb *DBSession) DestroyAll(rw http.ResponseWriter, user *entity.User) error {
	return nil
}

type ctxKey int

const sessionKey ctxKey = 1

func FromContext(ctx context.Context) (*Session, error) {
	sess, ok := ctx.Value(sessionKey).(*Session)
	if !ok {
		return nil, ErrorNoAuth
	}
	return sess, nil
}

var noAuthUrls = map[string]struct{}{
	"/":                {},
	"/api/auth/login":  {},
	"/api/auth/signup": {},
	"/api/courses":     {},
	"/api/authors":     {},
	"/courses/{id}":    {},
}

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
