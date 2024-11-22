package handler

import (
	"errors"
	"html/template"
	"net/http"
	"perisie.com/koran/src/user"
	"perisie.com/koran/src/util"
)

func Login(tmpl *template.Template, mngr_user user.Mngr) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			{
				username := r.FormValue("username")
				user, err := mngr_user.Get(username)
				if err != nil {
					util.Redirect_error_page(w, http.StatusInternalServerError, errors.New("user not found"))
					return
				}
				if user.Username == "" {
					util.Redirect_error_page(w, http.StatusInternalServerError, errors.New("user not found"))
					return
				}
				tmpl.ExecuteTemplate(w, "comp_login_ok.html", map[string]interface{}{
					"user": user,
				})
			}
		case http.MethodDelete:
			{
				tmpl.ExecuteTemplate(w, "comp_login_out.html", map[string]interface{}{})
			}
		default:
			{
				_ = tmpl.ExecuteTemplate(w, "page_login.html", map[string]interface{}{})
			}
		}
	}
}
