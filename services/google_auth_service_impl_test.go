package services_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
	"perisie.com/koran/services"
)

func TestAuthUserCode(t *testing.T) {
	setupEnv()

	googleAuthManager, err := services.NewGoogleAuthServiceImpl()
	assert.Nil(t, err)

	_, err = googleAuthManager.AuthUserCode("")
	assert.NotNil(t, err)
}

func TestGetGoogleUser(t *testing.T) {
	setupEnv()

	googleAuthManager, err := services.NewGoogleAuthServiceImpl()
	assert.Nil(t, err)

	token := oauth2.Token{
		AccessToken:  "ya29.a0AVA9y1smrq6OrsNpOajYOTqvZuUTQY6MsRw36KW6VUNzlo2qS9G_9piFGJ7b_qbRMZQHHrSeqgh97tVexGk8ii-jleHLffOIzN2aRaeY71O7aX4YZf48G-19uv0JT4as_BjDiAF9a6krbOCGw_z3206P2QFEaCgYKATASAQASFQE65dr8ikCDzIgBW2hOjVEode3CBQ0163",
		TokenType:    "Bearer",
		RefreshToken: "1//0gDCylQ0gm3YrCgYIARAAGBASNwF-L9IrXkyNlKv1GTtLoH31fR5kNlJFuhUQqf15FYbeUPwFS8DcR6M_M8oetg6wjhTS75Qg-Oo",
		Expiry:       time.UnixMilli(0),
	}

	_, err = googleAuthManager.GetGoogleUser(&token)
	assert.NotNil(t, err)
}

func setupEnv() {
	os.Setenv("GOOGLE_CLIENT_ID", "454337127208-8cfsr6ebdake7qjp93n98rlrjjm9qgo6.apps.googleusercontent.com")
	os.Setenv("GOOGLE_CLIENT_SECRET", "")
	os.Setenv("GOOGLE_REDIRECT_URL", "http://localhost:3000")
}
