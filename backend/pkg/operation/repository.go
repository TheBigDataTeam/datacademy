package operation

import (
	"github.com/globalsign/mgo/bson"

	"github.com/Serj1c/datalearn/api/pkg/entity"
)

type UserRepository interface {
	Create(email string, name string, surname string, password string) (string, error)
	Authenticate(email string, password string) (string, error)
	Get(id string) (*entity.User, error)
}

type AuthorRepository interface {
	List() ([]entity.Author, error)
	Get(id bson.ObjectId) (*entity.Author, error)
	GetByName(name string) (*entity.Author, error)
	Create(a entity.Author) error
	Update(a entity.Author) error
	Delete(id string) error
}

type CourseRepository interface {
	List() ([]*entity.Course, error)
	Get(id string) (*entity.Course, error)
	Create(a entity.Course) error
	Update(a entity.Course) error
	Delete(id string) error
}

type ModuleRepository interface {
	List(courseId string) ([]*entity.Module, error)
	Get(id string) (*entity.Module, error)
	Create(m entity.Module) error
	Update(m entity.Module) error
	Delete(id string) error
}
