package service

import (
	"github.com/Serj1c/datalearn/api/pkg/entity"
	"github.com/Serj1c/datalearn/api/pkg/operation"
)

type CourseProcessor struct {
	cr operation.CourseRepository
}

func NewCourseProcessor(cr operation.CourseRepository) *CourseProcessor {
	return &CourseProcessor{
		cr: cr,
	}
}

func (cp *CourseProcessor) List() ([]*entity.Course, error) {
	courses, err := cp.cr.List()
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (cp *CourseProcessor) Get(id string) (*entity.Course, error) {
	course, err := cp.cr.Get(id)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (cp *CourseProcessor) Create(c entity.Course) error {
	err := cp.cr.Create(c)
	return err
}
func (cp *CourseProcessor) Update(m entity.Course) error {
	return nil
}
func (cp *CourseProcessor) Delete(id string) error {
	return nil
}
