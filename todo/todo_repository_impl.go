package todo

import (
	"github.com/go-xorm/xorm"
	"time"
)

type RepositoryImpl struct {
	db *xorm.Engine
}

func NewRepository(db *xorm.Engine) Repository {
	return &RepositoryImpl{db}
}

func (r *RepositoryImpl) Create(title string) (*Model, error) {
	todo := Model{
		Title:    title,
		Complete: false,
		Created:  time.Now(),
	}
	_, err := r.db.Insert(&todo)
	return &todo, err
}

func (r *RepositoryImpl) SetComplete(id int64, complete bool) (int64, error) {
	return r.db.
		Id(id).
		Cols("complete").
		Update(&Model{Complete: complete})
}

func (r *RepositoryImpl) FindAll(hideComplete bool) (*[]Model, error) {
	var todos []Model
	var query string

	if hideComplete {
		query = "complete = false"
	}

	err := r.db.
		Where(query).
		Desc("created").
		Find(&todos)

	return &todos, err
}

func (r *RepositoryImpl) Delete(id int64) error {
	todo := Model{Id: id}
	_, err := r.db.Delete(&todo)
	return err
}
