package handler

import (
	"html/template"
	"net/http"
	"perisie.com/koran/src/user"
	"perisie.com/koran/src/util"
)

func User(tmpl *template.Template, mngr_user user.Mngr) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username, _ := util.Cookie_username_token(r.Cookies())
		user, _ := mngr_user.Get(username)
		if !user.Ok() {
			_ = tmpl.ExecuteTemplate(w, "page_error.html", map[string]interface{}{
				"code": http.StatusUnauthorized,
				"msg":  "please login",
			})
			return
		}
		tmpl.ExecuteTemplate(w, "page_user.html", map[string]interface{}{
			"user": user,
		})
	}
}
