package todo

import (
	"github.com/Jimeux/go-todos/app/common"
	"github.com/Jimeux/go-todos/app/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	IdParam           = "id"
	TitleParam        = "title"
	CompleteParam     = "complete"
	HideCompleteParam = "hide_complete"
)

type Handler struct {
	logger     common.Logger
	repository Repository
}

func NewHandler(logger common.Logger, repository Repository) *Handler {
	return &Handler{logger, repository}
}

func (h *Handler) Index(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"index.html",
		gin.H{"title": "Todo"},
	)
}

func (h *Handler) List(context *gin.Context) {
	hideComplete, parseErr := parseBool(context, HideCompleteParam, false)

	if parseErr != nil {
		context.AbortWithStatus(http.StatusNotAcceptable)
	} else {
		todos, repoErr := h.repository.FindAll(getUser(context).Id, hideComplete)

		if repoErr != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		} else {
			if todos == nil {
				todos = &[]Model{}
			}
			context.JSON(http.StatusOK, *todos)
		}
	}
}

func (h *Handler) Complete(context *gin.Context) {
	id, paramErr := strconv.Atoi(context.Param(IdParam))
	completed, queryErr := parseBool(context, CompleteParam, false)

	if paramErr != nil || queryErr != nil {
		context.AbortWithStatus(http.StatusNotAcceptable)
	} else {
		affected, err := h.repository.SetComplete(getUser(context).Id, int64(id), completed)

		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		} else if affected == 0 {
			context.AbortWithStatus(http.StatusNotFound)
		} else {
			h.logger.Forward("track.todo.completed", map[string]interface{}{
				"completed": completed,
			})
			context.JSON(http.StatusOK, "Updated")
		}
	}
}

func (h *Handler) Create(context *gin.Context) {
	title := context.PostForm(TitleParam)

	if title == "" {
		context.AbortWithStatus(http.StatusNotAcceptable)
	} else {
		todo, err := h.repository.Create(getUser(context).Id, title)

		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		} else {
			context.JSON(http.StatusOK, todo)
		}
	}
}

func getUser(context *gin.Context) *user.Model {
	return context.MustGet("user").(*user.Model)
}

func parseBool(context *gin.Context, key string, defaultValue bool) (bool, error) {
	boolString := strconv.FormatBool(defaultValue)
	value := context.DefaultQuery(key, boolString)
	return strconv.ParseBool(value)
}
