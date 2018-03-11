package user

import (
	"testing"
	"gin-todos/app/common"
)

var (
	db = common.InitTestDb()
	repository = &RepositoryImpl{db}
)
