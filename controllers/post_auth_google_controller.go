package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/arikama/koran-backend/beans"
	"github.com/arikama/koran-backend/managers"
	"github.com/arikama/koran-backend/utils"
	"github.com/gin-gonic/gin"
)

func PostAuthGoogleController(googleAuthManager managers.GoogleAuthManager) func(c *gin.Context) {
	return func(c *gin.Context) {
		bytes, _ := ioutil.ReadAll(c.Request.Body)
		code := string(bytes)
		googleUser, err := googleAuthManager.AuthUserCode(code)
		if err != nil {
			utils.JsonError(c, http.StatusInternalServerError, err)
			return
		}
		user := beans.User{
			Email:   googleUser.Email,
			Name:    googleUser.Name,
			Token:   googleUser.Token,
			Picture: googleUser.Picture,
		}
		utils.JsonData(c, http.StatusOK, user)
	}
}
