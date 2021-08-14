package session

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/Serj1c/datalearn/api/pkg/users"
	"github.com/Serj1c/datalearn/api/pkg/util"
)

type Session struct {
	ID     string
	UserID string
}

type SessionDB struct {
	DB *sql.DB
}

func NewSessionDB(db *sql.DB) *SessionDB {
	return &SessionDB{
		DB: db,
	}
}

type SessionManager interface {
	Create(http.ResponseWriter, string) error
	Check(*http.Request) (*Session, error)
	DestroyCurrent(http.ResponseWriter, *http.Request) error
	DestroyAll(http.ResponseWriter, *users.User) error
}

var (
	ErrorNoAuth = errors.New("No session found")
)

func (sdb *SessionDB) Create(w http.ResponseWriter, aUserID string) error {
	sessionID := util.RandString()
	theUserID := aUserID
	_, err := sdb.DB.Exec("INSERT into sessions(id, user_id) VALUES($1, $2)", sessionID, theUserID)
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

func (sbd *SessionDB) Check(r *http.Request) (*Session, error) {
	return &Session{}, nil
}

func (sbd *SessionDB) DestroyCurrent(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (sbd *SessionDB) DestroyAll(w http.ResponseWriter, user *users.User) error {
	return nil
}
