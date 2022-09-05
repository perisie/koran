package controllers

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
			RedirectURL:  "http://localhost:3000",
		}
		bytes, _ := ioutil.ReadAll(c.Request.Body)
		code := string(bytes)
		tok, err := conf.Exchange(ctx, code)
		if err != nil {
			log.Fatal(err)
		}
		oauth2Service, err := googleauth.NewService(ctx, option.WithTokenSource(conf.TokenSource(ctx, tok)))
		if err != nil {
			log.Fatal(err)
		}
		userInfo, err := oauth2Service.Userinfo.Get().Do()
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"token":     tok.AccessToken,
			"user_info": userInfo,
		})
	}
}
