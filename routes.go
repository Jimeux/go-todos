package main

import "gin-test/article"

func initializeRoutes(handler* article.Handler) {

	router.GET("/", handler.Index)

	router.GET("/article/:id", handler.Show)

}
