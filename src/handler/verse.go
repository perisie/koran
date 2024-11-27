package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"perisie.com/koran/src/opengraph"
	"perisie.com/koran/src/quran"
	"perisie.com/koran/src/user"
	"perisie.com/koran/src/util"
	"strconv"
)

func Verse(tmpl *template.Template, mngr_user user.Mngr, mngr_quran quran.Mngr) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username, _ := util.Cookie_username_token(r.Cookies())
		user, _ := mngr_user.Get(username)
		surah_id, _ := strconv.Atoi(r.PathValue("surah_id"))
		verse_id, _ := strconv.Atoi(r.PathValue("verse_id"))
		switch r.Method {
		default:
			{
				verse, _ := mngr_quran.Get_verse(surah_id, verse_id)
				og := opengraph.Og{
					Title: fmt.Sprintf("Quran %v:%v", surah_id, verse_id),
					Type:  "book",
					Url:   "https://koran.perisie.com/" + r.RequestURI,
					Image: "https://koran.perisie.com/static/koran.png",
				}
				_ = tmpl.ExecuteTemplate(w, "page_verse.html", map[string]interface{}{
					"verse":            verse,
					"user":             user,
					"show_verse":       user.Setting.Surah_verse,
					"show_translation": user.Setting.Surah_translation,
					"og":               og,
				})
			}
		}
	}
}
