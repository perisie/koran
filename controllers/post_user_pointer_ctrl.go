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

func PostUserPointerCtrl(userManager managers.UserManager) func(*gin.Context) {
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

		user, err := userManager.GetUser(accessToken)
		if err != nil {
			utils.JsonError(c, http.StatusNotFound, err)
			return
		}

		if request.Email != user.Email {
			utils.JsonError(c, http.StatusUnauthorized, errors.New("access token mismatch"))
			return
		}

		utils.JsonData(c, http.StatusOK, requestresponse.GetUserPointerResponse{
			CurrentPointer: user.CurrentPointer,
		})
	}
}
