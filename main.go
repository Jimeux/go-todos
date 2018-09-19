package main

import (
	"github.com/Jimeux/go-todos/app"
	"github.com/Jimeux/go-todos/app/auth"
	"github.com/Jimeux/go-todos/app/common"
	"github.com/Jimeux/go-todos/app/todo"
	"github.com/Jimeux/go-todos/app/user"
	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"strconv"
)

func initDb(env common.Env) *xorm.Engine {
	db, err := xorm.NewEngine("postgres", env.DatabaseHost)
	if err != nil {
		panic("database could not be initialised: " + err.Error())
	}

	db.ShowSQL(env.Debug)
	return db
}

func initCache(env common.Env) common.Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     env.RedisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return common.NewCache(client)
}

func initLogger(env common.Env) common.Logger {
	fluentdPort, _ := strconv.Atoi(env.FluentdPort)
	logger, err := fluent.New(fluent.Config{
		FluentPort: fluentdPort,
		FluentHost: env.FluentdHost,
	})

	if err != nil {
		panic("Fluentd logger could not be initialised: " + err.Error())
	}
	return common.NewLogger(logger)
}

func main() {
	debug := gin.Mode() == "debug"

	env := common.NewEnv(debug)
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

	app.InitializeRoutes(router, todoHandler, authHandler, authService)

	defer cache.Close()
	defer db.Close()
	defer logger.Close()

	router.Run()
}
