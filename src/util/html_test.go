package util

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_redirect_error_page(t *testing.T) {
	w := httptest.NewRecorder()

	Redirect_error_page(w, http.StatusInternalServerError, errors.New("some error"))

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, "/error?code=500&msg=some error", w.Header().Get("HX-Redirect"))
}
