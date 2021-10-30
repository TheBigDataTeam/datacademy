package authors

import (
	"errors"
	"fmt"
	"time"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Repo represents a database.
type Repo struct {
	mongo      *mgo.Session
	collection *mgo.Collection
}

// New returns an instance of Author repository.
func New(mongo *mgo.Session, collection *mgo.Collection) *Repo {
	return &Repo{
		mongo:      mongo,
		collection: collection,
	}
}

var (
	// ErrorNoRecord is returned when no record in database is found
	ErrorNoRecord = errors.New("No author record found")
	// ErrorAuthorAlreadyExists is returned when one tries to add author which already exists
	ErrorAuthorAlreadyExists = errors.New("Author already exists")
	// ErrorBadRequest is returned when wrong data provided
	ErrorBadRequest = errors.New("Wrong data provided")
)

// GetAuthors returns list of authors
func (r *Repo) GetAuthors() ([]*Author, error) {
	authors := []*Author{}
	err := r.collection.Find(bson.M{}).All(&authors)
	if err != nil {
		return nil, err
	}
	return authors, nil
}

// GetAuthorByID returns a single author which matches the id provided
func (r *Repo) GetAuthorByID(id bson.ObjectId) (*Author, error) {
	author := &Author{}
	err := r.collection.Find(bson.M{"_id": id}).One(&author)
	if err != nil {
		return nil, err
	}
	return author, nil
}

// GetAuthorByName returns a single author which matches the name provided
func (r *Repo) GetAuthorByName(name string) (*Author, error) {
	author := &Author{}
	err := r.collection.Find(bson.M{"fullname": name}).One(&author)
	if err != nil {
		return nil, err
	}
	fmt.Println(author)
	return author, nil
}

/* Administration part of repository functions */

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
	fmt.Println("from repo: ", a)
	err = r.collection.Insert(a)
	if err != nil {
		return ErrorBadRequest
	}
	return nil
}

// UpdateAuthor replaces an author in the DB with the given item.
func (r *Repo) UpdateAuthor(a Author) error {

	return nil
}

// DeleteAuthor deletes the author from the DB
func (r *Repo) DeleteAuthor(id string) error {

	return nil
}
