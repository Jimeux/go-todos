package main

import (
	"gin-todos/app/todo"
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-todos/app/auth"
)

func initializeRoutes(
	router *gin.Engine,
	todoHandler *todo.Handler,
	authHandler *auth.Handler,
	authService auth.Service,
) {

	authMiddleware := func(c *gin.Context) {
		token := c.Request.Header.Get("X-Auth-Token")
		userModel, err := authService.VerifyToken(token)

		if token == "" || err != nil {
			c.AbortWithStatus(http.StatusForbidden)
		} else {
			c.Set("user", userModel)
			c.Next()
		}
	}

	router.POST("/login", authHandler.Login)

	router.POST("/register", authHandler.Register)

	router.GET("/logout", authHandler.Logout)

	router.GET("/", todoHandler.Index)

	router.GET("/todo", authMiddleware, todoHandler.List)

	router.GET("/todo/:id/complete", authMiddleware, todoHandler.Complete)

	router.POST("/todo", authMiddleware, todoHandler.Create)

}
