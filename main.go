package main

import (
	_ "github.com/lib/pq"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"runtime"
	"gin-todos/app/todo"
	"time"
	"fmt"
	"gin-todos/app/user"
	"gin-todos/app/auth"
)

func initDb(env Env) (*xorm.Engine, error) {
	const driverName = "postgres"

	db, engineErr := xorm.NewEngine(driverName, env.DataSourceName)
	db.ShowSQL(true)

	if engineErr != nil {
		fmt.Println("error initialising xorm engine: \n" + engineErr.Error())
		return nil, engineErr
	}

	return db, nil
}

func initCache(env Env) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", env.RedisHost)
		},
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	env := NewEnv()

	var db *xorm.Engine

	for i := 1; i <= 10; i++ {
		dbTry, err := initDb(env)
		if err == nil {
			db = dbTry
			break
		} else {
			time.Sleep(time.Duration(i) * time.Second)
		}
	}

	cache := initCache(env)

	router := gin.Default()

	router.LoadHTMLGlob(env.ViewDir + "/*")
	router.Static("/assets", env.AssetDir)

	todoRepository := todo.NewRepository(db)
	userRepository := user.NewRepository(db)

	authService := auth.NewService(cache, userRepository)

	todoHandler := todo.NewHandler(todoRepository)
	authHandler := auth.NewHandler(authService)

	initializeRoutes(router, &todoHandler, &authHandler, authService)

	defer cache.Close()
	defer db.Close()

	router.Run()
}
