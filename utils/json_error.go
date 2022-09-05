package utils

import "github.com/gin-gonic/gin"

func JsonError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"error": err.Error(),
	})
}
