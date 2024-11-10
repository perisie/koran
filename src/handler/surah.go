package handler

import (
	"html/template"
	"net/http"
	"perisie.com/koran/src/quran"
	"strconv"
)

func Surah(tmpl *template.Template, mngr_quran quran.Mngr) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		surah_id, _ := strconv.Atoi(r.PathValue("surah_id"))
		switch r.Method {
		default:
			{
				surah, _ := mngr_quran.Get_surah(surah_id)
				_ = tmpl.ExecuteTemplate(w, "page_surah.html", map[string]interface{}{
					"surah": surah,
				})
			}
		}
	}
}
