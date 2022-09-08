package managers_test

import (
	"os"
	"testing"
	"time"

	"github.com/arikama/koran-backend/managers"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestAuthUserCode(t *testing.T) {
	setupEnv()

	googleAuthManager, err := managers.NewGoogleAuthManagerImpl()
	assert.Nil(t, err)

	_, err = googleAuthManager.AuthUserCode("")
	assert.NotNil(t, err)
}

func TestGetGoogleUser(t *testing.T) {
	setupEnv()

	googleAuthManager, err := managers.NewGoogleAuthManagerImpl()
	assert.Nil(t, err)

	token := oauth2.Token{
		AccessToken:  "ya29.a0AVA9y1smrq6OrsNpOajYOTqvZuUTQY6MsRw36KW6VUNzlo2qS9G_9piFGJ7b_qbRMZQHHrSeqgh97tVexGk8ii-jleHLffOIzN2aRaeY71O7aX4YZf48G-19uv0JT4as_BjDiAF9a6krbOCGw_z3206P2QFEaCgYKATASAQASFQE65dr8ikCDzIgBW2hOjVEode3CBQ0163",
		TokenType:    "Bearer",
		RefreshToken: "1//0gDCylQ0gm3YrCgYIARAAGBASNwF-L9IrXkyNlKv1GTtLoH31fR5kNlJFuhUQqf15FYbeUPwFS8DcR6M_M8oetg6wjhTS75Qg-Oo",
		Expiry:       time.UnixMilli(0),
	}

	googleUser, err := googleAuthManager.GetGoogleUser(&token)
	assert.Nil(t, err)

	assert.Equal(t, "amir.ariffin.920404@gmail.com", googleUser.Email)
	assert.Equal(t, "Amir", googleUser.Name)
	assert.Equal(t, "https://lh3.googleusercontent.com/a-/AFdZucpNm6qed2bDnlWlNS-SVbsBsuJl0EzqgItwA7jFOn4=s96-c", googleUser.Picture)
	assert.Equal(t, token.AccessToken, googleUser.Token)
}

func setupEnv() {
	os.Setenv("GOOGLE_CLIENT_ID", "454337127208-8cfsr6ebdake7qjp93n98rlrjjm9qgo6.apps.googleusercontent.com")
	os.Setenv("GOOGLE_CLIENT_SECRET", "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")
	os.Setenv("GOOGLE_REDIRECT_URL", "http://localhost:3000")
}
