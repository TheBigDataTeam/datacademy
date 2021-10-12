package authors

import (
	"errors"
	"time"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Repo represents a data base
type Repo struct {
	mongo      *mgo.Session
	collection *mgo.Collection
}

// NewRepo returns an instance of a Repo
func NewRepo(mongo *mgo.Session, collection *mgo.Collection) *Repo {
	return &Repo{
		mongo:      mongo,
		collection: collection,
	}
}

var (
	// ErrNoRecord is returned when no record in database is found
	ErrNoRecord = errors.New("No author record found")
	// ErrorAuthorAlreadyExists is returned when one tries to add user which already exists
	ErrorAuthorAlreadyExists = errors.New("Author already exists")
	// ErrorBadRequest is returned when one tries to create a user with a wrong data
	ErrorBadRequest = errors.New("Error inserting info into db")
)

// AddAuthor inserts a new author in the DB
func (r *Repo) AddAuthor(a Author) error {
	tempAuthor := &Author{}
	err := r.collection.Find(bson.M{"email": a.Email}).One(&tempAuthor)
	if err == nil {
		return ErrorAuthorAlreadyExists
	}
	/* TODO: here is possibly mistake in querying db not handled */
	a.ID = bson.NewObjectId()
	a.CreatedOn = time.Now().Format("2006-01-02 15:04:05")
	a.Version = 1
	err = r.collection.Insert(a)
	if err != nil {
		return ErrorBadRequest
	}
	return nil
}

// GetAuthors returns list of authors
func (r *Repo) GetAuthors() ([]*Author, error) {
	authorsFromDB := []*Author{}
	err := r.collection.Find(bson.M{}).All(&authorsFromDB)
	if err != nil {
		return nil, err
	}
	return authorsFromDB, nil
}

// GetAuthorByID returns a single author which matches the id from the DB
func (r *Repo) GetAuthorByID(id string) (*Author, error) {

	return nil, nil
}

// UpdateAuthor replaces an author in the DB with the given item.
func (r *Repo) UpdateAuthor(a Author) error {

	return nil
}

// DeleteAuthor deletes the author from the DB
func (r *Repo) DeleteAuthor(id string) error {

	return nil
}
