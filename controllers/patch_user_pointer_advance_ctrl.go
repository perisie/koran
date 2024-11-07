package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"perisie.com/koran/constants"
	"perisie.com/koran/managers"
	"perisie.com/koran/requestresponse"
	"perisie.com/koran/utils"
)

func PatchUserPointerAdvanceCtrl(userManager managers.UserManager) func(*gin.Context) {
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

		currentPointer, err := userManager.AdvanceUserCurrentPointer(request.Email, accessToken)
		if err != nil {
			utils.JsonError(c, http.StatusInternalServerError, err)
			return
		}

		utils.JsonData(c, http.StatusOK, requestresponse.PatchUserPointerAdvanceResponse{
			CurrentPointer: currentPointer,
		})
	}
}
