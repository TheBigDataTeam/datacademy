package service

import (
	"github.com/Serj1c/datalearn/api/pkg/entity"
	"github.com/Serj1c/datalearn/api/pkg/operation"
	"github.com/globalsign/mgo/bson"
)

type AuthorProcessor struct {
	ar operation.AuthorRepository
}

func NewAuthorProcessor(ar operation.AuthorRepository) *AuthorProcessor {
	return &AuthorProcessor{
		ar: ar,
	}
}

func (ap *AuthorProcessor) List() ([]entity.Author, error) {
	authors, err := ap.ar.List()
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (ap *AuthorProcessor) Get(id bson.ObjectId) (*entity.Author, error) {
	authr, err := ap.ar.Get(id)
	if err != nil {
		return nil, err
	}
	return authr, nil
}

func (ap *AuthorProcessor) GetByName(name string) (*entity.Author, error) {
	authr, err := ap.ar.GetByName(name)
	if err != nil {
		return nil, err
	}
	return authr, nil
}

func (ap *AuthorProcessor) Create(a entity.Author) error {
	err := ap.ar.Create(a)
	return err
}
func (ap *AuthorProcessor) Update(a entity.Author) error {
	return nil
}
func (ap *AuthorProcessor) Delete(id bson.ObjectId) error {
	return nil
}
