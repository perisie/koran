package handler

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"perisie.com/koran/src/quran"
	"perisie.com/koran/src/user"
	"testing"
)

func Test_surah(t *testing.T) {
	tmpl := template.Must(template.ParseGlob("../template/*.html"))
	mngr_user := user.Mngr_impl_fake()
	mngr_quran, _ := quran.Mngr_impl_new("../../qurancsv")

	_ = httptest.NewRecorder()
	_ = httptest.NewRequest(http.MethodGet, "/", nil)

	Surah(tmpl, mngr_user, mngr_quran)
}
