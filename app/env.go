package app

import (
	"os"
	"fmt"
)

type Env struct {
	Debug        bool
	RedisHost    string
	DatabaseHost string
	ViewDir      string
	AssetDir     string
	FluentdHost  string
	FluentdPort  string
}

func NewEnv(debug bool) Env {
	databaseHost := dbUrl()
	redisHost := redisUrl()
	viewDir := os.Getenv("VIEW_DIR")
	assetDir := os.Getenv("ASSET_DIR")
	fluentdHost := os.Getenv("FLUENTD_HOST")
	fluentdPort := os.Getenv("FLUENTD_PORT")

	if debug {
		databaseHost = "postgresql://default:default@127.0.0.1/todos?sslmode=disable"
		redisHost = "127.0.0.1:6379"
		viewDir = "public/views"
		assetDir = "public/assets"
		fluentdHost = "127.0.0.1"
		fluentdPort = "24224"
	}

	return Env{
		debug,
		redisHost,
		databaseHost,
		viewDir,
		assetDir,
		fluentdHost,
		fluentdPort,
	}
}

func redisUrl() string {
	return fmt.Sprintf(
		"%s:%s",
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
	)
}

func dbUrl() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}
