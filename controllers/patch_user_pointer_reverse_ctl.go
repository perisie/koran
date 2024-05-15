package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/arikama/koran-backend/constants"
	"github.com/arikama/koran-backend/managers"
	"github.com/arikama/koran-backend/requestresponse"
	"github.com/arikama/koran-backend/utils"
	"github.com/gin-gonic/gin"
)

func PatchUserPointerReverseCtrl(userManager managers.UserManager) func(*gin.Context) {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get(constants.XAccessToken())
		if accessToken == "" {
			utils.JsonError(c, http.StatusUnauthorized, errors.New(`missing x-access-token header`))
			return
		}
		request := requestresponse.PatchUserPointerAdvanceRequest{}
		err := json.NewDecoder(c.Request.Body).Decode(&request)
		if err != nil {
			utils.JsonError(c, http.StatusBadRequest, err)
			return
		}
		currentPointer, err := userManager.ReverseUserCurrentPointer(request.Email, accessToken)
		if err != nil {
			utils.JsonError(c, http.StatusInternalServerError, err)
			return
		}
		utils.JsonData(c, http.StatusOK, requestresponse.PatchUserPointerAdvanceResponse{
			CurrentPointer: currentPointer,
		})
	}
}
