package repository

import (
	"time"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/Serj1c/datalearn/api/pkg/entity"
	"github.com/Serj1c/datalearn/api/pkg/errs"
)

type AuthorRepository struct {
	mongo      *mgo.Session
	collection *mgo.Collection
}

func NewAuthorRepository(mongo *mgo.Session, collection *mgo.Collection) *AuthorRepository {
	return &AuthorRepository{
		mongo:      mongo,
		collection: collection,
	}
}

func (r *AuthorRepository) List() ([]entity.Author, error) {
	authors := []entity.Author{}
	err := r.collection.Find(bson.M{}).All(&authors)
	if err != nil {
		return nil, errs.OperationFailed(500, "authors LIST failed", err.Error(), nil)
	}
	return authors, nil
}

func (r *AuthorRepository) Get(id bson.ObjectId) (*entity.Author, error) {
	author := &entity.Author{}
	err := r.collection.Find(bson.M{"_id": id}).One(&author)
	if err != nil {
		return nil, errs.NotFound
	}
	return author, nil
}

func (r *AuthorRepository) GetByName(name string) (*entity.Author, error) {
	author := &entity.Author{}
	err := r.collection.Find(bson.M{"fullname": name}).One(&author)
	if err != nil {
		return nil, errs.NotFound
	}
	return author, nil
}

/* Administration part of repository functions */

func (r *AuthorRepository) Create(a entity.Author) error {
	tempAuthor := &entity.Author{}
	err := r.collection.Find(bson.M{"email": a.Email}).One(&tempAuthor)
	if err == nil {
		return errs.AlreadyExists
	}
	/* TODO: here is possibly mistake in querying db not handled */
	a.ID = bson.NewObjectId()
	a.CreatedOn = time.Now().Format("2006-01-02 15:04:05")
	a.Version = 1
	err = r.collection.Insert(a)
	if err != nil {
		return errs.BadRequest
	}
	return nil
}

func (r *AuthorRepository) Update(a entity.Author) error {
	return nil
}

func (r *AuthorRepository) Delete(id string) error {
	return nil
}
