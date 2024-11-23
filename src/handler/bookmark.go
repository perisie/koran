package handler

import (
	"html/template"
	"net/http"
	"perisie.com/koran/src/quran"
	"perisie.com/koran/src/user"
	"perisie.com/koran/src/util"
)

func Bookmark(tmpl *template.Template, mngr_user user.Mngr, mngr_quran quran.Mngr) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username, token := util.Cookie_username_token(r.Cookies())
		user, _ := mngr_user.Get(username)
		if !user.Ok() || user.Password != token {
			tmpl.ExecuteTemplate(w, "page_login.html", nil)
			return
		}
		switch r.Method {
		default:
			{
				tmpl.ExecuteTemplate(w, "page_bookmark.html", map[string]interface{}{
					"user": user,
				})
			}
		}
	}
}
