package user

import (
	"gin-todos/app/common"
	"testing"
)

var (
	db         = common.InitTestDb()
	repository = &RepositoryImpl{db}
)
