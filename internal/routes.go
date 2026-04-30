package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run(port string) error {
	r := gin.Default()

	r.GET("/health", healthCheck)

	err := r.Run(":" + port)
	return fmt.Errorf("failed to run server: %w", err)
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
