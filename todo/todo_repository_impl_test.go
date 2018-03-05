package todo

import (
	"testing"
	"gin-test/common"
)

var (
	db = common.InitTestDb()
	repository = &XormRepository{db}
)

func TestFindById(t *testing.T) {
	db.Exec("DROP TABLE articles")
	db.Sync2(new(Model))

	inserted, _ := repository.Create("test", "content")

	article, exists, _ := repository.FindById(inserted.Id)

	check(t, exists, "article not found")
	check(t, article.Title == inserted.Title, "Title not equal", article.Title, inserted.Title)
	check(t, article.Content == inserted.Content, "Content not equal", article.Content, inserted.Content)
}

func TestFindAll(t *testing.T) {
	db.Exec("DROP TABLE articles")
	db.Sync2(new(Model))

	var inserts = make([]Model, 5)
	for i := 1; i <= 5; i++  {
		inserted, _ := repository.Create("test" + string(i), "content" + string(i))
		inserts = append(inserts, *inserted)
	}

	articles, _ := repository.FindAll()

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
