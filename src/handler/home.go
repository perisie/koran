package handler

import (
	"net/http"
	"perisie.com/koran/src/guice"
	"perisie.com/koran/src/util"
)

func Home(dep *guice.Dep) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username, _ := util.Cookie_username_token(r.Cookies())
		user, _ := dep.Mngr_user.Get(username)
		surah_infos, _ := dep.Mngr_quran.Get_surah_infos()
		_ = dep.Tmpl.ExecuteTemplate(w, "page_home.html", map[string]interface{}{
			"user":        user,
			"surah_infos": surah_infos,
		})
	}
}
