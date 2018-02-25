package article

type Repository interface {
	Create(title string, content string) (*Model, error)
	FindAll() (*[]Model, error)
	FindById(id int64) (*Model, bool, error)
}
