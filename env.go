package main

import (
	"os"
)

type Env struct {
	RedisHost      string
	DataSourceName string
	ViewDir        string
	AssetDir       string
}

func NewEnv() Env {
	datasourceName := dbUrl()
	redisHost := os.Getenv("REDIS_HOST")
	viewDir := os.Getenv("VIEW_DIR")
	assetDir := os.Getenv("ASSET_DIR")

	if len(redisHost) == 0 {
		redisHost = "127.0.0.1"
	}
	if len(viewDir) == 0 {
		viewDir = "public/views"
	}
	if len(assetDir) == 0 {
		assetDir = "public/assets"
	}

	return Env{
		redisHost + ":6379",
		datasourceName,
		viewDir,
		assetDir,
	}
}

func dbUrl() string {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" {
		dbHost = "127.0.0.1"
	}
	if dbUser == "" {
		dbUser = "default"
	}
	if dbPass == "" {
		dbPass = "default"
	}
	if dbName == "" {
		dbName = "gin_todos"
	}

	return "postgresql://" + dbUser + ":" + dbPass + "@" + dbHost +
		":5432/" + dbName + "?sslmode=disable"
}
