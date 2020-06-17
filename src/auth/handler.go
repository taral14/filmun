package auth

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service}
}

func (h *Handler) RegisterHTTPEndpoints(r *gin.Engine) {
	r.GET("/sign-in", h.SignIn)
}

func (h *Handler) SignIn(c *gin.Context) {
	token, err := h.service.LogIn("taral", "9924054")
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(200, token)
	}
}
