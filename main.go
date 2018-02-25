package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"runtime"
	"gin-test/article"
)

const driverName = "postgres"
const dataSourceName = "postgres://localhost:5432/xorm_test?sslmode=disable"

var (
	db = initDb()
	articleRepository = article.NewRepository(db)
	articleHandler = article.NewHandler(articleRepository)
	router = gin.Default()
)

func initDb() *xorm.Engine {
	db, err := xorm.NewEngine(driverName, dataSourceName)

	db.Sync2(new(article.Model))

	if err != nil {
		panic("couldn't connect to database: \n" + err.Error())
	} else {
		return db
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	router.LoadHTMLGlob("views/*")

	initializeRoutes(&articleHandler)

	router.Run()
}
