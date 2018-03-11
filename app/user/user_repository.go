package user

import "github.com/go-xorm/xorm"

type Repository interface {
	Create(username string, password string) (*Model, error)
	FindByCredentials(username string, password string) (*Model, error)
}

type RepositoryImpl struct {
	db *xorm.Engine
}

func NewRepository(db *xorm.Engine) Repository {
	return &RepositoryImpl{db}
}

func (r *RepositoryImpl) Create(username string, password string) (*Model, error) {
	todo := Model{
		Username: username,
		Password: password,
	}
	_, err := r.db.Insert(&todo)
	return &todo, err
}

func (r *RepositoryImpl) FindByCredentials(username string, password string) (*Model, error) {
	user := Model{ // TODO: Stop blank values returning a row
		Username: username,
		Password: password,
	}
	_, err := r.db.Get(&user)
	return &user, err
}
