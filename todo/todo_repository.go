package todo

type Repository interface {
	Create(title string) (*Model, error)
	FindAll(hideCompleted bool) (*[]Model, error)
	SetComplete(id int64, complete bool) (int64, error)
	Delete(id int64) error
}
