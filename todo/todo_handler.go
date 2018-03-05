package todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"fmt"
)

type Handler struct {
	repository Repository
}

func NewHandler(repository Repository) Handler {
	return Handler{repository}
}

func (h *Handler) Index(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"index.html",
		gin.H{"title": "Todo"},
	)
}

func (h *Handler) List(context *gin.Context) {
	hideComplete, _ := strconv.ParseBool(context.DefaultQuery("hide_complete", "false"))
	todos, err := h.repository.FindAll(hideComplete)

	if err != nil {
		fmt.Println(err)
		context.AbortWithStatus(http.StatusInternalServerError)
	} else {
		context.JSON(http.StatusOK, todos)
	}
}

func (h *Handler) Show(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err == nil {
		todo, err := h.repository.FindById(int64(id))

		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		} else if todo == nil {
			context.AbortWithStatus(http.StatusNotFound)
		} else {
			context.HTML(
				http.StatusOK,
				"todo.html",
				gin.H{
					"title":   todo.Title,
					"payload": todo,
				},
			)
		}
	} else {
		context.AbortWithStatus(http.StatusNotAcceptable)
	}
}

func (h *Handler) Complete(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	completed, err2 := strconv.ParseBool(context.DefaultQuery("completed", "false"))

	if err == nil && err2 == nil {
		affected, err := h.repository.SetComplete(int64(id), completed)

		if err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
			//context.AbortWithStatus(http.StatusInternalServerError)
		} else if affected == 0 {
			context.AbortWithStatus(http.StatusNotFound)
		} else {
			context.JSON(http.StatusOK, "Updated")
		}
	} else {
		context.AbortWithStatus(http.StatusNotAcceptable)
	}
}

func (h *Handler) Create(context *gin.Context) {
	title := context.PostForm("title")
	todo, err := h.repository.Create(title)

	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	} else if todo == nil {
		context.AbortWithStatus(http.StatusNotAcceptable)
	} else {
		context.JSON(http.StatusOK, todo)
	}
}
