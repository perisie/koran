package utils

import (
	"github.com/gin-gonic/gin"
	"perisie.com/koran/requestresponse"
)

func JsonError(c *gin.Context, status int, err error) {
	c.JSON(status, requestresponse.JsonError{
		Error: err.Error(),
	})
}
