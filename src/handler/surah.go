package handler

import (
	"html/template"
	"net/http"
	"perisie.com/koran/src/quran"
	"perisie.com/koran/src/user"
	"perisie.com/koran/src/util"
	"strconv"
)

func Surah(tmpl *template.Template, mngr_user user.Mngr, mngr_quran quran.Mngr) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username, _ := util.Cookie_username_token(r.Cookies())
		user, _ := mngr_user.Get(username)
		surah_id, _ := strconv.Atoi(r.PathValue("surah_id"))
		switch r.Method {
		default:
			{
				surah, _ := mngr_quran.Get_surah(surah_id)
				verses := []map[string]interface{}{}
				for _, verse := range surah.Verses {
					verses = append(verses, map[string]interface{}{
						"verse":            verse,
						"show_verse":       user.Setting.Surah_verse,
						"show_translation": user.Setting.Surah_translation,
					})
				}
				_ = tmpl.ExecuteTemplate(w, "page_surah.html", map[string]interface{}{
					"surah":  surah,
					"user":   user,
					"verses": verses,
				})
			}
		}
	}
}
