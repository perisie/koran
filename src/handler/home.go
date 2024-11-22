package handler

import (
	"html/template"
	"net/http"
	"perisie.com/koran/src/quran"
	"perisie.com/koran/src/user"
	"perisie.com/koran/src/util"
)

func Home(tmpl *template.Template, mngr_user user.Mngr, mngr_quran quran.Mngr) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username, _ := util.Cookie_username_token(r.Cookies())
		user, _ := mngr_user.Get(username)
		surah_infos, _ := mngr_quran.Get_surah_infos()
		_ = tmpl.ExecuteTemplate(w, "page_home.html", map[string]interface{}{
			"user":        user,
			"surah_infos": surah_infos,
		})
	}
}
