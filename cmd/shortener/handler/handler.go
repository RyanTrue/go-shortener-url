package handler

import (
	"github.com/RyanTrue/go-shortener-url.git/cmd/shortener/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.ServiceContainer
}

func NewHandler(services *service.ServiceContainer) *Handler {
	return &Handler{
		services: services,
	}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/", h.ShortenURL)
	router.GET("/:id", h.ExpandURL)

	return router
}
