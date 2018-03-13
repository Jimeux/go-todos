package main

import (
	_ "github.com/lib/pq"
	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/Jimeux/go-todos/app/user"
	"github.com/Jimeux/go-todos/app/auth"
	"github.com/Jimeux/go-todos/app"
	"github.com/Jimeux/go-todos/app/todo"
	"time"
	"strconv"
)

func initDb(env app.Env) *xorm.Engine {
	db, err := xorm.NewEngine("postgres", env.DatabaseHost)
	if err != nil {
		panic("database could not be initialised: " + err.Error())
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

func initLogger(env app.Env) app.Logger {
	fluentdPort, _ := strconv.Atoi(env.FluentdPort)
	logger, err := fluent.New(fluent.Config{
		FluentPort: fluentdPort,
		FluentHost: env.FluentdHost,
	})

	if err != nil {
		panic("Fluentd logger could not be initialised: " + err.Error())
	}
	return app.NewLogger(logger)
}

func main() {
	debug := gin.Mode() == "debug"

	env := app.NewEnv(debug)
	db := initDb(env)
	cache := initCache(env)
	logger := initLogger(env)

	todoRepository := todo.NewRepository(db)
	userRepository := user.NewRepository(db)

	authService := auth.NewService(cache, userRepository)

	todoHandler := todo.NewHandler(logger, todoRepository)
	authHandler := auth.NewHandler(userRepository, authService)

	router := gin.Default()

	router.LoadHTMLGlob(env.ViewDir + "/*")
	router.Static("/assets", env.AssetDir)

	initializeRoutes(router, todoHandler, authHandler, authService)

	defer cache.Close()
	defer db.Close()
	defer logger.Close()

	router.Run()
}
