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
	"gin-todos/app"
)

func initDb(env app.Env) *xorm.Engine {
	db, err := xorm.NewEngine("postgres", env.DatabaseHost)

	if err != nil {
		panic("database could not be initialised")
	}

	db.ShowSQL(env.Debug)

	return db
}

func initCache(env app.Env) app.Cache {
	client := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", env.RedisHost)
		},
	}
	return app.NewCache(client)
}

func main() {
	debug := gin.Mode() == "debug"

	env := app.NewEnv(debug)
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
