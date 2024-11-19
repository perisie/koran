package handler

import (
	"html/template"
	"net/http"
)

func Error(tmpl *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		switch r.Method {
		default:
			{
				code := q.Get("code")
				msg := q.Get("msg")
				_ = tmpl.ExecuteTemplate(w, "page_error.html", map[string]interface{}{
					"code": code,
					"msg":  msg,
				})
			}
		}
	}
}
