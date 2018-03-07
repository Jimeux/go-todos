package todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"fmt"
)

const IdParam = "id"
const TitleParam = "title"
const CompleteParam = "complete"
const HideCompleteParam = "hide_complete"

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
	hideComplete, parseErr := parseBool(context, HideCompleteParam,false)

	if parseErr != nil {
		context.AbortWithStatus(http.StatusNotAcceptable)
	} else {
		todos, repoErr := h.repository.FindAll(hideComplete)

		if repoErr != nil {
			fmt.Println(repoErr)
			context.AbortWithStatus(http.StatusInternalServerError)
		} else {
			context.JSON(http.StatusOK, todos)
		}
	}
}

func (h *Handler) Complete(context *gin.Context) {
	id, paramErr := strconv.Atoi(context.Param(IdParam))
	completed, queryErr := parseBool(context, CompleteParam, false)

	fmt.Println(completed, queryErr)

	if paramErr != nil || queryErr != nil {
		context.AbortWithStatus(http.StatusNotAcceptable)
	} else {
		affected, err := h.repository.SetComplete(int64(id), completed)

		if err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
		} else if affected == 0 {
			context.AbortWithStatus(http.StatusNotFound)
		} else {
			context.JSON(http.StatusOK, "Updated")
		}
	}
}

func (h *Handler) Create(context *gin.Context) {
	title := context.PostForm(TitleParam)

	if title == "" {
		context.AbortWithStatus(http.StatusNotAcceptable)
	} else {
		todo, err := h.repository.Create(title)

		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		} else {
			context.JSON(http.StatusOK, todo)
		}
	}
}

func parseBool(context *gin.Context, key string, defaultValue bool) (bool, error) {
	boolString := strconv.FormatBool(defaultValue)
	value := context.DefaultQuery(key, boolString)

	fmt.Println(key)

	return strconv.ParseBool(value)
}
