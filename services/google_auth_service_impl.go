package services

import (
	"context"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	googleauth "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type GoogleAuthServiceImpl struct {
	context context.Context
	config  oauth2.Config
}

func NewGoogleAuthServiceImpl() (*GoogleAuthServiceImpl, error) {
	context := context.Background()
	config := oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"email", "profile", "openid"},
		Endpoint:     google.Endpoint,
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
	}
	return &GoogleAuthServiceImpl{
		context: context,
		config:  config,
	}, nil
}

func (g *GoogleAuthServiceImpl) AuthUserCode(userAuthCode string) (*GoogleUser, error) {
	token, err := g.config.Exchange(g.context, userAuthCode)
	if err != nil {
		return nil, err
	}
	return g.GetGoogleUser(token)
}

func (g *GoogleAuthServiceImpl) GetGoogleUser(userToken *oauth2.Token) (*GoogleUser, error) {
	client, err := googleauth.NewService(g.context, option.WithTokenSource(g.config.TokenSource(g.context, userToken)))
	if err != nil {
		return nil, err
	}
	userInfo, err := client.Userinfo.Get().Do()
	if err != nil {
		return nil, err
	}
	user := GoogleUser{
		Email:   userInfo.Email,
		Name:    userInfo.GivenName,
		Token:   userToken.AccessToken,
		Picture: userInfo.Picture,
	}
	return &user, nil
}
