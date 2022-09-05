package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRootController() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "koran-backend at your service!",
		})
	}
}
