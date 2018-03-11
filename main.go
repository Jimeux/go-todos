package main

import (
	_ "github.com/lib/pq"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"gin-todos/app/todo"
	"time"
	"gin-todos/app/user"
	"gin-todos/app/auth"
)

func initDb(env Env) *xorm.Engine {
	db, err := xorm.NewEngine("postgres", env.DataSourceName)

	if err != nil {
		panic("database could not be initialised")
	}

	db.ShowSQL(false)

	return db
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
	env := NewEnv()
	db := initDb(env)
	cache := initCache(env)

	todoRepository := todo.NewRepository(db)
	userRepository := user.NewRepository(db)

	authService := auth.NewService(cache, userRepository)

	todoHandler := todo.NewHandler(todoRepository)
	authHandler := auth.NewHandler(userRepository, authService)

	router := gin.Default()
	router.LoadHTMLGlob(env.ViewDir + "/*")
	router.Static("/assets", env.AssetDir)

	initializeRoutes(router, &todoHandler, &authHandler, authService)

	defer cache.Close()
	defer db.Close()

	router.Run()
}
