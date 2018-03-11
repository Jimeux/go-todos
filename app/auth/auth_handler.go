package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const UsernameParam = "username"
const PasswordParam = "password"

type Handler struct {
	authService Service
}

func NewHandler(authService Service) Handler {
	return Handler{authService}
}

func (h *Handler) Login(context *gin.Context) {
	username := context.PostForm(UsernameParam)
	password := context.PostForm(PasswordParam)

	token, err := h.authService.Authenticate(username, password)

	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
	} else if token == "" {
		context.AbortWithStatus(http.StatusForbidden)
	} else {
		context.JSON(http.StatusOK, token)
	}
}
