package utils

import (
	"github.com/arikama/koran-backend/requestresponse"
	"github.com/gin-gonic/gin"
)

func JsonError(c *gin.Context, status int, err error) {
	c.JSON(status, requestresponse.JsonError{
		Error: err.Error(),
	})
}
