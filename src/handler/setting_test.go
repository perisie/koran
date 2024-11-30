package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"perisie.com/koran/src/guice"
)

func Test_setting(t *testing.T) {
	dep := guice.Dep_test(
		"../template",
		"../../qurancsv",
		"../../static",
	)
	mux := Mux(dep)

	username := "cglotr"
	password := "password"

	user, _ := dep.Mngr_user.Create(username, password)

	assert.True(t, user.Setting.Bookmark_verse)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPatch, "/setting?name=bookmark_verse&value=false", nil)
	r.AddCookie(&http.Cookie{
		Name:  "username",
		Value: user.Username,
	})

	mux.ServeHTTP(
		w,
		r,
	)

	user, _ = dep.Mngr_user.Get(username)

	assert.False(t, user.Setting.Bookmark_verse)
}
