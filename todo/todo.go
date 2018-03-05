package todo

import "time"

type Model struct {
	Id       int64     `json:"id"`
	Title    string    `json:"title"`
	Complete bool      `json:"complete"`
	Created  time.Time `json:"created"`
}

func (a Model) TableName() string {
	return "todos"
}
