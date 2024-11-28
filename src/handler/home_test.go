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

func Test_home(t *testing.T) {
	dep := guice.Dep_test(
		"../template",
		"../../qurancsv",
		"../../static",
	)
	mux := Mux(dep)
	w := httptest.NewRecorder()

	mux.ServeHTTP(
		w,
		httptest.NewRequest(http.MethodGet, "/", nil),
	)

	b, _ := io.ReadAll(w.Body)
	b_str := string(b)

	assert.True(t, strings.Contains(b_str, "The Opening"))
	assert.True(t, strings.Contains(b_str, "Mankind"))
}
