package main

import (
	"gin-test/todo"
)

func initializeRoutes(handler* todo.Handler) {

	router.GET("/", handler.Index)

	router.GET("/todo", handler.List)

	router.GET("/todo/:id", handler.Show)

	router.GET("/todo/:id/complete", handler.Complete)

	router.POST("/todo", handler.Create)

}
