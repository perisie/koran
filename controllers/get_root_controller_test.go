package controllers_test

import (
	"html/template"
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

	tmpl, _ := template.ParseGlob("../template/*.html")
	quranManager, _ := managers.NewQuranManagerImpl("./../qurancsv")

	r.GET("/", controllers.GetRootController(tmpl, quranManager))

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.Nil(t, err)

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Result().StatusCode)

	_, err = io.ReadAll(w.Result().Body)
	assert.Nil(t, err)
}
