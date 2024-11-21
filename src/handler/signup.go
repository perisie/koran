package handler

import (
	"errors"
	"html/template"
	"net/http"
	"perisie.com/koran/src/user"
	"perisie.com/koran/src/util"
)

func Signup(tmpl *template.Template, mngr_user user.Mngr) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			{
				username := r.FormValue("username")
				password := r.FormValue("password")
				user_exist, _ := mngr_user.Get(username)
				if user_exist.Username == "" {
					_, err := mngr_user.Create(username, password)
					if err != nil {
						util.Redirect_error_page(w, http.StatusBadRequest, err)
						return
					}
					w.Header().Set("HX-Redirect", "/login")
				} else {
					util.Redirect_error_page(w, http.StatusForbidden, errors.New("username taken"))
				}
			}
		default:
			{
				_ = tmpl.ExecuteTemplate(w, "page_signup.html", map[string]interface{}{})
			}
		}
	}
}
