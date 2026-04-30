package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
//
//	@Summary		Check health status
//	@Description	ヘルスチェック
//	@Success		200	{object}	response
//	@Router			/health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, response{Message: "ok"})
}
