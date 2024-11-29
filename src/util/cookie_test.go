package util

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_cookie_username_token(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.AddCookie(&http.Cookie{
		Name:  "username",
		Value: "cglotr",
	})
	r.AddCookie(&http.Cookie{
		Name:  "token",
		Value: "token_value",
	})
	username, token := Cookie_username_token(r)
	assert.Equal(t, "cglotr", username)
	assert.Equal(t, "token_value", token)
}
