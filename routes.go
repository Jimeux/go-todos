package main

import (
	"gin-todos/todo"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine, handler* todo.Handler) {

	router.GET("/", handler.Index)

	router.GET("/todo", handler.List)

	router.GET("/todo/:id/complete", handler.Complete)

	router.POST("/todo", handler.Create)

}
