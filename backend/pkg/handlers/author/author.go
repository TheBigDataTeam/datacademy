package author

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"

	"github.com/Serj1c/datalearn/api/pkg/entity"
	"github.com/Serj1c/datalearn/api/pkg/errs"
	"github.com/Serj1c/datalearn/api/pkg/middleware"
	"github.com/Serj1c/datalearn/api/pkg/service"
)

type AuthorHandler struct {
	l *log.Logger
	v *middleware.Validation
	p *service.AuthorProcessor
}

func NewAuthorHandler(l *log.Logger, v *middleware.Validation, p *service.AuthorProcessor) *AuthorHandler {
	return &AuthorHandler{
		l: l,
		v: v,
		p: p,
	}
}

type authorData struct {
	Author entity.Author `json:"author"`
}

func (a *AuthorHandler) List(rw http.ResponseWriter, r *http.Request) {
	authors, err := a.p.List()
	switch {
	case err == nil:
	case err == errs.BadRequest:
		http.Error(rw, "Wrong data provided", http.StatusBadRequest)
	}
	response, err := json.Marshal(authors)
	if err != nil {
		http.Error(rw, "Error marshaling response", http.StatusInternalServerError)
	}
	rw.Write(response)
}

func (a *AuthorHandler) Get(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["id"]) {
		http.Error(rw, "ID is of a wrong format", http.StatusBadRequest)
		return
	}
	authorID := bson.ObjectIdHex(vars["id"])
	author, err := a.p.Get(authorID)
	switch err {
	case nil:
	case errs.NotFound:
		http.Error(rw, "There is no such an author", http.StatusNotFound)
	}
	response, err := json.Marshal(author)
	if err != nil {
		http.Error(rw, "Error marshaling response", http.StatusInternalServerError)
	}
	rw.Write(response)
}

func (a *AuthorHandler) GetByName(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorName := vars["name"]
	author, err := a.p.GetByName(authorName)
	switch err {
	case nil:
	case errs.NotFound:
		http.Error(rw, "There is no such an author", http.StatusNotFound)
	}
	response, err := json.Marshal(author)
	if err != nil {
		http.Error(rw, "Error marshaling response", http.StatusInternalServerError)
	}
	rw.Write(response)
}

/* Administration endpoints handlers */

func (a *AuthorHandler) Create(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Cannot read body", http.StatusInternalServerError)
	}
	r.Body.Close()

	data := &authorData{}
	err = json.Unmarshal(body, data)
	if err != nil {
		http.Error(rw, "Error unmarshaling request body", http.StatusInternalServerError)
	}

	err = a.p.Create(data.Author)
	switch {
	case err == nil:
		rw.WriteHeader(http.StatusCreated)
	case err == errs.BadRequest:
		http.Error(rw, "Wrong data provided", http.StatusBadRequest)
	case err == errs.AlreadyExists:
		http.Error(rw, "Such author already exists", http.StatusConflict)
	default:
		http.Error(rw, "Internal error", http.StatusInternalServerError)
	}
}

func (a *AuthorHandler) Update(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Cannot read body", http.StatusInternalServerError)
	}
	r.Body.Close()

	data := &authorData{}
	err = json.Unmarshal(body, data)
	if err != nil {
		http.Error(rw, "Error unmarshaling request body", http.StatusInternalServerError)
	}
	_ = a.p.Update(data.Author)

	/* TODO currently this handler is not implemented */

	http.Error(rw, "Update resourse is not allowed", http.StatusForbidden)
}

func (a *AuthorHandler) Delete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["id"]) {
		http.Error(rw, "ID is of a wrong format", http.StatusBadRequest)
		return
	}
	authorID := bson.ObjectIdHex(vars["id"])
	_ = a.p.Delete(authorID)

	/* TODO currently this handler is not implemented */

	http.Error(rw, "Delete resourse is not allowed", http.StatusForbidden)
}
