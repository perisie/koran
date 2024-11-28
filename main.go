package main

import (
	"net/http"
	"perisie.com/koran/src/guice"
	"perisie.com/koran/src/handler"
)

func main() {
	dep := guice.Dep_new(
		"src/template",
		"qurancsv",
		"static",
		"data",
	)

	tmpl := dep.Tmpl
	mngr_user := dep.Mngr_user
	mngr_quran := dep.Mngr_quran
	fs := dep.Fs

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Home(dep))
	mux.HandleFunc("/surah/{surah_id}", handler.Surah(tmpl, mngr_user, mngr_quran))
	mux.HandleFunc("/surah/{surah_id}/verse/{verse_id}", handler.Verse(tmpl, mngr_user, mngr_quran))

	mux.HandleFunc("/bookmark", handler.Bookmark(tmpl, mngr_user, mngr_quran))
	mux.HandleFunc("/setting", handler.Setting(tmpl, mngr_user))

	mux.HandleFunc("/login", handler.Login(tmpl, mngr_user))
	mux.HandleFunc("/signup", handler.Signup(tmpl, mngr_user))
	mux.HandleFunc("/user/{username}", handler.User(tmpl, mngr_user))

	mux.HandleFunc("/error", handler.Error(tmpl))
	mux.HandleFunc("/static/", handler.Static(fs))

	_ = http.ListenAndServe(":8001", mux)
}
