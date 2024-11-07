package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"perisie.com/koran/controllers"
	"perisie.com/koran/managers"
	"perisie.com/koran/requestresponse"
)

func TestGetSurahVerseControllerBadSurahQuery(t *testing.T) {
	r, w := setupGetSurahVerseController(t)

	req, err := http.NewRequest(http.MethodGet, "/surah/x/verse/x", nil)
	assert.Nil(t, err)

	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Result().StatusCode)

	var jsonError requestresponse.JsonError
	json.NewDecoder(w.Result().Body).Decode(&jsonError)
	assert.Equal(t, `strconv.Atoi: parsing "x": invalid syntax`, jsonError.Error)
}

func TestGetSurahVerseControllerBadVerseQuery(t *testing.T) {
	r, w := setupGetSurahVerseController(t)

	req, err := http.NewRequest(http.MethodGet, "/surah/1/verse/x", nil)
	assert.Nil(t, err)

	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Result().StatusCode)

	var jsonError requestresponse.JsonError
	json.NewDecoder(w.Result().Body).Decode(&jsonError)
	assert.Equal(t, `strconv.Atoi: parsing "x": invalid syntax`, jsonError.Error)
}

func TestGetSurahVerseControllerNotFound(t *testing.T) {
	r, w := setupGetSurahVerseController(t)

	req, err := http.NewRequest(http.MethodGet, "/surah/1/verse/8", nil)
	assert.Nil(t, err)

	r.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Result().StatusCode)

	var jsonError requestresponse.JsonError
	json.NewDecoder(w.Result().Body).Decode(&jsonError)
	assert.Equal(t, "verse does not exist", jsonError.Error)
}

func TestGetSurahVerseController(t *testing.T) {
	r, w := setupGetSurahVerseController(t)

	req, err := http.NewRequest(http.MethodGet, "/surah/1/verse/7", nil)
	assert.Nil(t, err)

	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)

	var jsonData requestresponse.JsonData
	json.NewDecoder(w.Result().Body).Decode(&jsonData)
	assert.Equal(t, "صِرَاطَ الَّذِينَ أَنْعَمْتَ عَلَيْهِمْ غَيْرِ الْمَغْضُوبِ عَلَيْهِمْ وَلَا الضَّالِّينَ", jsonData.Data["verse"])
}

func setupGetSurahVerseController(t *testing.T) (*gin.Engine, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	quranManager, err := managers.NewQuranManagerImpl("./../qurancsv")
	assert.Nil(t, err)

	r.GET("/surah/:surah_id/verse/:verse_id", controllers.GetSurahVerseController(quranManager))

	return r, w
}
