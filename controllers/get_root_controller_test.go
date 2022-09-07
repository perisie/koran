package controllers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arikama/koran-backend/controllers"
	"github.com/arikama/koran-backend/managers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetRootController(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	quranManager, err := managers.NewQuranManagerImpl("./../qurancsv")
	assert.Nil(t, err)

	r.GET("/", controllers.GetRootController(quranManager))

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.Nil(t, err)

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Result().StatusCode)

	bytes, err := io.ReadAll(w.Result().Body)
	assert.Nil(t, err)
	assert.Equal(t, 16894, len(string(bytes)))
}
