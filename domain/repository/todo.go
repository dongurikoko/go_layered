package repository

import (
	"github.com/dongurikoko/GO_lesson1/domain/model"
)

// infra層、usecase層がこのinterfaceに依存する
type TodoRepository interface {
	Insert(title string) error
	GetAll() ([]*model.Todo, error)
	GetAllByTitle(title string) ([]*model.Todo, error)
	Update(id int, title string) error
	Delete(id int) error
}
