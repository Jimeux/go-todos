package article

import "github.com/go-xorm/xorm"

type XormRepository struct {
	db *xorm.Engine
}

func NewRepository(db *xorm.Engine) Repository {
	return &XormRepository{db}
}

func (r *XormRepository) Create(title string, content string) (*Model, error) {
	article := Model{
		Title:   title,
		Content: content,
	}
	_, err := r.db.Insert(article)
	return &article, err
}

func (r *XormRepository) FindAll() (*[]Model, error) {
	var articles []Model
	err := r.db.Find(&articles)
	return &articles, err
}

func (r *XormRepository) FindById(id int64) (*Model, bool, error) {
	article := Model{Id: id}
	exists, err := r.db.Get(&article)
	return &article, exists, err
}
