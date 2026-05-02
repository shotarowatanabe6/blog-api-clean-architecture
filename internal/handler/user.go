package handler

import (
	"fmt"
	"net/http"

	"blog-api-clean-architecture/internal/domain/models"
	"blog-api-clean-architecture/internal/domain/service"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	FindByID(c *gin.Context)
	Save(c *gin.Context)
}

type UserHandler struct {
	Service service.IUserService
}

func NewUserHandler(service service.IUserService) IUserHandler {
	return UserHandler{service}
}

type findUserByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (h UserHandler) FindByID(c *gin.Context) {
	var req findUserByIDRequest
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

func (h UserHandler) Save(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, response{Message: "bad request"})
		return
	}

	u, err := h.Service.Save(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response{Message: "failed to save user data"})
		return
	}

	c.JSON(http.StatusCreated, u)
}
