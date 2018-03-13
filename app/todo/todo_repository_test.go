package todo

import (
	"github.com/Jimeux/go-todos/app/common"
	"testing"
)

var (
	db = common.InitTestDb()
	repository = &RepositoryImpl{db}
)

func TestFindAll(t *testing.T) {
	db.Exec("DROP TABLE articles")
	db.Sync2(new(Model))

	var inserts = make([]Model, 5)
	for i := 1; i <= 5; i++  {
		inserted, _ := repository.Create("test" + string(i))
		inserts = append(inserts, *inserted)
	}

	articles, _ := repository.FindAll(false)

	var contains = func(article *Model) bool {
		for _, a := range inserts {
			if a.Title == article.Title {
				return true
			}
		}
		return false
	}

	check(t, len(*articles) == 5,"length was not 5 but", len(*articles))

	for _, article := range *articles {
		check(t, contains(&article),"Mismatch between local and DB.")
	}
}

func check(t *testing.T, condition bool, msg string, values ...interface{}) {
	if !condition {
		t.Fatal(msg, values)
	}
}
