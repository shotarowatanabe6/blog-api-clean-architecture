package internal

import (
	"fmt"

	_ "blog-api-clean-architecture/docs"
	"blog-api-clean-architecture/internal/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run(port string) error {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", handler.HealthCheck)

	err := r.Run(":" + port)
	return fmt.Errorf("failed to run server: %w", err)
}
