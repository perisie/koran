package handler

import (
	"html/template"
	"net/http"

	"perisie.com/koran/src/user"
	"perisie.com/koran/src/util"
)

func Setting(tmpl *template.Template, mngr_user user.Mngr) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username, _ := util.Cookie_username_token(r)
		switch r.Method {
		case http.MethodPatch:
			{
				q := r.URL.Query()
				name := q.Get("name")
				value := r.FormValue("value")
				_ = mngr_user.Update_setting(username, name, value)
			}
		default:
			{
			}
		}
	}
}
