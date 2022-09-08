package controllers

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/arikama/koran-backend/beans"
	"github.com/arikama/koran-backend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	googleauth "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func PostAuthGoogleController() func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.Background()
		conf := &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Scopes:       []string{"email", "profile", "openid"},
			Endpoint:     google.Endpoint,
			RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		}
		bytes, _ := ioutil.ReadAll(c.Request.Body)
		code := string(bytes)
		token, err := conf.Exchange(ctx, code)
		if err != nil {
			utils.JsonError(c, http.StatusBadRequest, err)
			return
		}
		oauth2Service, err := googleauth.NewService(ctx, option.WithTokenSource(conf.TokenSource(ctx, token)))
		if err != nil {
			utils.JsonError(c, http.StatusBadRequest, err)
			return
		}
		userInfo, err := oauth2Service.Userinfo.Get().Do()
		if err != nil {
			utils.JsonError(c, http.StatusBadRequest, err)
			return
		}
		user := beans.User{
			Email:   userInfo.Email,
			Name:    userInfo.GivenName,
			Token:   token.AccessToken,
			Picture: userInfo.Picture,
		}
		utils.JsonData(c, http.StatusOK, user)
	}
}
