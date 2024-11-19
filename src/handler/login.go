package handler

import (
	"html/template"
	"net/http"
	"perisie.com/koran/src/user"
)

func Login(tmpl *template.Template, mngr_user user.Mngr) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			{
				username := r.FormValue("username")
				user, err := mngr_user.Get(username)
				if err != nil {
					w.Header().Set("HX-Redirect", "/error?code=code&msg=msg")
					return
				}
				if user.Username == "" {
					w.Header().Set("HX-Redirect", "/error?code=code&msg=msg")
					return
				}
				w.Header().Set("HX-Redirect", "/")
			}
		default:
			{
				_ = tmpl.ExecuteTemplate(w, "page_login.html", map[string]interface{}{})
			}
		}
	}
}
