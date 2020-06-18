package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service}
}

func (h *Handler) RegisterHTTPEndpoints(r *gin.Engine) {
	r.GET("/my-profile", h.MyProfile)
}

func (h *Handler) MyProfile(c *gin.Context) {
	userId := c.GetInt("UserId")
	user, err := h.service.FindById(userId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(200, user)
}
