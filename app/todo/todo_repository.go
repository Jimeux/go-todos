package todo

import (
	"github.com/go-xorm/xorm"
	"time"
)

type Repository interface {
	Create(userId int64, title string) (*Model, error)
	FindAll(userId int64, hideCompleted bool) (*[]Model, error)
	SetComplete(userId int64, todoId int64, complete bool) (int64, error)
	Delete(id int64) error
}

func NewRepository(db *xorm.Engine) Repository {
	return &RepositoryImpl{db}
}

type RepositoryImpl struct {
	db *xorm.Engine
}

func (r *RepositoryImpl) Create(userId int64, title string) (*Model, error) {
	todo := Model{
		UserId:   userId,
		Title:    title,
		Complete: false,
		Created:  time.Now(),
	}
	_, err := r.db.Insert(&todo)
	return &todo, err
}

func (r *RepositoryImpl) SetComplete(userId int64, todoId int64, complete bool) (int64, error) {
	return r.db.
		Id(todoId).
		Where("user_id = ?", userId).
		Cols("complete").
		Update(&Model{Complete: complete})
}

func (r *RepositoryImpl) FindAll(userId int64, hideComplete bool) (*[]Model, error) {
	var todos []Model
	var query string

	if hideComplete {
		query = "complete = false"
	}

	err := r.db.
		Where("user_id = ?", userId).
		And(query).
		Desc("created").
		Find(&todos)

	return &todos, err
}

func (r *RepositoryImpl) Delete(id int64) error {
	todo := Model{Id: id}
	_, err := r.db.Delete(&todo)
	return err
}
