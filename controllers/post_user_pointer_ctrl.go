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
