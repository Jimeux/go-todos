package user

type Model struct {
	Id       int64  `json:"id" xorm:"SERIAL pk"`
	Username string `json:"username" xorm:"not null unique"`
	Password string `json:"-" xorm:"not null"`
}

func (a Model) TableName() string {
	return "users"
}
