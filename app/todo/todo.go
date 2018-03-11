package todo

import "time"

type Model struct {
	Id       int64     `json:"id" xorm:"SERIAL pk"`
	UserId   int64     `json:"userId" xorm:"index not null"`
	Title    string    `json:"title" xorm:"not null"`
	Complete bool      `json:"complete" xorm:"not null default false"`
	Created  time.Time `json:"created" xorm:"not null default CURRENT_TIMESTAMP"`
}

func (a Model) TableName() string {
	return "todos"
}
