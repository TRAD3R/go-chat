package handler

import (
	"github.com/gin-gonic/gin"
	"github/trad3r/go_temp.git/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.GET("/log-in", h.Login)
		auth.POST("/sign-up", h.Signin)
	}

	return router
}
