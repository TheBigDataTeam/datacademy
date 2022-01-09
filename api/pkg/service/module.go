package service

import (
	"github.com/Serj1c/datalearn/api/pkg/entity"
	"github.com/Serj1c/datalearn/api/pkg/operation"
)

type ModuleProcessor struct {
	mr operation.ModuleRepository
}

func NewModuleProcessor(mr operation.ModuleRepository) *ModuleProcessor {
	return &ModuleProcessor{
		mr: mr,
	}
}

func (mp *ModuleProcessor) List() ([]*entity.Module, error) {
	modules, err := mp.mr.List()
	if err != nil {
		return nil, err
	}
	return modules, nil
}

func (mp *ModuleProcessor) Get(id string) (*entity.Module, error) {
	module, err := mp.mr.Get(id)
	if err != nil {
		return nil, err
	}
	return module, nil
}

func (mp *ModuleProcessor) Create(m entity.Module) error {
	err := mp.mr.Create(m)
	return err
}
func (mp *ModuleProcessor) Update(m entity.Module) error {
	return nil
}
func (mp *ModuleProcessor) Delete(id string) error {
	return nil
}
