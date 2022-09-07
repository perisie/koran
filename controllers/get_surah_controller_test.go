package controllers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arikama/koran-backend/controllers"
	"github.com/arikama/koran-backend/managers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetSurahController(t *testing.T) {
	r, w := setupGetSurahController(t)

	req, err := http.NewRequest(http.MethodGet, "/surah/1", nil)
	assert.Nil(t, err)

	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)

	bytes, err := io.ReadAll(w.Result().Body)
	assert.Nil(t, err)
	assert.True(t, strings.Contains(string(bytes), "بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ"))
	assert.True(t, strings.Contains(string(bytes), "صِرَاطَ الَّذِينَ أَنْعَمْتَ عَلَيْهِمْ غَيْرِ الْمَغْضُوبِ عَلَيْهِمْ وَلَا الضَّالِّينَ"))
}

func TestGetSurahControllerInvalidRequest(t *testing.T) {
	r, w := setupGetSurahController(t)

	req, err := http.NewRequest(http.MethodGet, "/surah/x", nil)
	assert.Nil(t, err)

	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Result().StatusCode)
}

func TestGetSurahControllerSurahNotExist(t *testing.T) {
	r, w := setupGetSurahController(t)

	req, err := http.NewRequest(http.MethodGet, "/surah/115", nil)
	assert.Nil(t, err)

	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Result().StatusCode)

	bytes, err := io.ReadAll(w.Result().Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"error":"surah does not exist"}`, string(bytes))
}

func setupGetSurahController(t *testing.T) (*gin.Engine, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	quranManager, err := managers.NewQuranManagerImpl("./../qurancsv")
	assert.Nil(t, err)

	r.GET("/surah/:surah_id", controllers.GetSurahController(quranManager))

	return r, w
}
