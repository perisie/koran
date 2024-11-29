package handler

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"perisie.com/koran/src/guice"
	"strings"
	"testing"
)

func Test_user(t *testing.T) {
	dep := guice.Dep_test(
		"../template",
		"../../qurancsv",
		"../../static",
	)
	mux := Mux(dep)
	w := httptest.NewRecorder()

	mux.ServeHTTP(
		w,
		httptest.NewRequest(http.MethodGet, "/user", nil),
	)

	b, _ := io.ReadAll(w.Body)
	b_str := string(b)

	assert.True(t, strings.Contains(b_str, "401"))
	assert.True(t, strings.Contains(b_str, "please login"))

	user, _ := dep.Mngr_user.Create("cglotr", "")

	w = httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/user", nil)
	r.AddCookie(&http.Cookie{
		Name:  "username",
		Value: user.Username,
	})

	mux.ServeHTTP(
		w,
		r,
	)

	b, _ = io.ReadAll(w.Body)
	b_str = string(b)

	assert.True(t, strings.Contains(b_str, "User"))
	assert.True(t, strings.Contains(b_str, user.Username))
	assert.True(t, strings.Contains(b_str, "Logout"))
}
