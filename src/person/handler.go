package person

import (
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
	r.GET("/persons/:id", h.GetPerson)
	r.GET("/persons", func(c *gin.Context) {
		if term := c.Query("term"); term == "" {
			h.ListPersons(c)
		} else {
			h.SearchByTerm(c)
		}
	})
}

func (h *Handler) GetPerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {

	}
	person, _ := h.service.GetPerson(id)
	c.JSON(200, person)
}

func (h *Handler) ListPersons(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	persons, _ := h.service.ListPersons(page)
	c.JSON(200, persons)
}

func (h *Handler) SearchByTerm(c *gin.Context) {
	term := c.Query("term")
	persons, _ := h.service.SearchByTerm(term)
	c.JSON(200, persons)
}
