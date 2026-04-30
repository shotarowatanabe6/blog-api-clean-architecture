// Description: Clean Architecture の Compotion Root として扱う
package main

import (
	"blog-api-clean-architecture/internal/domain/repository"
	"blog-api-clean-architecture/internal/domain/service"
	"blog-api-clean-architecture/internal/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Blog API
//	@version		1.0
//	@description	This is a blog server.

//	@contact.name	shotarowatanabe6
//	@contact.email	shotarowatanabe6@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/
func main() {
	port := "8080"

	cacheRepo := repository.NewCacheRepository()
	service := service.NewUserService(cacheRepo)
	handler := handler.NewHandler(service)

	r := newRouter(handler)

	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}

func newRouter(h handler.IHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", h.HealthCheck)

	{
		userGroup := r.Group("/api/v1")
		userGroup.GET("/users/:id", h.GetUser)
	}

	return r
}
