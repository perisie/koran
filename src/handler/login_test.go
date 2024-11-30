package handler

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"perisie.com/koran/src/guice"
)

func Test_login(t *testing.T) {
	dep := guice.Dep_test(
		"../template",
		"../../qurancsv",
		"../../static",
	)
	mux := Mux(dep)
	w := httptest.NewRecorder()

	mux.ServeHTTP(
		w,
		httptest.NewRequest(http.MethodGet, "/login", nil),
	)

	b, _ := io.ReadAll(w.Body)
	b_str := string(b)

	assert.True(t, strings.Contains(b_str, "Login"))

	username := "cglotr"
	password := "password"

	w = httptest.NewRecorder()
	form := url.Values{}
	form.Add("username", username)
	form.Add("password", password)
	r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	mux.ServeHTTP(
		w,
		r,
	)

	assert.Equal(t, "/error?code=401&msg=Wrong username or password", w.Header().Get("HX-Redirect"))

	user, err := dep.Mngr_user.Create(username, password)

	assert.Nil(t, err)
	assert.Equal(t, username, user.Username)

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	mux.ServeHTTP(
		w,
		r,
	)

	b, _ = io.ReadAll(w.Body)
	b_str = string(b)

	assert.True(t, strings.Contains(b_str, "Logging in..."))

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodDelete, "/login", nil)

	mux.ServeHTTP(
		w,
		r,
	)

	b, _ = io.ReadAll(w.Body)
	b_str = string(b)

	assert.True(t, strings.Contains(b_str, "Logging out..."))

}
