package main

import (
	"html/template"
	"net/http"
	"perisie.com/koran/src/handler"
	"perisie.com/koran/src/quran"
	"perisie.com/koran/src/user"
)

func main() {
	var tmpl *template.Template
	var mngr_quran quran.Mngr
	var mngr_user user.Mngr

	tmpl = template.Must(template.ParseGlob("src/template/*.html"))
	fs := http.FileServer(http.Dir("static"))
	mngr_quran, _ = quran.Mngr_impl_new("qurancsv")
	mngr_user = user.Mngr_impl_new()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Home(tmpl, mngr_user, mngr_quran))
	mux.HandleFunc("/surah/{surah_id}", handler.Surah(tmpl, mngr_user, mngr_quran))

	mux.HandleFunc("/bookmark", handler.Bookmark(tmpl, mngr_user, mngr_quran))

	mux.HandleFunc("/login", handler.Login(tmpl, mngr_user))
	mux.HandleFunc("/signup", handler.Signup(tmpl, mngr_user))

	mux.HandleFunc("/error", handler.Error(tmpl))
	mux.HandleFunc("/static/", handler.Static(fs))

	_ = http.ListenAndServe(":8001", mux)
}
