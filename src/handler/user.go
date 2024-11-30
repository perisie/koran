package handler

import (
	"html/template"
	"net/http"
	"perisie.com/koran/src/user"
	"perisie.com/koran/src/util"
)

func User(tmpl *template.Template, mngr_user user.Mngr) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username, _ := util.Cookie_username_token(r)
		user, _ := mngr_user.Get(username)
		if !user.Ok() {
			_ = tmpl.ExecuteTemplate(w, "page_login.html", map[string]interface{}{})
			return
		}
		tmpl.ExecuteTemplate(w, "page_user.html", map[string]interface{}{
			"user": user,
		})
	}
}
