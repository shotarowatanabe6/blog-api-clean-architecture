package handler

import (
	"blog-api-clean-architecture/internal/domain/service"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	HealthCheck(c *gin.Context)

	GetUser(c *gin.Context)
}

type Handler struct {
	Service service.IUserService
}

func NewHandler(service service.IUserService) IHandler {
	return Handler{service}
}
