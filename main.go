package main

import (
	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"runtime"
	"gin-todos/todo"
	"os"
	"time"
	"fmt"
)

func initDb() (*xorm.Engine, error) {
	const driverName = "postgres"
	dataSourceName := os.Getenv("DATABASE_URL")

	databaseHost := os.Getenv("DB_HOST")

	if len(databaseHost) == 0 {
		databaseHost = "127.0.0.1:5433"
	}

	dataSourceName = "postgresql://default:default@" + databaseHost + "/gin_todos?sslmode=disable"
	fmt.Println(dataSourceName)

	db, engineErr := xorm.NewEngine(driverName, dataSourceName)

	if engineErr != nil {
		fmt.Println("error initialising xorm engine: \n" + engineErr.Error())
		return nil, engineErr
	}

	syncErr := db.Sync2(new(todo.Model))

	if syncErr != nil {
		fmt.Println("couldn't sync todos table: \n" + syncErr.Error())
		return nil, syncErr
	}

	return db, nil
}

func main() {
	runtime.GOMAXPROCS(2)

	viewDir := os.Getenv("VIEW_DIR")
	assetsDir := os.Getenv("ASSETS_DIR")

	if len(viewDir) == 0 {
		viewDir = "views"
	}
	if len(assetsDir) == 0 {
		assetsDir = "assets"
	}

	var db *xorm.Engine

	for i := 1; i <= 10; i++ {
		dbTry, err := initDb()
		if err == nil {
			db = dbTry
			break
		} else {
			time.Sleep(time.Duration(i) * time.Second)
		}
	}

	todoRepository := todo.NewRepository(db)
	todoHandler := todo.NewHandler(todoRepository)
	router := gin.Default()

	router.LoadHTMLGlob(viewDir + "/*")
	router.Static("/assets", assetsDir)

	initializeRoutes(router, &todoHandler)

	router.Run()
}
