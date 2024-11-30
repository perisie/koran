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

func Test_signup(t *testing.T) {
	dep := guice.Dep_test(
		"../template",
		"../../qurancsv",
		"../../static",
	)
	mux := Mux(dep)
	w := httptest.NewRecorder()

	mux.ServeHTTP(
		w,
		httptest.NewRequest(http.MethodGet, "/signup", nil),
	)

	b, _ := io.ReadAll(w.Body)
	b_str := string(b)

	assert.True(t, strings.Contains(b_str, "Signup"))

	w = httptest.NewRecorder()
	form := url.Values{}
	form.Add("username", "")
	form.Add("password", "")
	r := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBufferString(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	mux.ServeHTTP(
		w,
		r,
	)

	assert.Equal(t, "/error?code=400&msg=username invalid", w.Header().Get("HX-Redirect"))

	username := "cglotr"
	password := "password"

	w = httptest.NewRecorder()
	form = url.Values{}
	form.Add("username", username)
	form.Add("password", password)
	r = httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBufferString(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	mux.ServeHTTP(
		w,
		r,
	)

	assert.Equal(t, "/login", w.Header().Get("HX-Redirect"))

	w = httptest.NewRecorder()

	mux.ServeHTTP(
		w,
		r,
	)

	assert.Equal(t, "/error?code=403&msg=username taken", w.Header().Get("HX-Redirect"))
}
