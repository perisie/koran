package controllers

import (
	"errors"
	"net/http"

	"github.com/arikama/koran-backend/constants"
	"github.com/arikama/koran-backend/requestresponse"
	"github.com/arikama/koran-backend/utils"
	"github.com/gin-gonic/gin"
)

func PatchUserPointerAdvanceCtrl() func(*gin.Context) {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get(constants.XAccessToken())
		if accessToken == "" {
			utils.JsonError(c, http.StatusUnauthorized, errors.New(`missing x-access-token header`))
			return
		}
		utils.JsonData(c, http.StatusOK, requestresponse.PatchUserPointerAdvanceResponse{
			CurrentPointer: accessToken,
		})
	}
}
