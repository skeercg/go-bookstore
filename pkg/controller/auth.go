package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) signUp(ctx *gin.Context) {
	id, err := c.services.AuthService.CreateUser()
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (c *Controller) signIn(ctx *gin.Context) {
	id, err := c.services.AuthService.GetUser()
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
