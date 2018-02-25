package article

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"gin-test/common"
)

type Handler struct {
	repository Repository
}

func NewHandler(repository Repository) Handler {
	return Handler{repository}
}

func (h *Handler) Index(context *gin.Context) {
	articles, err := h.repository.FindAll()

	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
	} else {
		common.Render(
			context,
			gin.H{
				"title":   "Articles",
				"payload": articles,
			},
			"index.html",
		)
	}
}

func (h *Handler) Show(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err == nil {
		article, exists, err := h.repository.FindById(int64(id))

		if !exists {
			context.AbortWithStatus(http.StatusNotFound)
		} else if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		} else {
			common.Render(
				context,
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
				"article.html",
			)
		}
	} else {
		context.AbortWithStatus(http.StatusNotAcceptable)
	}
}
