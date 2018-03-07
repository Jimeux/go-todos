package main

import (
	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"runtime"
	"gin-todos/todo"
	"os"
)

var (
	db             = initDb()
	todoRepository = todo.NewRepository(db)
	todoHandler    = todo.NewHandler(todoRepository)
	router         = gin.Default()
)

func initDb() *xorm.Engine {
	const driverName = "postgres"
	dataSourceName := os.Getenv("DATABASE_URL")

	if len(dataSourceName) == 0 {
		dataSourceName = "postgres://localhost:5433/gin_todos?user=default&password=default&sslmode=disable"
	}

	db, err := xorm.NewEngine(driverName, dataSourceName)
	err2 := db.Sync2(new(todo.Model))

	if err2 != nil {
		panic("couldn't sync todos table: \n" + err2.Error())
	}
	if err != nil {
		panic("error initialising xorm engine: \n" + err.Error())
	}

	return db
}

func main() {
	runtime.GOMAXPROCS(2)

	viewDir := os.Getenv("VIEW_DIR")

	if len(viewDir) == 0 {
		viewDir = "views"
	}

	router.LoadHTMLGlob(viewDir + "/*")
	router.Static("/assets", "./assets")

	initializeRoutes(&todoHandler)

	router.Run()
}
