package todo

import (
	"github.com/go-xorm/xorm"
	"time"
)

type XormRepository struct {
	db *xorm.Engine
}

func NewRepository(db *xorm.Engine) Repository {
	return &XormRepository{db}
}

func (r *XormRepository) Create(title string) (*Model, error) {
	todo := Model{
		Title:    title,
		Complete: false,
		Created:  time.Now(),
	}
	_, err := r.db.Insert(&todo)
	return &todo, err
}

func (r *XormRepository) SetComplete(id int64, complete bool) (int64, error) {
	return r.db.
		Id(id).
		Cols("complete").
		Update(&Model{Complete: complete})
}

func (r *XormRepository) FindAll(hideComplete bool) (*[]Model, error) {
	var todos []Model
	var err error

	if hideComplete {
		err = r.db.Where("complete = false").
			Desc("created").
			Find(&todos)
	} else {
		err = r.db.Desc("created").
			Find(&todos)
	}
	return &todos, err
}

func (r *XormRepository) FindById(id int64) (*Model, error) {
	todo := Model{Id: id}
	_, err := r.db.Get(&todo)
	return &todo, err
}

func (r *XormRepository) Delete(id int64) error {
	todo := Model{Id: id}
	_, err := r.db.Delete(&todo)
	return err
}
