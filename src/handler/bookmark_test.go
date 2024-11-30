package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"perisie.com/koran/src/guice"
)

func Test_bookmark(t *testing.T) {
	dep := guice.Dep_test(
		"../template",
		"../../qurancsv",
		"../../static",
	)

	user, _ := dep.Mngr_user.Create("cglotr", "")

	mux := Mux(dep)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/bookmark", nil)
	r.AddCookie(&http.Cookie{
		Name:  "username",
		Value: user.Username,
	})
	r.AddCookie(&http.Cookie{
		Name:  "token",
		Value: user.Password,
	})

	mux.ServeHTTP(
		w,
		r,
	)

	b, _ := io.ReadAll(w.Body)
	b_str := string(b)

	assert.True(t, strings.Contains(b_str, "1:1"))

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodPatch, "/bookmark?move=next", nil)
	r.AddCookie(&http.Cookie{
		Name:  "username",
		Value: user.Username,
	})
	r.AddCookie(&http.Cookie{
		Name:  "token",
		Value: user.Password,
	})

	mux.ServeHTTP(
		w,
		r,
	)

	b, _ = io.ReadAll(w.Body)
	b_str = string(b)

	assert.True(t, strings.Contains(b_str, "1:2"))

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodPatch, "/bookmark?move=prev", nil)
	r.AddCookie(&http.Cookie{
		Name:  "username",
		Value: user.Username,
	})
	r.AddCookie(&http.Cookie{
		Name:  "token",
		Value: user.Password,
	})

	mux.ServeHTTP(
		w,
		r,
	)

	b, _ = io.ReadAll(w.Body)
	b_str = string(b)

	assert.True(t, strings.Contains(b_str, "1:1"))
}

func Test_bookmark_error(t *testing.T) {
	dep := guice.Dep_test(
		"../template",
		"../../qurancsv",
		"../../static",
	)

	mux := Mux(dep)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/bookmark", nil)

	mux.ServeHTTP(
		w,
		r,
	)

	b, _ := io.ReadAll(w.Body)
	b_str := string(b)

	assert.True(t, strings.Contains(b_str, "401"))
	assert.True(t, strings.Contains(b_str, "please login"))
}
