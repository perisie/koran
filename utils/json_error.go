package utils

import (
	"github.com/arikama/koran-backend/beans"
	"github.com/gin-gonic/gin"
)

func JsonError(c *gin.Context, status int, err error) {
	c.JSON(status, beans.JsonError{
		Error: err.Error(),
	})
}
