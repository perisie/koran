package handler

import (
	"html/template"
	"net/http"
	"perisie.com/koran/src/quran"
)

func Home(tmpl *template.Template, mngr_quran quran.Mngr) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		surah_infos, _ := mngr_quran.Get_surah_infos()
		_ = tmpl.ExecuteTemplate(w, "page_home.html", map[string]interface{}{
			"surah_infos": surah_infos,
		})
	}
}
