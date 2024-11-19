package handler

import (
	"html/template"
	"net/http"
	"perisie.com/koran/src/user"
)

func Signup(tmpl *template.Template, mngr_user user.Mngr) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			{
				username := r.FormValue("username")
				password := r.FormValue("password")
				_, _ = mngr_user.Create(username, password)
				w.Header().Set("HX-Redirect", "/login")
			}
		default:
			{
				_ = tmpl.ExecuteTemplate(w, "page_signup.html", map[string]interface{}{})
			}
		}
	}
}
