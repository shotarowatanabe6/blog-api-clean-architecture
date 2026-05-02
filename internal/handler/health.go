package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IHealthCheckHandler interface {
	Check(c *gin.Context)
}

type HealthCheckHandler struct{}

func NewHealthCheckHandler() IHealthCheckHandler {
	return HealthCheckHandler{}
}

// HealthCheck godoc
//
//	@Summary		Check health status
//	@Description	ヘルスチェック
//	@Success		200	{object}	response
//	@Router			/health [get]
func (h HealthCheckHandler) Check(c *gin.Context) {
	c.JSON(http.StatusOK, response{Message: "ok"})
}
