package film

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service}
}

func (h *Handler) RegisterHTTPEndpoints(r *gin.Engine) {
	r.GET("/actors/:id/films", h.ActorFilms)
	r.GET("/directors/:id/films", h.DirectorFilms)
}

func (h *Handler) DirectorFilms(c *gin.Context) {
	directorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)

		c.AbortWithStatus(404)
		return
	}
	films, err := h.service.ListDirectorFilms(directorId)
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, films)
}

func (h *Handler) ActorFilms(c *gin.Context) {
	actorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)

		c.AbortWithStatus(404)
		return
	}
	films, err := h.service.ListActorFilms(actorId)
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, films)
}
