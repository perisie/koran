package handler

import (
	"net/http"
)

func Static(fs http.Handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", fs).ServeHTTP(w, r)
	}
}
