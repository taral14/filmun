package auth

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service}
}

func (h *Handler) RegisterHTTPEndpoints(r *gin.Engine) {
	r.Use(h.AuthMiddleware)
	r.POST("/sign-in", h.SignIn)
}

type login struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (h *Handler) SignIn(c *gin.Context) {
	var login login
	err := c.BindJSON(&login)
	if err != nil {
		log.Printf("SignIn: %v", err.Error())
		return
	}
	user, token, err := h.service.LogIn(login.Username, login.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": "Login or password is incorrect"})
	} else {
		c.JSON(200, gin.H{
			"token":    token,
			"userId":   user.ID,
			"username": user.Username,
		})
	}
}

func (h *Handler) AuthMiddleware(c *gin.Context) {
	tokenHeader := c.GetHeader("Authorization")
	if tokenHeader == "" { // if empty header do not try auth user
		return
	}
	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	tokenString := splitted[1]
	userId, err := h.service.GetUserIdByToken(tokenString)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	} else {
		c.Set("UserId", userId)
	}
	c.Next()
}
