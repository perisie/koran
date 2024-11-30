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

func Test_error(t *testing.T) {
	dep := guice.Dep_test(
		"../template",
		"../../qurancsv",
		"../../static",
	)
	mux := Mux(dep)
	w := httptest.NewRecorder()

	mux.ServeHTTP(
		w,
		httptest.NewRequest(http.MethodGet, "/error?code=500&msg=error", nil),
	)

	b, _ := io.ReadAll(w.Body)
	b_str := string(b)

	assert.True(t, strings.Contains(b_str, "500"))
	assert.True(t, strings.Contains(b_str, "error"))
}
