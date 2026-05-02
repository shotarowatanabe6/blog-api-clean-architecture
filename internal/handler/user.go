package handler

import (
	"fmt"
	"net/http"

	"blog-api-clean-architecture/internal/domain/service"

	"github.com/gin-gonic/gin"
)

type getUserRequest struct {
	ID string `uri:"id" binding:"required"`
}

type IUserHandler interface {
	FindByID(c *gin.Context)
}

type UserHandler struct {
	Service service.IUserService
}

func NewUserHandler(service service.IUserService) IUserHandler {
	return UserHandler{service}
}

func (h UserHandler) FindByID(c *gin.Context) {
	var req getUserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		fmt.Printf("failed to bind request: %s", err.Error()) // FIXME: ロガーを導入する
		c.JSON(http.StatusBadRequest, response{Message: "bad request"})
		return
	}

	user, err := h.Service.FindByID(req.ID)
	if err != nil {
		fmt.Printf("failed to get user: %s", err.Error()) // FIXME: ロガーを導入する
		c.JSON(http.StatusInternalServerError, response{Message: "failed to get user"})
		return
	}

	c.JSON(http.StatusOK, user)
}
