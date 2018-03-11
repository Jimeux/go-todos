package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-todos/app/user"
)

const (
	UsernameParam = "username"
	PasswordParam = "password"
	TokenHeader   = "X-Auth-Token"
)

type Handler struct {
	userRepository user.Repository
	authService    Service
}

func NewHandler(userRepository user.Repository, authService Service) Handler {
	return Handler{userRepository, authService}
}

func (h *Handler) Login(context *gin.Context) {
	username := context.PostForm(UsernameParam)
	password := context.PostForm(PasswordParam)
	h.authenticate(context, username, password)
}

func (h *Handler) Logout(context *gin.Context) {
	token := context.Request.Header.Get(TokenHeader)
	err := h.authService.RevokeToken(token)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
	} else {
		context.Status(http.StatusOK)
	}
}

func (h *Handler) Register(context *gin.Context) {
	username := context.PostForm(UsernameParam)
	password := context.PostForm(PasswordParam)

	_, err := h.userRepository.Create(username, password)

	if err != nil {
		context.AbortWithStatus(http.StatusNotAcceptable)
	} else {
		h.authenticate(context, username, password)
	}
}

func (h *Handler) authenticate(context *gin.Context, username string, password string) {
	token, err := h.authService.Authenticate(username, password)

	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
	} else if token == "" {
		context.AbortWithStatus(http.StatusForbidden)
	} else {
		context.JSON(http.StatusOK, token)
	}
}
