package article

type Model struct {
	Id      int64  `xorm:"SERIAL" json:"id"`
	Title   string `json:"title"`
	Content string `xorm:"TEXT" json:"content"`
}

func (a Model) TableName() string {
	return "articles"
}
