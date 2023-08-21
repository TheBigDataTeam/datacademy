package module

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Serj1c/datalearn/api/pkg/errs"
	"github.com/Serj1c/datalearn/api/pkg/middleware"
	"github.com/Serj1c/datalearn/api/pkg/service"
)

type ModuleHandler struct {
	l *log.Logger
	v *middleware.Validation
	p *service.ModuleProcessor
}

func NewModuleHandler(l *log.Logger, v *middleware.Validation, p *service.ModuleProcessor) *ModuleHandler {
	return &ModuleHandler{
		l: l,
		v: v,
		p: p,
	}
}

func (m *ModuleHandler) List(rw http.ResponseWriter, r *http.Request) {
	req, reqErr := m.parseListRequest(r)
	if reqErr != nil {
		http.Error(rw, "", http.StatusBadRequest)
	}
	modules, err := m.p.List(req.CourseId)
	switch {
	case err == nil:
	case err == errs.BadRequest:
		http.Error(rw, "Wrong data provided", http.StatusBadRequest)
	}
	response, err := json.Marshal(modules)
	if err != nil {
		http.Error(rw, "Error marshaling response", http.StatusInternalServerError)
	}
	rw.Write(response)
}

func (m *ModuleHandler) Get(rw http.ResponseWriter, r *http.Request) {
	req, reqErr := m.parseGetRequest(r)
	if reqErr != nil {
		http.Error(rw, "", http.StatusBadRequest)
	}
	mdl, err := m.p.Get(req.Id)
	switch err {
	case nil:
	case errs.NotFound:
		http.Error(rw, "There is no such a module", http.StatusNotFound)
	}
	response, err := json.Marshal(mdl)
	if err != nil {
		http.Error(rw, "Error marshaling response", http.StatusInternalServerError)
	}
	rw.Write(response)
}

func (m *ModuleHandler) Create(rw http.ResponseWriter, r *http.Request) {}

func (m *ModuleHandler) Delete(rw http.ResponseWriter, r *http.Request) {}

func (m *ModuleHandler) Update(rw http.ResponseWriter, r *http.Request) {}

func (m *ModuleHandler) parseGetRequest(r *http.Request) (req *GetRequest, e *errs.InternalError) {
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		e = errs.OperationFailed(http.StatusBadRequest, "...", err.Error(), nil)
	}
	return
}

func (m *ModuleHandler) parseListRequest(r *http.Request) (req *ListRequest, e *errs.InternalError) {
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		e = errs.OperationFailed(http.StatusBadRequest, "...", err.Error(), nil)
	}
	return
}
