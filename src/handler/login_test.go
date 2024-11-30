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

	mux.ServeHTTP(
		w,
		r,
	)

	assert.Equal(t, "/error?code=401&msg=Wrong username or password", w.Header().Get("HX-Redirect"))

	_, _ = dep.Mngr_user.Create(username, password)

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(form.Encode()))

	mux.ServeHTTP(
		w,
		r,
	)
}
