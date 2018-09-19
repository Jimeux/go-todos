package common

import (
	"fmt"
	"os"
)

type Env struct {
	Debug        bool
	RedisHost    string
	DatabaseHost string
	ViewDir      string
	AssetDir     string
}

func NewEnv(debug bool) Env {
	databaseHost := dbUrl()
	redisHost := redisUrl()
	viewDir := os.Getenv("VIEW_DIR")
	assetDir := os.Getenv("ASSET_DIR")

	if debug {
		databaseHost = "postgresql://default:default@127.0.0.1/todos?sslmode=disable"
		redisHost = "127.0.0.1:6379"
		viewDir = "public/views"
		assetDir = "public/assets"
	}

	return Env{
		debug,
		redisHost,
		databaseHost,
		viewDir,
		assetDir,
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
