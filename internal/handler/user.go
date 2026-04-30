package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getUserRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (h Handler) GetUser(c *gin.Context) {
	var req getUserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		fmt.Printf("failed to bind request: %s", err.Error()) // FIXME: ロガーを導入する
		c.JSON(http.StatusBadRequest, response{Message: "bad request"})
		return
	}

	user, err := h.Service.GetUser(req.ID)
	if err != nil {
		fmt.Printf("failed to get user: %s", err.Error()) // FIXME: ロガーを導入する
		c.JSON(http.StatusInternalServerError, response{Message: "failed to get user"})
		return
	}

	c.JSON(http.StatusOK, user)
}
