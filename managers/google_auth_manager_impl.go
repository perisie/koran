package managers

import (
	"context"
	"os"

	"github.com/arikama/koran-backend/daos"
	"github.com/arikama/koran-backend/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	googleauth "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type GoogleAuthManagerImpl struct {
	context context.Context
	config  oauth2.Config
	userDao daos.UserDao
}

func NewGoogleAuthManagerImpl() (*GoogleAuthManagerImpl, error) {
	context := context.Background()
	config := oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"email", "profile", "openid"},
		Endpoint:     google.Endpoint,
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
	}
	return &GoogleAuthManagerImpl{
		context: context,
		config:  config,
	}, nil
}

func (g *GoogleAuthManagerImpl) AuthUserCode(userAuthCode string) (*GoogleUser, error) {
	token, err := g.config.Exchange(g.context, userAuthCode)
	if err != nil {
		return nil, err
	}
	googleUser, err := g.GetGoogleUser(token)
	if err != nil {
		return nil, err
	}

	_, err = g.userDao.UpsertUser(&models.User{
		Email:   googleUser.Email,
		Name:    googleUser.Name,
		Token:   googleUser.Token,
		Picture: googleUser.Picture,
	})
	if err != nil {
		return nil, err
	}

	return googleUser, nil
}

func (g *GoogleAuthManagerImpl) GetGoogleUser(userToken *oauth2.Token) (*GoogleUser, error) {
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
