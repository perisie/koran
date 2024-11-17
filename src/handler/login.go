package handler

import (
	"html/template"
	"net/http"
	"perisie.com/koran/src/user"
)

func Login(tmpl *template.Template, mngr_user user.Mngr) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = tmpl.ExecuteTemplate(w, "page_login.html", map[string]interface{}{})
	}
}
