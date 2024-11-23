package handler

import (
	"html/template"
	"net/http"

	"perisie.com/koran/src/user"
	"perisie.com/koran/src/util"
)

func User(tmpl *template.Template, mngr_user user.Mngr) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username_login, _ := util.Cookie_username_token(r.Cookies())
		username := r.PathValue("username")
		user_login := false
		if username == username_login {
			user_login = true
		}
		user, _ := mngr_user.Get(username)
		tmpl.ExecuteTemplate(w, "page_user.html", map[string]interface{}{
			"user":       user,
			"user_login": user_login,
		})
	}
}
